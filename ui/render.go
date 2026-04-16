// Package ui
package ui

import (
	"fmt"
	"sort"
)

func PrintResult(portsMap map[int]string) {
	keys := make([]int, 0, len(portsMap))
	for k := range portsMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	fmt.Println("PORT", "PROTOCOLE")
	for _, k := range keys {
		fmt.Println(k, portsMap[k])
	}
}
