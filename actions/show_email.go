package actions

import (
	"fmt"
	"net/http"
)

func ShowEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Email opened")
}
