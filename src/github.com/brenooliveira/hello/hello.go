package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Hey Hoo Let`s GO")
	fmt.Println("Gooool do Corinthians às ", time.Now())
	fmt.Println("Random number", rand.Intn(10))
}

func sysout(text string) {
	fmt.Println("Append: ", text)
}
