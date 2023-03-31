package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	hurufPertama := true
	for _, v := range data {
		isInputAvailable := strings.Contains(v, input)

		if isInputAvailable {
			if hurufPertama {
				result += v
				hurufPertama = false
			} else {
				result += "," + v 
			}
		}
	}

	

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
