package ip

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aleksandra0KR/networks-itmo-is/internal/subnet"
)

type IPAddress struct {
	octet1 uint8
	octet2 uint8
	octet3 uint8
	octet4 uint8
}

func FromString(str string) (IPAddress, error) {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return IPAddress{}, fmt.Errorf("invalid IP address format")
	}

	var ip IPAddress
	for i, octetStr := range octets {
		octet, err := strconv.ParseUint(octetStr, 10, 8)
		if err != nil || octet < 0 || octet > 255 {
			return IPAddress{}, fmt.Errorf("invalid octet value: %s", octetStr)
		}

		switch i {
		case 0:
			ip.octet1 = uint8(octet)
		case 1:
			ip.octet2 = uint8(octet)
		case 2:
			ip.octet3 = uint8(octet)
		case 3:
			ip.octet4 = uint8(octet)
		}
	}

	return ip, nil
}

func (ip IPAddress) Add(n int) IPAddress {
	newIP := int(ip.octet1)<<24 |
		int(ip.octet2)<<16 |
		int(ip.octet3)<<8 |
		int(ip.octet4)

	newIP += n

	return IPAddress{
		octet1: uint8((newIP >> 24) & 0xFF),
		octet2: uint8((newIP >> 16) & 0xFF),
		octet3: uint8((newIP >> 8) & 0xFF),
		octet4: uint8(newIP & 0xFF),
	}
}

func (ip IPAddress) String() string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		ip.octet1,
		ip.octet2,
		ip.octet3,
		ip.octet4,
	)
}

func (ip IPAddress) NetworkAddress(mask subnet.Subnet) IPAddress {
	return IPAddress{
		octet1: ip.octet1 & mask.GetOctet1(),
		octet2: ip.octet2 & mask.GetOctet2(),
		octet3: ip.octet3 & mask.GetOctet3(),
		octet4: ip.octet4 & mask.GetOctet4(),
	}
}
