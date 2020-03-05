package Handler

import (
	"fmt"
	"github.com/gorilla/context"
	"github.com/haisum/recaptcha"
	"net/http"
)

func Puzzle(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	url := r.Form.Get("url")
	if url == "" {
		fmt.Fprintf(w, "Please enter url")
		return
	}
	indexDataa.Title = context.Get(r, "title").(string)
	indexDataa.V2Key = context.Get(r, "v2").(string)

	indexDataa.TargetUrl = url

	sitekey := indexDataa.V2Key

	form := fmt.Sprintf(`
		<html>
			<head>
				<script src='https://www.google.com/recaptcha/api.js'></script>
			</head>
 <style>
   div.container6 {
  height: 10em;
  display: flex;
  align-items: center;
  justify-content: center }
div.container6 p {
  margin: 0 }
  </style>
			<body>
				<div class="container6">
				<form action="/v2?url=%s" method="post">
					<div class="g-recaptcha" data-sitekey="%s"></div>
  				    <input type="hidden" id="adreess"  value="%s">
					<input type="submit" value="Enter">
				</form>
				</div>
			</body>
		</html>
	`, url, sitekey, url)
	fmt.Fprintf(w, form)

}
func V2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	private := context.Get(r, "v2private").(string)
	re := recaptcha.R{
		Secret: private,
	}
	isValid := re.Verify(*r)
	if isValid {
		r.ParseForm()
		url := r.Form.Get("url")
		if url == "" {
			fmt.Fprintf(w, "Please enter url")
			return
		}
		http.Redirect(w, r, url, http.StatusSeeOther)
	} else {
		fmt.Fprintf(w, "Not found")
	}
}
