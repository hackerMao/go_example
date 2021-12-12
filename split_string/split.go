package split_string

import "strings"

func Split(s string, sep string) []string {
	var ret []string

	index := strings.Index(s, sep)
	for index >= 0 {
		ret = append(ret, s[:index])
		s = s[index+1:]
 		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return ret
}
