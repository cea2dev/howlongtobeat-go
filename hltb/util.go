package hltb

import "strings"

func splitStrTerms(s, sep string) []string {
	platforms := strings.Split(s, sep)

	res := []string{}

	for i := 0; i < len(platforms); i++ {
		value := strings.TrimSpace(platforms[i])

		if value != "" {
			res = append(res, value)
		}
	}

	return res
}
