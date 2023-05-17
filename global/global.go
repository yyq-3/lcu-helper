package global

import (
	"github.com/sacOO7/gowebsocket"
	"lcu-helper/dto"
)

var ClientUx = &dto.ClientStatus{
	ProcessName: "LeagueClientUx.exe",
	Port:        0,
	Token:       "",
}

var ClientSocket gowebsocket.Socket
