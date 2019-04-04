package EasyGoSSH

import "github.com/axgle/mahonia"

const RsaKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQC8fpi06NfVYHAOAnxNMVnTXr/ptsLsNjP+uAt2eO0cc5J9H5XV
8lFVujOrRu/JWi1TDmAvOaf/6A3BphIA1Pwp0AAqlZdwizIum8j0KzpsGYH5qReN
QDwF3oUSKMsQCCGCDHrDYifG/pRi9bN1ZVjEXPr35HJuBe+FQpZTs8DewwIDAQAB
AoGARfNxNknmtx/n1bskZ/01iZRzAge6BLEE0LV6Q4gS7mkRZu/Oyiv39Sl5vUlA
+WdGxLjkBwKNjxGN8Vxw9/ASd8rSsqeAUYIwAeifXrHhj5DBPQT/pDPkeFnp9B1w
C6jo+3AbBQ4/b0ONSIEnCL2xGGglSIAxO17T1ViXp7lzXPECQQDe63nkRdWM0OCb
oaHQPT3E26224maIstrGFUdt9yw3yJf4bOF7TtiPLlLuHsTTge3z+fG6ntC0xG56
1cl37C3ZAkEA2HdVcRGugNp/qmVz4LJTpD+WZKi73PLAO47wDOrYh9Pn2I6fcEH0
CPnggt1ko4ujvGzFTvRH64HXa6aPCv1j+wJBAMQMah3VQPNf/DlDVFEUmw9XeBZg
VHaifX851aEjgXLp6qVj9IYCmLiLsAmVa9rr6P7p8asD418nZlaHUHE0eDkCQQCr
uxis6GMx1Ka971jcJX2X696LoxXPd0KsvXySMupv79yagKPa8mgBiwPjrnK+EPVo
cj6iochA/bSCshP/mwFrAkBHEKPi6V6gb94JinCT7x3weahbdp6bJ6/nzBH/p9VA
HoT1JtwNFhGv9BCjmDydshQHfSWpY9NxlccBKL7ITm8R
-----END RSA PRIVATE KEY-----`

const SHELL_TYPE_BASH = "bash"
const SHELL_TYPE_POWERSHELL = "powershell"

func GBK2Utf8(src string) []byte {
	srcCoder := mahonia.NewDecoder("GB18030")
	srcResult := srcCoder.ConvertString(src)
	tagCoder := mahonia.NewDecoder("utf-8")
	_, cdata, _ := tagCoder.Translate([]byte(srcResult), true)
	return cdata
}

func Utf82GBK(src []byte) []byte {
	enc := mahonia.NewEncoder("GB18030")
	cdata := enc.ConvertString(string(src))
	return []byte(cdata)
}

var TranslateInput = func(src []byte) []byte {
	return src
}

var TranslateOutput = func(src []byte) []byte {
	return src
}
