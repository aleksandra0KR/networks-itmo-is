package arithmeticOperations

import "strings"

func BinaryAnd(bin1, bin2 string) string {
	var result strings.Builder
	for i := 0; i < len(bin1); i++ {
		if bin1[i] == '1' && bin2[i] == '1' {
			result.WriteByte('1')
		} else {
			result.WriteByte('0')
		}
	}
	return result.String()
}
