package morestrings

import (
	"fmt"
	"strings"
)

func Exclaim(s string) string {
	return fmt.Sprintf("%s!!!", strings.ToUpper(s))
}
