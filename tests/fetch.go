package main

import (

	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main(){
	for _, url := range os.Args[1:]{
		resp, _ := http.Get(url)
		b, _ := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Printf("%s", b)
	}
}
