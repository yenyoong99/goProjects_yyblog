package ioc

type Object interface {
	Init() error
	Destroy() error
}
