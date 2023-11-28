package util

func GenerateString(length int) string {
	var str string
	for i := 0; i < length; i++ {
		str += "a"
	}
	return str
}
