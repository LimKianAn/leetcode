package main

func findLadders(beginWord string, endWord string, wordList []string) (ladders [][]string) {
	dict := map[string]struct{}{}
	for _, word := range wordList {
		dict[word] = struct{}{}
	}

	if _, found := dict[endWord]; !found {
		return
	}
	delete(dict, endWord) // avoids going back and forth

	curDict := map[string]struct{}{beginWord: {}} // current dict
	found := false
	wordLen := len(beginWord)
	children := map[string][]string{}
	for len(curDict) != 0 && !found {
		for word := range curDict {
			delete(dict, word) // avoids going back and forth
		}

		tmpDict := map[string]struct{}{}
		for word := range curDict {
			runes := []rune(word) // new runes for manipulating each letter
			for i := 0; i < wordLen; i++ {
				rn := runes[i]        // records the original letter
				for letter := 'a'; letter <= 'z'; letter++ {
					runes[i] = letter
					if child := string(runes); child == endWord {
						found = true
						children[word] = append(children[word], child)
					} else if _, ok := dict[child]; ok && !found {
						tmpDict[child] = struct{}{} // next depth
						children[word] = append(children[word], child)
					}
				}
				runes[i] = rn // sets it back
			}
		}
		curDict = tmpDict
	}

	if found {
		var appendLadders func(ladder []string, word string)
		appendLadders = func(ladder []string, word string) {
			if word == endWord {
				copiedLadder := make([]string, len(ladder)) // ladder is shared. We must make a copy!
				copy(copiedLadder, ladder)
				ladders = append(ladders, copiedLadder)
				return
			}

			for _, child := range children[word] {
				ladder = append(ladder, child)
				appendLadders(ladder, child)
				ladder = ladder[:len(ladder)-1]
			}
		}
		appendLadders([]string{beginWord}, beginWord)
	}

	return
}

// func main() {
// 	log.Println(findLadders("red", "tax", []string{"ted", "tex", "red", "tax", "tad", "den", "rex", "pee"}))
// }
