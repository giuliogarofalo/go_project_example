package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	fmt.Fprint(w, "This is the About Page")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("2 + +2 is %d", sum))

}

// best practice
//
//	func AddValues(x, y int) (int error) {
//		var sum int
//		sum = x + y
//		return sum nil
//	}
func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// n, err := fmt.Fprintf(w, "Hello, world!")
		// if err != nil {
		// 	fmt.Println(err)

		// }

		fmt.Sprintf(fmt.Sprintf("Number of bytes written: %d", n))

	})

	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(":8080", nil)
}
