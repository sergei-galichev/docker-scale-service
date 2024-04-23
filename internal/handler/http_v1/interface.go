package http_v1

type ServerHttp interface {
	Run() error
	Shutdown() error
}
