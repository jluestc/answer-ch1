package main

import(
	"fmt"
	"os"
)

func main(){
	for n,arg :=range os.Args[1:]{
		fmt.Printf("%v\t%v\n",n,arg)
	}
}
