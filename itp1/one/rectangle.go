package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var stdin = bufio.NewScanner(os.Stdin)
	stdin.Scan()
	var line = stdin.Text()
	var aAndB = strings.Split(line, " ")
	var a, _ = strconv.Atoi(aAndB[0])
	var b, _ = strconv.Atoi(aAndB[1])
    var area = a * b
    var length = a * 2 + b * 2
    fmt.Println(area, length)
}
