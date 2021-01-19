package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a := fmt.Sprintf(`"%x"`, "1234")
	b := `"` + hex.EncodeToString([]byte("1234")) + `"`
	fmt.Println(a)
	fmt.Println(b)
}
