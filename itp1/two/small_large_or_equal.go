package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	var aAndB = strings.Split(sc.Text(), " ")
	var a, err1 = strconv.Atoi(aAndB[0])
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	var b, err2 = strconv.Atoi(aAndB[1])
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	if a > 1000 || a < -1000 || b > 1000 || b < -1000 {
		fmt.Println("Invalid parameter.")
		return
	}
	if a > b {
		fmt.Println("a > b")
	}
	if a == b {
		fmt.Println("a == b")
	}
	if a < b {
        fmt.Println("a < b")
	}
}
