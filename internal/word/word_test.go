package word

import (
	"log"
	"testing"
)

func TestUnderscoreToUpperCamelCase(t *testing.T) {
	s := UnderscoreToUpperCamelCase("tao_rui_cheng")
	if s != "TaoRuiCheng" {
		log.Fatalln("格式转换错误")
	}
}
