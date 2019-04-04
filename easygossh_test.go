package EasyGoSSH

import "testing"

func TestStartSSH(t *testing.T) {
	err := StartSSH("127.0.0.1:2222", "admin", "admin", SHELL_TYPE_POWERSHELL)
	if err != nil {
		println(err.Error())
	}
}
