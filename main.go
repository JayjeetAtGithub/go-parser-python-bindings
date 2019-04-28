package main
import "C"
import (
 "fmt"
 "os"
 "bufio"
 "github.com/actions/workflow-parser/parser"
)

//export Validate
func Validate(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return false
	}

	reader := bufio.NewReader(f)
	_, err_parse := parser.Parse(reader)
	if err_parse != nil {
		fmt.Println(err_parse)
		f.Close()
		return false
	}
	f.Close()
	return true
}

func main() {}
