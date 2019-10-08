package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	args := os.Args

	usage := "voer je volledige naam in: "
	uitkomst := "de generator heeft deze pincode voor u bedacht: "

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	v, a := args[1], args[2]

	now := time.Now()
	pincode := now.Unix()

	rand.Seed(pincode)
	x := rand.Intn(100000)
	fmt.Println(uitkomst, x, "bedankt voor het gebruiken van dit programma", v, a)

}
