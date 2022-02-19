package main

import (
	"log"
	"net/http"
)

func main() {

	//TODO

	port := ":8080"

	router := http.NewServeMux()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	})

	log.Fatal(http.ListenAndServe(port, router))

}

// func main() {

// 	port := ":8080"

// 	router := http.NewServeMux()

// 	// router.HandleFunc("/", CreateCaca("dura"))

// 	log.Fatal(http.ListenAndServe(port, router))
// }

// func Get(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "GET")
// }

// type HttpCaca func(w http.ResponseWriter, r *http.Request)

// func CreateCaca(s string) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "desde Crear Caca %s", s)
// 	}
// }

// func GetCaca(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Get Caca")
// }

// type CacaMiddleWare func(http.HandlerFunc) http.HandlerFunc

// func CacaMiddleware(f http.HandlerFunc, middlewares ...CacaMiddleWare) http.HandlerFunc {
// 	for _, m := range middlewares {
// 		f = m(f)
// 	}
// 	return f
// }

// func LoggingMiddleware(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		log.Println("LoggingMiddleware")
// 		f(w, r)
// 	}
// }

// func RecoveryMiddleware(f http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		defer func() {
// 			if err := recover(); err != nil {
// 				log.Println("RecoveryMiddleware", err)
// 			}
// 		}()
// 		f(w, r)
// 	}
// }
