package util

import (
	"fmt"
	"testing"
)

func TestGetConf(t *testing.T) {

	Conf, err := GetConf("conf.yaml")
	fmt.Println(Conf, err)
}
