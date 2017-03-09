package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp := r.PostFormValue("answer")
	fmt.Print(resp)
	t, err := template.ParseFiles("C:\\Work\\testWork\\templates\\index.html", "C:\\Work\\testWork\\templates\\header.html") //"C:\\Work\\testWork\\templates\\footer.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())

	}

	t.ExecuteTemplate(w, "index", nil)

}

//test
func main() {
	fmt.Println("listening on port :3000")

	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(":3000", nil)

}
