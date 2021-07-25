package main

import (
	"fmt"
	"homework2/servive"
	"net/http"
)

func main() {
	http.HandleFunc("/homework2", servive.TestService)
	fmt.Println(http.ListenAndServe(":3000",nil))
}
