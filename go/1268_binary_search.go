// 2022.01.18

package main

import (
	"sort"
)

func suggestedProducts_b(products []string, searchWord string) [][]string {
	sort.Strings(products)

	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	ans := [][]string{}
	rns := []rune{}
	for _, rn := range searchWord {
		rns = append(rns, rn)
		l := sort.Search(len(products), func(j int) bool { // left:=lower bound
			return products[j] >= string(rns)
		})
		r := sort.Search(len(products), func(j int) bool { // right:=upper bound
			return products[j] > string(append(rns, 'z'+1)) // so that product[j] > all rns*, * any English alphabet
		})
		if l == r {
			break
		}
		ans = append(ans, products[l:min(l+3, r)])
	}

	for len(ans) < len(searchWord) {
		ans = append(ans, []string{})
	}
	return ans
}
