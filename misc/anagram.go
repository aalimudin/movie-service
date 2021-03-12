func groupAnagram(stringSlice []string) []string {
	anagramGroup := make([][]string, len(stringSlice))
	currentIndex := 0
	nextIndex := false
	a := 0
	for i := 0; i < len(stringSlice); i++ {

		if i == 0 {
			anagramGroup[i] = append(anagramGroup[i], stringSlice[i])
			continue
		}

		splittedAnagram := strings.Split(stringSlice[i], "")

		if i > 0 {
			for o := 0; o < len(anagramGroup); o++ {
				if len(anagramGroup[o]) > 0 {
					a++
				}
			}

			for x := 0; x < a; x++ {
				splittedResult := strings.Split(result[x][0], "")

				if len(splittedAnagram) == len(splittedResult) {
					for y := 0; y < len(splittedAnagram); y++ {
						for z := 0; z < len(splittedResult); z++ {
							if splittedAnagram[y] == splittedResult[z] {
								nextIndex = false
								break
							} else {
								nextIndex = true
							}
						}
					}

					if !nextIndex {
						anagramGroup[x] = append(anagramGroup[x], stringSlice[i])
						continue
					} else if nextIndex {
						currentIndex++
						anagramGroup[currentIndex] = append(anagramGroup[currentIndex], stringSlice[i])
						continue
					}
				} else if len(splittedAnagram) != len(splittedResult) && stringSlice[i] != anagramGroup[x][0] {
					currentIndex++
					anagramGroup[currentIndex] = append(anagramGroup[currentIndex], stringSlice[i])
					break
				}
			}
		}
	}

	return anagramGroup
}

func anagram () {
	strings := []string{"kita", "atik", "tika", "aku", "kia", "tes", "makan", "set", "kua"}
	anagramGroup := groupAnagram(strings)

	fmt.Println(anagramGroup)
}