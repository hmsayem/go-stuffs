package main
import (
	"flag"
	"fmt"
)
func main() {
	output := flag.Bool("output", false, "Should there be output" )
	flag.Parse()
	fmt.Println(*output)
}
