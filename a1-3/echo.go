package echo

import(
	"fmt"
	"strings"
)

func Echo1(args []string){
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func Echo2(args []string){
	fmt.Println(strings.Join(args[0:], " "))
}
