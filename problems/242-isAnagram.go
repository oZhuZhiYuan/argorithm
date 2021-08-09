package problems

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map1, map2 := make(map[byte]int), make(map[byte]int)
	for _, b := range []byte(s) {
		map1[b]++
	}
	for _, b := range []byte(t) {
		map2[b]++
	}
	if len(map1) != len(map2) {
		return false
	}
	for k := range map1 {
		if map1[k] != map2[k] {
			return false
		}
	}
	return true
}
