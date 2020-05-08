package main

import (
	"antalk-go/pkg/apigw"
	"log"
)

func main() {
	c := &apigw.Config{}
	s, err := apigw.NewApigw(c)
	if err != nil {
		log.Fatal("NewApigw error, ", err)
	}
	defer s.Close()
}
