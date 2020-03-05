package main

import (
	"fmt"
	captcha "github.com/netdata/go-recaptcha"
)

func main() {
	client := captcha.NewClient("6LeVPN4UAAAAAKceUc88CvyFxc52-unnA2lAfVyN")
	token := "03AERD8XpT-P2dcOlFE7Y64AXJlPFo7vvezlNH6b9aEnZdOTjNShw4FI8CjhInuIrBLA6CT3AQHiTLYEtBqZkmZncgmzVdyTkCk4d3MWIzUZvuLUCQEAkYlYDu-k2CKcuWtsz5USNxOiT-FClQf5Rzln2SkKifoPXziAvzvb_gL78ojMGyiqLChmU2Y7-R2h2iKvRFtn1vbzaxNK_H2Q51ppVr8kFVMXgCAIgYfVuQF3s09IQvgBIT4VsSzdBz9hgTISk38LPZd3DoOIFLIs6B1SD3cCQFMBK1Z6-Xo-M7FZMCKVAaNBPSBTnZ2ZsC_Mal3fcwN0lFppA-ocX1SktJjflGQUNvl0w1zduQJ6q9o8gu0KZL0IzL9p49I06JaITq9RLa78N0QdNl"
	r, err := client.Score(token, "198.23.143.100")
	fmt.Println(err)
	fmt.Println(r)
}
