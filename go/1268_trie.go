// 2022.01.18

package main

import "sort"

func suggestedProducts_t(products []string, searchWord string) [][]string {
	type trie struct {
		child map[rune]*trie
		words []string
	}

	addProduct := func(root *trie, product string) {
		for _, rn := range product {
			if root.child[rn] == nil {
				root.child[rn] = &trie{
					child: map[rune]*trie{},
				}
			}
			root = root.child[rn]
			if len(root.words) < 3 {
				root.words = append(root.words, product)
			}
		}
	}

	getWords := func(root *trie, prefix string) [][]string {
		ans := make([][]string, len(prefix))
		for i, rn := range prefix {
			root = root.child[rn]
			if root == nil {
				break
			}
			ans[i] = root.words
		}
		return ans
	}

	root := &trie{
		child: map[rune]*trie{},
	}
	sort.Strings(products)
	for _, product := range products {
		addProduct(root, product)
	}
	return getWords(root, searchWord)
}
