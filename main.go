package gitcredentialhelper

import (
	"fmt"
	"os"
	"os/exec"
)

/*
Notes about the code comments
- When referred to the "Host" it's the aplication that calles spawns git
- When refrered to the "Client" it's the program created by git
*/

type endRun struct {
	Err error
	Out []byte
}

// Run runs git commands that need credentials
// ask() gets executed when git needs credentials
// The ask argument will be "username" or "password"
func Run(cmd *exec.Cmd, ask func(string) string) ([]byte, error) {
	totalListener++
	currentListener := fmt.Sprintf("%v", totalListener)

	listener := listenerMetaType{AskFunction: ask}
	listenerMeta[currentListener] = listener

	defer delete(listenerMeta, currentListener)

	end := make(chan endRun)
	hostPort := getFreePort()
	serverStartedChan := make(chan struct{})

	go func() {
		ex, err := os.Executable()
		if err != nil {
			ex = os.Args[0]
		}

		cmd.Env = append(
			os.Environ(),
			"LAZYGIT_CLIENT_COMMAND=GET_CREDENTIAL",
			"LAZYGIT_HOST_PORT="+hostPort,
			"LAZYGIT_LISTENER="+currentListener,
			"GIT_ASKPASS="+ex, // tell git where lazygit is located so it can ask lazygit for credentials
			"LANG=en_US.UTF-8",
			"LC_ALL=en_US.UTF-8",
		)
		out, err := cmd.CombinedOutput()
		end <- endRun{
			Err: err,
			Out: out,
		}
	}()

	go runServer(
		serverStartedChan,
		end,
		hostPort,
		currentListener,
	)

	endData := <-end
	return endData.Out, endData.Err
}
