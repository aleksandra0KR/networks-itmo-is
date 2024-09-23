package subnet

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Subnet struct {
	octet1 uint8
	octet2 uint8
	octet3 uint8
	octet4 uint8
}

func FromString(str string) (Subnet, error) {
	octets := strings.Split(str, ".")
	if len(octets) != 4 {
		return Subnet{}, fmt.Errorf("invalid IP address format")
	}

	var subnet Subnet

	for i, octetStr := range octets {
		octet, err := strconv.ParseUint(octetStr, 10, 8)
		if err != nil || octet < 0 || octet > 255 {
			return Subnet{}, fmt.Errorf("invalid octet value: %s", octetStr)
		}

		switch i {
		case 0:
			subnet.octet1 = uint8(octet)
		case 1:
			subnet.octet2 = uint8(octet)
		case 2:
			subnet.octet3 = uint8(octet)
		case 3:
			subnet.octet4 = uint8(octet)
		}
	}

	return subnet, nil
}

func CalculateMask(requiredSize int) (Subnet, error) {
	if requiredSize <= 0 {
		return Subnet{}, fmt.Errorf("invalid required size")
	}

	bitsNeeded := int(math.Ceil(math.Log2(float64(requiredSize))))
	if bitsNeeded > 32 {
		return Subnet{}, fmt.Errorf("too large network")
	}

	prefixLength := 32 - bitsNeeded
	return FromPrefixLength(prefixLength)
}

func FromPrefixLength(prefixLength int) (Subnet, error) {
	if prefixLength < 0 || prefixLength > 32 {
		return Subnet{}, fmt.Errorf("invalid prefix length")
	}

	var subnet Subnet
	for i := 0; i < prefixLength; i++ {
		switch {
		case i < 8:
			subnet.octet1 |= 1 << (7 - i)
		case i < 16:
			subnet.octet2 |= 1 << (15 - i)
		case i < 24:
			subnet.octet3 |= 1 << (23 - i)
		case i < 32:
			subnet.octet4 |= 1 << (31 - i)
		}
	}

	return subnet, nil
}

func (subnet Subnet) AvailableIPs() int {
	ones := subnet.PrefixLength()
	return int(math.Pow(2, float64(32-ones)))
}

func (subnet Subnet) PrefixLength() int {
	count := 0
	for _, octet := range []uint8{
		subnet.octet1,
		subnet.octet2,
		subnet.octet3,
		subnet.octet4,
	} {
		for i := 7; i >= 0; i-- {
			if (octet>>i)&1 == 1 {
				count++
			} else {
				return count
			}
		}
	}
	return count
}

func (subnet Subnet) String() string {
	return fmt.Sprintf(
		"%d.%d.%d.%d",
		subnet.octet1,
		subnet.octet2,
		subnet.octet3,
		subnet.octet4,
	)
}

func (subnet *Subnet) GetOctet1() uint8 {
	return subnet.octet1
}

func (subnet *Subnet) GetOctet2() uint8 {
	return subnet.octet1
}

func (subnet *Subnet) GetOctet3() uint8 {
	return subnet.octet1
}

func (subnet *Subnet) GetOctet4() uint8 {
	return subnet.octet1
}
