package calculate

import (
	"github.com/aleksandra0KR/networks-itmo-is/internal/ip"
	"github.com/aleksandra0KR/networks-itmo-is/internal/network"
	"github.com/aleksandra0KR/networks-itmo-is/internal/subnet"
)

type Data struct {
	ip       ip.IPAddress
	subnet   subnet.Subnet
	networks []network.Network
}

func NewData(ipaddress, subnetmask string, userAmounts []int) (*Data, error) {
	ip, err := ip.FromString(ipaddress)
	if err != nil {
		return &Data{}, err
	}

	subnet, err := subnet.FromString(subnetmask)
	if err != nil {
		return &Data{}, err
	}

	networks, err := network.FromInts(userAmounts)
	if err != nil {
		return &Data{}, err
	}

	return &Data{
		ip:       ip,
		subnet:   subnet,
		networks: networks,
	}, nil
}

func Calculate() {
}
