package main

type WordSet map[string]struct{}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	// from slice to set
	wordSet := WordSet{}
	for _, word := range wordList {
		wordSet[word] = struct{}{}
	}

	// Return if the end word is not in the list at all.
	if _, found := wordSet[endWord]; !found {
		return 0
	}

	wordLen := len(beginWord)

	smallWordSet := WordSet{}
	smallWordSet[beginWord] = struct{}{}

	bigWordSet := WordSet{}
	bigWordSet[endWord] = struct{}{}

	depth := 0
	for len(smallWordSet) != 0 && len(bigWordSet) != 0 {
		depth++

		if len(smallWordSet) > len(bigWordSet) {
			smallWordSet, bigWordSet = bigWordSet, smallWordSet
		}

		// Loop through the small list and record the words on the next level in a tmp set
		tmpWordSet := WordSet{}
		for keyWord := range smallWordSet {
			for i := 0; i < wordLen; i++ {
				runes := []rune(keyWord) // []rune allows for runes[i] = newRune
				tmpRune := runes[i]      // saves it temporarily
				for rn := 'a'; rn <= 'z'; rn++ {
					runes[i] = rn
					word := string(runes)
					if _, found := bigWordSet[word]; found {
						return depth + 1
					}

					if _, found := wordSet[word]; !found {
						continue // next rune
					}

					delete(wordSet, word) // not going backward
					tmpWordSet[string(runes)] = struct{}{}
				}
				runes[i] = tmpRune // sets it back
			}
		}
		smallWordSet = tmpWordSet // the tmp set becoming the new small set since we are done with the old small set
	}

	return 0
}

func main() {
	start := "hit"
	end := "cog"
	words := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	print(ladderLength(start, end, words))
}
