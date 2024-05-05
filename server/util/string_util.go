package util

import "strconv"

func Unit2String(v uint) string {
	return strconv.FormatUint(uint64(v), 10)
}
