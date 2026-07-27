import "reflect"

func isAnagram(s string, t string) bool {
	s_map := make(map[rune]int)
	t_map := make(map[rune]int)
	for _, val := range s {
		s_map[val] += 1
	}
	for _, val := range t {
		t_map[val] += 1
	}
	return reflect.DeepEqual(s_map, t_map)	
}
