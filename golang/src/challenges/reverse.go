package challenges

import "strings"

func Reverse(s string) string {
	list := strings.Split(s, "")

	strReverse := ""
	for _, c := range list {
		strReverse = c + strReverse
	}

	return strReverse
}
