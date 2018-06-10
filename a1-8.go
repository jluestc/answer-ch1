package main

import(
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

func main(){
	for _,url :=range os.Args[1:]{
		flag :=strings.HasPrefix(url,"https://")
		if !flag {
			url ="https://"+url
		}
		resp,err :=http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}
		
		b,err :=ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch: reading %s: %v\n",url,err)
			os.Exit(1)
		}
		fmt.Printf("%s",b)
	}
}


