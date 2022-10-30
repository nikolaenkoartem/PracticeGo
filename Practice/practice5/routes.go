package practice5

import (
	"fmt"
	"net/http"
)

func Mux(port string) *http.ServeMux {
	fmt.Printf("Порт %v занят сервером", port)

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/func1", functia1)
	mux.HandleFunc("/func2", functia2)
	mux.HandleFunc("/func2/{inp}", productsHandler)

	return mux
}