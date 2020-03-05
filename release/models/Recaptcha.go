package models

type Recaptcha struct {
	PublicKey  string
	PrivateKey string
	Score      float32
}
