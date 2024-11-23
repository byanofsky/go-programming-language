package echo

import "strings"

func EchoInefficient(args []string) string {
	var s, sep string
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func EchoJoin(args []string) string {
	return strings.Join(args, " ")
}
