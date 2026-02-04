package main

import (
	"fmt"
	"log"
	"os"
)
func main(){
	file,err := os.Open("./messages.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer file.Close()
	for{
		buff := make([]byte,8)
		n,err := file.Read(buff)
		if err != nil{
			break
		}
		fmt.Printf("read: %s\n",string(buff[:n]))
	}
}
