package config

import "fmt"

type ServerConfiguration struct {
	Address string
	Port    uint
}

func NewServerConfig(port uint, hostaddress string) *ServerConfiguration {
	return &ServerConfiguration{
		Port:    port,
		Address: hostaddress,
	}
}

func (sc *ServerConfiguration) GetEndpoint() string {
	return fmt.Sprintf("%s:%d", sc.Address, sc.Port)
}
