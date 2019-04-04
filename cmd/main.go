package main

import EasyGoSSH ".."

func main() {
	err := EasyGoSSH.StartSSH("127.0.0.1:2222", "admin", "admin", EasyGoSSH.SHELL_TYPE_POWERSHELL)
	if err != nil {
		println(err.Error())
	}

}
