package main

func main() {
	println("1. ", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println("2. ", longestCommonPrefix([]string{"dog", "racecar", "car"}), "end")
	println("3. ", longestCommonPrefix([]string{""}), "end")
	println("4. ", longestCommonPrefix([]string{"a"}), "end")
	println("5. ", longestCommonPrefix([]string{"ab", "a"}), "end")
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	minStr := strs[0]
	for _, str := range strs {
		if str == "" {
			return ""
		}

		if len(str) < len(minStr) {
			minStr = str
		}
	}

	end := 1
	//benchmark := strs[0][:end] // sub str
	//println(benchmark)

	for end <= len(minStr) {
		for _, elem := range strs {
			if elem[:end] != strs[0][:end] {
				//return benchmark[:end-1]
				return strs[0][:end-1]
			}
		}

		if end == len(minStr) {
			//return benchmark
			return strs[0][:end]
		}
		// update
		end++
		//benchmark = strs[0][:end]
	}

	return ""
}
