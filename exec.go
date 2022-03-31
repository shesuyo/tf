package tf

import (
	"bytes"
	"os/exec"
)

// BashRet RunBashResult
type BashRet struct {
	ErrBuff *bytes.Buffer
	OutBuff *bytes.Buffer
	Err     error
}

// RunBash run bash
func RunBash(command string) BashRet {
	br := BashRet{}
	br.ErrBuff = new(bytes.Buffer)
	br.OutBuff = new(bytes.Buffer)
	cmd := exec.Command("bash", "-c", command)
	cmd.Stderr = br.ErrBuff
	cmd.Stdout = br.OutBuff
	br.Err = cmd.Run()
	return br
}
