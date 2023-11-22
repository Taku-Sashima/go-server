package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello from h1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprint(w, "Hello from h2!\n")
	}

	http.HandleFunc("/", viewHandler)

	http.HandleFunc("/1", h1)
	http.HandleFunc("/h2", h2)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Server Start Up........")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

// helloworldを返すハンドラー, main関数の中で呼び出してる
func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

// HTMLを返すハンドラーを作成
func viewHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("./default/disney.html")

	if err != nil {
		log.Fatal(err)
	}
	if err := html.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
}
