package subnet

import (
	"fmt"
	"strconv"
	"strings"
)

type Subnet struct {
	octet1 int8
	octet2 int8
	octet3 int8
	octet4 int8
}

func FromString(str string) (Subnet, error) {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return Subnet{}, fmt.Errorf("invalid IP address format")
	}

	var subnet Subnet

	for i, octetStr := range octets {
		octet, err := strconv.ParseInt(octetStr, 10, 8)
		if err != nil || octet < 0 || octet > 255 {
			return Subnet{}, fmt.Errorf("invalid octet value: %s", octetStr)
		}

		switch i {
		case 0:
			subnet.octet1 = int8(octet)
		case 1:
			subnet.octet2 = int8(octet)
		case 2:
			subnet.octet3 = int8(octet)
		case 3:
			subnet.octet4 = int8(octet)
		}
	}

	return subnet, nil
}
