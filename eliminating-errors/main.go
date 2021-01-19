package eliminating_errors

//// inspired by https://dave.cheney.net/2019/01/27/eliminate-error-handling-by-eliminating-errors
//import (
//	"bufio"
//	"os"
//	"log"
//	"fmt"
//	"io"
//)
//
//func main() {
//	fd, err := os.Open("in.txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	n, err := doRead(fd)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(n)
//}
//
//func doRead(r io.Reader) (int, error) {
//	var count int
//	s := bufio.NewScanner(r)
//
//	for s.Scan() {
//		count++
//	}
//
//	return count, s.Err()
//}
