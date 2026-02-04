package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)
func main(){
	file,err := os.Open("./messages.txt")
	if err != nil{
		log.Fatal(err)
	}
	linescanner := bufio.NewScanner(file)
	for linescanner.Scan(){
		fmt.Printf("read: %s\n",linescanner.Text())
	}
	if err := linescanner.Err() ;err != nil{
		log.Fatal(err)
	}
}
