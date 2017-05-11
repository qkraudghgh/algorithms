package backtrack

func allPerms(elements string) []string {
	tmp := []string{}
	if len(elements) <= 1 {
		return append(tmp, elements)
	} else {
		for _, value := range allPerms(elements[1:]) {
			for i := 0; i < len(elements); i++ {
				tmp = append(tmp, value[:i]+elements[0:1]+value[i:])
			}
		}
		return tmp
	}
}
