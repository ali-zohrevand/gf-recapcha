package main

import (
	"fmt"
	"gopkg.in/ezzarghili/recaptcha-go.v4"
	"time"
)

type VerifyOption struct {
	Threshold      float32 // ignored in v2 recaptcha
	Action         string  // ignored in v2 recaptcha
	Hostname       string
	ApkPackageName string
	ResponseTime   time.Duration
	RemoteIP       string
}

func main() {
	captcha, err := recaptcha.NewReCAPTCHA("6LeVPN4UAAAAAFYYsewKaTq2F8oKZ_pG9sXR8GoI", recaptcha.V3, 10*time.Second) // for v3 API use https://g.co/recaptcha/v3 (apperently the same admin UI at the time of writing)
	fmt.Println(captcha, err)
	var token = "03AERD8XqxbFE8gqzNmUsOcWBqQ-rUa5uhWBgwQ98BdhiGGsUBgsXKX9cb9p5rz5xpmUR6jyc-485d8WWFIgq-kg8hoFSrCXp6pd6PqruEu2v4SL499HbRUKwjT6uuyaWRshsVUZTXHm8wxWjLJCbnsDjSeJ8UfIhdt97M4CKrMWMu_FLunp9WECmagBKaWnvJtzAbUF0iIunnBR_EoVo5G31vZvg7RJP0V6BhWdGlAhcQEAT2tlHr4HumqmbfH8ltXE7ZGD_9T-LRE6wokB4lUaXQIEfShQr_KXI3HMbziNdl6gamc-U4KrBxIZkjd6uB1SPUDPyMI23r1T_tWpDAtuz95zC6uM0nKeobSRbFzDmZKGhh1TjaSSGlX83cK6-sPodJsCXHxAVS"
	err = captcha.VerifyWithOptions(token, recaptcha.VerifyOption(VerifyOption{RemoteIP: "198.23.143.100", Threshold: 0.8}))
	fmt.Println(err)
}
