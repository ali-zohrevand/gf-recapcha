package Handler

import (
	"fmt"
	"github.com/ali-zohrevand/gf-recapcha/release/models"
	"github.com/gorilla/context"
	"html/template"
	"log"
	"net/http"
)

var indexDataa models.IndexData

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./assets/index.html") //parse the html file homepage.html
	if err != nil {                                      // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	r.ParseForm()
	url := r.Form.Get("url")
	if url == "" {
		fmt.Fprintf(w, "Please enter url")
		return
	}
	indexDataa.Title = context.Get(r, "title").(string)
	indexDataa.V3Key = context.Get(r, "v3").(string)
	indexDataa.TargetUrl = url
	err = t.Execute(w, indexDataa) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}
