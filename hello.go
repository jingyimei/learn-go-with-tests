package learn_go_with_tests
//import "fmt"

func hello(s string, l string) string {
	if s == "" {
		s = "world"
	}

	return greetingVerb(l) + s
}

func greetingVerb (l string) (verb string) {
	switch l {
	case "Spanish":
		verb = "Hola "
	case "French" :
		verb = "Bonjour "
	default:
		verb = "Hello "
	}

	return
}

//func main() {
//	fmt.Println(hello())
//}
