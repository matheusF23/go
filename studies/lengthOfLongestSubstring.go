package main

func lengthOfLongestSubstring(s string) int {
	mapper := make(map[byte]int)
	max := 0
	for i := 0; i < len(s); i++ {
		if _, existKey := mapper[s[i]]; !existKey {
			mapper[s[i]] = i
			if len(mapper) > max {
				max = len(mapper)
			}
		} else {
			i = mapper[s[i]]
			mapper = map[byte]int{}
		}
	}
	return max
}
