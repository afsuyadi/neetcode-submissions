func isPalindrome(s string) bool {
	// remove the space first,
	// change all letter to lowercase
	// compare first element with last element
	// if all same, return true
	// else return false
	var rawword []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == ' ' {
			continue
		}
		
		isLower := c >= 'a' && c <= 'z'
		isUpper := c >= 'A' && c <= 'Z'
		isNumber := c >= '0' && c <= '9'
		if !isLower && !isUpper && !isNumber {
			continue
		} 

		if isUpper {
			c = c + ('a' - 'A') // 97 - 65
		}

		rawword = append(rawword, c)
	}
	combinedword := string(rawword)

	for i := 0; i < len(combinedword)/2; i++ {
		j := len(combinedword) - i - 1
		if combinedword[i] != combinedword[j] {
			return false
		}
	}
	return true //true
}
