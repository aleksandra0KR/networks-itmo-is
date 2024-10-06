package internal

import (
	"github.com/aleksandra0KR/networks-itmo-is/labs/lab2/internal/arithmeticOperations"
	"github.com/aleksandra0KR/networks-itmo-is/labs/lab2/internal/convert"
)

func CalculateNetworkIP(ipAddress, mask string) string {
	ipBinary := convert.ConvertIpToBinary(ipAddress)
	maskBinary := convert.ConvertIpToBinary(mask)

	networkIP := arithmeticOperations.BinaryAnd(ipBinary, maskBinary)
	return networkIP
}
