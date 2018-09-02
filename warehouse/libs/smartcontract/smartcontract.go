package smartcontract

type smartcontract interface {
	Install() error
	Query() error
	Invoke() error
}
