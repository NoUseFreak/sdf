package utils

import (
	"regexp"
)

func ReSubMatchMap(r *regexp.Regexp, str string) map[string]string {
	match := r.FindStringSubmatch(str)
	subMatchMap := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if name != "" && i != 0 && len(match) > i {
			subMatchMap[name] = match[i]
		}
	}

	return subMatchMap
}
