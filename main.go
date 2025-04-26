package main

import (
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello, World!</h1>"))
}

func main() {
<<<<<<< HEAD
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
=======
    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"  
    }
>>>>>>> 981f8df0fe2d4bc9692a9716c430dcd9ecae3f1b

	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)

	println("Launching", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		println("Server error:", err.Error())
	}

	println("Session is end")
}
