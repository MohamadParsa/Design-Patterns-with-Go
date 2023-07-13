package economic

type EconomicTransport struct {
	transportType string
}

func (economicTransport *EconomicTransport) SetTransportType(transportType string) {
	economicTransport.transportType = transportType
}
func (economicTransport *EconomicTransport) GetTransportType() string {
	return economicTransport.transportType
}
func (economicTransport *EconomicTransport) GetService() string {
	return "Nothing"
}
