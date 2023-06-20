// declare main package
package main
// import libraries 
import  (
"fmt"
"log"
"example.com/greetings"
)
// run main function
// handle error with a non-error value
func main() {
	// set properties of predefined Logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// request a greeting messaging
	message, err := greetings.Hello("")
	// if an error was returned, print it to console and exit
	if err != nil {
		log.Fatal(err)
	}
	// if no error, print message
	fmt.Println(message)
}
// run the app with `go run .`

