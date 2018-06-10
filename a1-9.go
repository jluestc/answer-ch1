package main

import(
	"fmt"
	"os"
	"net/http"
	"strings"
	"io/ioutil"
)

func main(){
	for _,url :=range os.Args[1:]{
		if flag :=strings.HasPrefix(url,"https://");!flag{
			url ="https://"+url
		}
		resp,err :=http.Get(url)
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch: %v\n",err)
			os.Exit(1)
		}

		b,err:=ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err!=nil{
			fmt.Fprintf(os.Stderr,"fetch: reading %s %v\n", url,err)
			os.Exit(1)
		}
		fmt.Printf("%s\n",b)
		a :=resp.Status
		fmt.Printf("%v\n",a)
	}
}

