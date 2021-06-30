package ascTrim

func AscTrim20(s string) string {
	arr := []byte(s)
	var result []byte
	for _, el := range arr {
		if el >= 32 {
			result = append(result, el)
		}
	}
	return string(result)
}
