package gitcredentialhelper

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	cmd := exec.Command("git", "push")
	out, err := Run(cmd, func(question string) string {
		fmt.Println("git asked for:", question)
		return ""
	})
	assert.NoError(t, err)
	assert.Equal(t, out, "")
}
