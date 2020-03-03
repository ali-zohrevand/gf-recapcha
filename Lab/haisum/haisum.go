package main

//https://github.com/haisum/recaptcha/blob/master/example/main.go
import (
	"fmt"
	"github.com/haisum/recaptcha"
	"log"
	"net/http"
)

func main() {
	sitekey := "6LcCSN4UAAAAAEE0bX4d3GSFaZ7z3qVANzgEIwV5"
	re := recaptcha.R{
		Secret: "6LcCSN4UAAAAAISdM6pN-Jr9k-IyroIjkA65woqH",
	}

	form := fmt.Sprintf(`
		<html>
			<head>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
			<body>
				<form action="/submit" method="post">
					<div class="g-recaptcha" data-sitekey="%s"></div>
					<input type="submit">
				</form>
			</body>
		</html>
	`, sitekey)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, form)
	})
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		isValid := re.Verify(*r)
		if isValid {
			fmt.Fprintf(w, "Valid")
		} else {
			fmt.Fprintf(w, "Invalid! These errors ocurred: %v", re.LastError())
		}
	})

	log.Printf("\n Starting server on http://localhost:8100 . Check example by opening this url in browser.\n")

	err := http.ListenAndServe(":8100", nil)

	if err != nil {
		log.Fatalf("Could not start server. %s", err)
	}
}
