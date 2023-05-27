// ゠゠゠゠゠゠゠゠゠゠ここから追加する゠゠゠゠゠゠゠゠゠゠
package main

import (
	// ゠゠゠゠゠゠゠゠゠゠ここから追加する゠゠゠゠゠゠゠゠゠゠
	"flag"
	// ゠゠゠゠゠゠゠゠゠゠ここまで追加する゠゠゠゠゠゠゠゠゠゠
	"fmt"
	"os"
)

func main() {
	// ゠゠゠゠゠゠゠゠゠゠ここから追加する゠゠゠゠゠゠゠゠゠゠
	flag.Parse()
	arg := flag.Arg(0)
	msg := fmt.Sprintf("Hello %s\n", arg)
	// ゠゠゠゠゠゠゠゠゠゠ここまで追加する゠゠゠゠゠゠゠゠゠゠

	f, err := os.Create("hello.txt") // 変更する
	if err != nil {
		fmt.Printf("failed to create file \n: %v", err)
		return
	}
	defer f.Close() // 追加する

	// ゠゠゠゠゠゠゠゠゠゠ここから追加する゠゠゠゠゠゠゠゠゠゠
	_, err = f.WriteString(msg)
	if err != nil {
		fmt.Printf("failed to write message to file \n: %v", err)
		return
	}
	// ゠゠゠゠゠゠゠゠゠゠ここまで追加する゠゠゠゠゠゠゠゠゠゠

	fmt.Println("Done!")
}