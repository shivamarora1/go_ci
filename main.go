// an anagram of a string is a same string
// but the sequence of characters inside string
// is different
package main

import (
	"fmt"
	"reflect"
)

func main() {
	str1 := "potato"
	str2 := "tatopo"
	fmt.Printf("%v", isAnagram(str1, str2))
}

func isAnagram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	// using has maps for this
	occ1 := getChracterOccurenceMapFromString(str1)
	occ2 := getChracterOccurenceMapFromString(str2)

	return reflect.DeepEqual(occ1, occ2)
}

func getChracterOccurenceMapFromString(str string) map[string]int {
	res := make(map[string]int)
	for _, c := range str {
		if occ, ok := res[string(c)]; ok {
			res[string(c)] = occ + 1
		} else {
			res[string(c)] = 1
		}
	}
	return res
}
