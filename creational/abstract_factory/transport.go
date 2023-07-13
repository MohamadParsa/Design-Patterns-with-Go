package abstract_factory

type TransportFactory interface {
}

type Transport interface {
	GetTransportType() string
	GetService() string
}
