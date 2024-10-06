package internal

import (
	"github.com/aleksandra0KR/networks-itmo-is/labs/lab2/internal/convert"
	"math"
	"strings"
)

func CalculateMask(amountOfComputers int) string {
	bytesInIPv4 := 32
	amountZeros := calculateBitsToRepresentAmountOfNodes(amountOfComputers)
	amountOfOnes := bytesInIPv4 - amountZeros
	maskBinary := strings.Repeat("1", amountOfOnes) + strings.Repeat("0", amountZeros)
	return convert.ConvertBinaryToIp(maskBinary)
}

func calculateBitsToRepresentAmountOfNodes(amount int) int {
	amountInLogForm := math.Log2(float64(amount))
	rounded := int(math.Ceil(amountInLogForm))
	return rounded
}
