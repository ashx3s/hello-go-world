// declare main package
package main
// import libraries 
import  (
"fmt"
"example.com/greetings"
)
// run main function
func main() {
	// print Hello Gladys
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
// run the app with `go run .`

