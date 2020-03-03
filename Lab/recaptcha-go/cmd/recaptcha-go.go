package main

import (
	"fmt"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"time"
)

func main() {
	captcha, err := recaptcha.NewReCAPTCHA("6LeVPN4UAAAAAFYYsewKaTq2F8oKZ_pG9sXR8GoI", recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
	fmt.Println(captcha, err)
	captcha.Verify()
}
