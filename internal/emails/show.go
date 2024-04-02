package emails

import (
	"fmt"
	"net/http"
)

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Email opened")
}
