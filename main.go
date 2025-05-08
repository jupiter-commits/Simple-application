package main
import (
"fmt"
"os"
"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func printerFunc(writer http.ResponseWriter, request *http.Request){
	message := `{"text":"Hello world"}`
	fmt.Fprintln(writer, message)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    p, _ := loadPage(title)
    fmt.Fprintf(w, `<h1 style="background-color: yellow;text-transform: capitalize;">%s</h1><div>%s</div>`, p.Title, p.Body)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := os.ReadFile(filename)

	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main(){
	const port = ":81"

	fmt.Println("Starting go app");

	http.HandleFunc("/view/", viewHandler)
	
	http.HandleFunc("/print", printerFunc)
	http.ListenAndServe(port, nil);
}
