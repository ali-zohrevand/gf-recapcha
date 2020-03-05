package Handler

import (
	"fmt"
	"github.com/ali-zohrevand/gf-recapcha/release/models"
	services "github.com/ali-zohrevand/gf-recapcha/release/util"
	"github.com/gorilla/context"
	"gopkg.in/ezzarghili/recaptcha-go.v4"

	"net/http"
	"time"
)

var conf2 models.Conf

type VerifyOption struct {
	Threshold      float32 // ignored in v2 recaptcha
	Action         string  // ignored in v2 recaptcha
	Hostname       string
	ApkPackageName string
	ResponseTime   time.Duration
	RemoteIP       string
}

func Ver3(w http.ResponseWriter, r *http.Request) {
	Private := context.Get(r, "v3private").(string)
	Score := context.Get(r, "score").(float32)
	r.ParseForm()
	token := r.Form.Get("token")

	captcha, err := recaptcha.NewReCAPTCHA(Private, recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	ip := services.GetIpOfRequest(r)
	err = captcha.VerifyWithOptions(token, recaptcha.VerifyOption(VerifyOption{RemoteIP: ip, Threshold: Score}))
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusUnauthorized)
	}
	w.WriteHeader(http.StatusAccepted)

}
