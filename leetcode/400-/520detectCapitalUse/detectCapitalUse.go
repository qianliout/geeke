package main

func main() {

}

func detectCapitalUse(word string) bool {
	hasUp, hasLow, up := false, false, 0
	for i := range word {
		if word[i] >= 'a' && word[i] <= 'z' {
			hasLow = true
		}
		if word[i] >= 'A' && word[i] <= 'Z' {
			hasUp = true
			up = i
		}
	}
	if hasUp && hasLow && up != 0 {
		return false
	}

	return true
}
