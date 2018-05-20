package server

type Cfg struct {
	CfgRootWebServer *CfgWebServer
}

type CfgWebServer struct {
	Addr string
}