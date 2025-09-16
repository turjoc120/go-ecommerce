package handlers

import (
	"fmt"
	"net/http"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "im inside the test contoroller")
	fmt.Println("im inside the test contoroller")
}
