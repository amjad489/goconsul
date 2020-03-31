package core

func SanitizeName(name string, limit int) string {
	result := name
	chars := 0
	for i := range name {
		if chars >= limit {
			result = name[:i]
			break
		}
		chars++
	}
	return result
}
