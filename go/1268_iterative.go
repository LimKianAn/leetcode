// 2022.01.18

package main

import (
	"sort"
)

func suggestedProducts_i(products []string, searchWord string) [][]string {
	sort.Strings(products)

	ans := [][]string{}
	for i := 0; i < len(searchWord); i++ {
		ss := []string{}
		for _, product := range products {
			if len(ss) == 3 {
				break
			}

			if len(product) < i+1 {
				continue // next product might be long enough!
			}

			if product[:i+1] == searchWord[:i+1] {
				ss = append(ss, product)
			}
		}
		ans = append(ans, ss)
	}
	return ans
}
