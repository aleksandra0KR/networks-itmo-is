package calculate

import (
	"fmt"
	"sort"
	"strings"

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

func Calculate(data *Data) string {
	sort.Slice(data.networks, func(i, j int) bool {
		return data.networks[i].UsersAmount > data.networks[j].UsersAmount
	})

	var result strings.Builder

	currentIP := data.ip

	for _, net := range data.networks {
		const reservedAddresses = 2
		requiredSize := net.UsersAmount + reservedAddresses

		subnetMask, err := subnet.CalculateMask(requiredSize)
		if err != nil {
			return fmt.Sprintf("failed to calculate mask: %v", err)
		}

		startIP := currentIP.NetworkAddress(subnetMask)
		availableIPs := subnetMask.AvailableIPs()

		if availableIPs < requiredSize {
			return fmt.Sprintf(
				"not enough addresses for network with %d users",
				net.UsersAmount,
			)
		}

		const inclusiveRange = 1
		endIP := currentIP.Add(availableIPs - inclusiveRange)

		result.WriteString(fmt.Sprintf(
			"Network with %d users: %s/%d, Range: %s - %s\n",
			net.UsersAmount,
			startIP,
			subnetMask.PrefixLength(),
			startIP,
			endIP,
		))

		currentIP = endIP.Add(1)
	}

	return result.String()
}
