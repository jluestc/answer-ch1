package main

import(
	"fmt"
	"os"
	"net/http"
	"io"
)
func main(){
	for _,url :=range os.Args[1:]{
		resp,err :=http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
	   _,err1 :=io.Copy(os.Stdout,resp.Body)//return int64,error
		resp.Body.Close()
		if err1 != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err1)
			os.Exit(1)
		} 
	}
}
