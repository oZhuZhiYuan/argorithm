package argorithm

func BracketIsValid(s string) bool {
	brackets_map := map[byte]byte{'}': '{', ']': '[', ')': '('}
	stack := make([]byte, 0)
	for _, a := range []byte(s) {
		if _, ok := brackets_map[a]; !ok {
			stack = append(stack, a)
			continue
		}
		if len(stack) == 0 || brackets_map[a] != stack[len(stack)-1] {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	return len(stack) == 0
}
