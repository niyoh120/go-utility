package strutil

import "reflect"

func FirstOfNonZero(vs ...interface{}) int {
	for i, v := range vs {
		if reflect.ValueOf(v).IsValid() {
			return i
		}
	}
	return 0
}
func FirstOfNonEmpty(strs ...string) string {
	return strs[FirstOfNonZero(strs)]
}
