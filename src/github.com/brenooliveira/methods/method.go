package main

import "fmt"

func split(sum int) (x, y int, t string) {
    x = sum * 4 / 9
    y = sum - x
    t = "zapata"
    return
}

func main() {
  fmt.Println(split(1))
}
