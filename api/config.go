package api

func AcceptGame(baseAddr string) (err error) {
	_, err = post(baseAddr, CONFIG_AUTO_ACCEPT)
	return err
}
