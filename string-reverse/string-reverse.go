package main

import (
	"log"
)

func main() {
	str := "Program"

	r := []rune(str)

	for i,j:=0,len(r)-1; i<j; i,j=i+1,j-1 {
		r[i],r[j] = r[j],r[i]
	}

	str = string(r)
	log.Println(str)
}
