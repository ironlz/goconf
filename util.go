package goconf

func spiltFileNameAndType(fileName string) (string, string) {
	lastDot := 0
	for i, ch := range fileName {
		if ch == '.' {
			lastDot = i
		}
	}
	if lastDot == 0 {
		return fileName, ""
	}
	return fileName[:lastDot], fileName[lastDot+1:]
}
