package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := newWordMapFromList(wordList)
	if _, found := wordMap[endWord]; !found {
		return 0
	}

	depth := 0
	queue := []string{beginWord}
	wordLen := len(beginWord)
	for len(queue) > 0 {
		depth++

		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			word := []rune(queue[0]) // []rune allows for word[i] = newRune // front
			queue = queue[1:]        // pop

			for i := 0; i < wordLen; i++ {
				tmpLetter := word[i] // saves it temporarily
				for letter := 'a'; letter <= 'z'; letter++ {
					word[i] = letter
					if string(word) == endWord {
						return depth + 1
					}

					if _, found := wordMap[string(word)]; !found {
						continue // next letter
					}

					queue = append(queue, string(word)) // next level
					delete(wordMap, string(word))       // not going back
				}
				word[i] = tmpLetter // sets it back
			}
		}
	}

	return 0
}

func newWordMapFromList(wordList []string) map[string]struct{} {
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	return wordMap
}
