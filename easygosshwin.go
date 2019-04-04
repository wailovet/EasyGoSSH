package EasyGoSSH

func RunningOnWin() {
	TranslateInput = func(src []byte) []byte {
		return Utf82GBK(src)
	}
	TranslateOutput = func(src []byte) []byte {
		return GBK2Utf8(string(src))
	}
}
