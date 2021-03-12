func findFirstStringInBracket(str string) (result string) {
	if len(str) > 0 {

		firstSplit := strings.Split(str, "(")

		if len(firstSplit) == 1 {
			return ""
		}

		secondSplit := strings.Split(firstSplit[1], ")")

		if len(secondSplit) == 1 {
			return ""
		}

		splittedResult := strings.Split(secondSplit[0], "")
		maxLength := len(splittedResult)

		result = strings.Join(splittedResult[:maxLength-1], "")
	}

	return
}