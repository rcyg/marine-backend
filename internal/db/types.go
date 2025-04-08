package db

type PortThroughput struct {
	PortName string    `json:"port_name,omitempty"`
	PortCode string    `json:"port_code,omitempty"`
	In       []float64 `json:"in,omitempty"`
	Out      []float64 `json:"out,omitempty"`
}
