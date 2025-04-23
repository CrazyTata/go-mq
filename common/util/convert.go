package util

// BoolToUint8 将布尔值转换为uint8
func BoolToUint8(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}
