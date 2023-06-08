package lcu

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"lcu-helper/internal/util"
	"lcu-helper/pkg/logger"
	"net/http"
	"net/http/httputil"
	"time"
)

func startProxy() {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	http.HandleFunc("/", serverProxy)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		logger.Info("代理启动失败！！！")
	}
	logger.Info("启动代理成功！！")
}

func serverProxy(w http.ResponseWriter, r *http.Request) {
	director := func(req *http.Request) { // 最终要的，转发定向处理
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString(util.Str2byte(fmt.Sprintf("riot:%s", ClientUx.Token))))
		req.Header.Set("Content-Type", "application/json")
		req.Host = fmt.Sprintf("127.0.0.1:%d", ClientUx.Port)
		req.URL.Scheme = "https"
		req.URL.Host = fmt.Sprintf("127.0.0.1:%d", ClientUx.Port) // 具体转发ip和端口
		req.URL.Path = r.RequestURI                               // 转发的服务具体接口路径  /api/v1/almsvr
		req.Method = r.Method                                     // 接口请求方法类型get、post、put、delete
		req.Body = r.Body
	}

	transport := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, ResponseHeaderTimeout: time.Second}

	errHandler := func(writer http.ResponseWriter, request *http.Request, err error) {
		if err != nil {
			logger.Infof("发生错误，错误信息： %s, 请求地址%s%s", err.Error(), request.Host, request.RequestURI)
		}
	}

	resHandler := func(response *http.Response) error { // 返回结果处理
		return nil
	}

	proxy := &httputil.ReverseProxy{Director: director, ModifyResponse: resHandler, ErrorHandler: errHandler}
	proxy.Transport = transport
	proxy.ServeHTTP(w, r) // http服务请求
}
