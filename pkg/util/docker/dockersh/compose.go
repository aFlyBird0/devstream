package dockersh

import (
	"bytes"
	"strings"
)

func (op *ShellOperator) ComposeUp() error {
	return ExecInSystemWithParams(".", []string{"docker", "compose", "up", "-d"}, nil, true)
}

func (op *ShellOperator) ComposeDown() error {
	return ExecInSystemWithParams(".", []string{"docker", "compose", "down", "-v"}, nil, true)
}

func (op *ShellOperator) ComposeState() (map[string]interface{}, error) {
	var buf = &bytes.Buffer{}
	err := ExecInSystemWithParams(".", []string{"docker", "compose", "ls"}, buf, true)
	if err != nil {
		return nil, err
	}

	bufStr := buf.String()
	// TODO(daniel-hutao): enhancement is needed
	return map[string]interface{}{
		"Running": strings.Contains(bufStr, "running"),
	}, nil
}
