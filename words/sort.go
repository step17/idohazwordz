package words

import (
	"sort"
	"strings"
)

func Sort(letters string) string {
	rs := strings.Split(letters, "")
	sort.Sort(sort.StringSlice(rs))
	return strings.Join(rs, "")
}
