package main

import (
	"flag" // 追加する
	"fmt"
)

func main() {
	// ゠゠゠゠゠゠゠゠゠゠ここから変更゠゠゠゠゠゠゠゠゠゠
	flag.Parse()
	arg := flag.Arg(0)
	fmt.Printf("Hello %s\n", arg)
	// ゠゠゠゠゠゠゠゠゠゠ここまで変更゠゠゠゠゠゠゠゠゠゠
}