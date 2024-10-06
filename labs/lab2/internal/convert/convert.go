package convert

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertIpToBinary(ip string) string {
	parts := strings.Split(ip, ".")
	var binaryParts []string
	for _, part := range parts {
		intPart, _ := strconv.Atoi(part)
		binaryParts = append(binaryParts, fmt.Sprintf("%8b", intPart))
	}
	return strings.Join(binaryParts, "")
}

func ConvertBinaryToIp(binaryIp string) string {
	var parts []string
	for i := 0; i < len(binaryIp); i += 8 {
		octet := binaryIp[i : i+8]
		parts = append(parts, fmt.Sprintf("%d", binaryToDecimal(octet)))
	}
	return strings.Join(parts, ".")
}

func binaryToDecimal(binary string) int {
	var decimal int
	for _, bit := range binary {
		decimal = decimal*2 + int(bit-'0')
	}
	return decimal
}
