func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string) // buat array yang terdiri dari key of 26 arrays,
										// dengan value type slice of string
	for _, str := range strs { // "act"
		var count[26]int
		for _, char := range str {
			count[char - 'a']++ // menggunakan index sebagai posisi huruf, 
								// dan value sebagai frequensi
		}
		groups[count] = append(groups[count], str)		
	}
	result := make([][]string, 0, len(groups))
	for _, group := range groups {
			result = append(result, group)
		}
	return result
}
