package main


func MakeUpperCase(str string) string {
	var upperCase string

	for i := 0; i < len(str); i++ {
		if str[i] <= 'z' && str[i] >= 'a' {
			upperCase += string(str[i]-'a' + 'A')
		} else {
			upperCase += string(str[i])
		}
	}
	return upperCase
}