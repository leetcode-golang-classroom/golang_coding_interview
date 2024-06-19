package internal

func CountSubString(inputs []string) []int32 {
	results := []int32{}
	for _, input := range inputs {
		results = append(results, CountSubPrefix(input))
	}
	return results
}

func CountSubPrefix(input string) int32 {
	var count int32 = int32(len(input))
	prefix := ""
	for pos := range input {
		prefix += string(input[pos])
		suffix := input[pos+1:]
		// calculate common substring length
		subLen := CommonSubStringLen(suffix, prefix)
		count += subLen
	}
	return count
}

func CommonSubStringLen(a, b string) int32 {
	var count int32 = 0
	alen, blen := len(a), len(b)
	slen := alen
	if slen > blen {
		slen = blen
	}
	for pos := 0; pos < slen; pos++ {
		if a[pos] != b[pos] {
			return count
		} else {
			count++
		}
	}
	return count
}
