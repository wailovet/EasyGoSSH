package EasyGoSSH

import (
	"github.com/gliderlabs/ssh"
	"github.com/pkg/sftp"
	"log"
	"os/exec"
	"runtime"
)

func ListenSSH(addr string, username string, password string, shellPath string) error {
	sshHandler := func(s ssh.Session) {

		println("test2")
		serverOptions := []sftp.ServerOption{
			func(server *sftp.Server) error {
				println("test")
				return nil
			},
		}
		sftps, _ := sftp.NewServer(s, serverOptions...)
		sftps.Serve()
		log.Println("new session:", s.User())

		cmd := exec.Command(shellPath)
		stdIn, _ := cmd.StdinPipe()
		stdoutIn, _ := cmd.StdoutPipe()
		stderrIn, _ := cmd.StderrPipe()

		err := cmd.Start()
		if err != nil {
			return
		}
		go func() {
			var stderrInBuf = make([]byte, 1024)
			for {
				i, err := stderrIn.Read(stderrInBuf)
				if err != nil {
					log.Println("stderrInBuf:", err.Error())
					break
				} else {
					_, err = s.Write(TranslateOutput(stderrInBuf[:i]))
					if err != nil {
						log.Println("stderrInBuf:", err.Error())
						break
					}
				}
			}
		}()

		go func() {
			var stdoutInBuf = make([]byte, 1024)
			for {
				i, err := stdoutIn.Read(stdoutInBuf)
				if err != nil {
					log.Println("stdoutInBuf:", err.Error())
					break
				} else {
					_, err = s.Write(TranslateOutput(stdoutInBuf[:i]))
					if err != nil {
						log.Println("stdoutInBuf:", err.Error())
						break
					}
				}
			}
		}()
		go func() {
			var buf = make([]byte, 1024)
			for {
				i, err := s.Read(buf)
				if err != nil {
					cmd.Process.Kill()
					break
				}

				_, err = stdIn.Write(TranslateInput(buf[:i]))
				println("input:", i)
				println("input:", string(buf[:i]))
				if err != nil {
					cmd.Process.Kill()
					break
				}
			}
		}()

		println(cmd.Wait().Error())
	}

	passHandler := func(ctx ssh.Context, pass string) bool {
		return ctx.User() == username && pass == password
	}

	s := &ssh.Server{
		Addr:            addr,
		Handler:         sshHandler,
		PasswordHandler: passHandler,
	}

	err := ssh.HostKeyPEM([]byte(RsaKey))(s)
	if err != nil {
		return err
	}

	return s.ListenAndServe()
}

func StartSSH(addr string, username string, password string, shellPath string) error {
	if runtime.GOOS == "windows" {
		RunningOnWin()
	}
	return ListenSSH(addr, username, password, shellPath)
}
