package abstract_factory

type ExpensiveTransport struct {
	transportType string
	service       string
}

func (expensiveTransport *ExpensiveTransport) SetTransportType(transportType string) {
	expensiveTransport.transportType = transportType
}
func (expensiveTransport *ExpensiveTransport) SetService(service string) {
	expensiveTransport.service = service
}
func (expensiveTransport *ExpensiveTransport) GetTransportType() string {
	return expensiveTransport.transportType
}
func (expensiveTransport *ExpensiveTransport) GetService() string {
	return expensiveTransport.service
}
