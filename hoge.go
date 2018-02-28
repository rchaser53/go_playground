package main

import (
	"regexp"
	"strings"
)

func main() {
	targetStr := "AbcHoge"
	re := regexp.MustCompile("[A-Z]")
	nyan := re.FindAllStringIndex(targetStr, -1)

	var result string

	for i, r := range targetStr {
		earlyReturnFlg := false
		if i == 0 {
			result = result + strings.ToLower(string(r))
			continue
		}

		for _, a := range nyan {
			if i == a[0] {
				result = result + "_" + strings.ToLower(string(r))
				earlyReturnFlg = true
			}
		}

		if earlyReturnFlg == true {
			continue
		}

		result = result + string(r)
	}

	println(result)
}
