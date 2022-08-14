package kubectl

import (
	"os/exec"
	"strings"

	"github.com/devstream-io/devstream/pkg/util/log"
)

const (
	Create string = "create"
	Apply  string = "apply"
	Delete string = "delete"
)

// KubeCreate runs "kubectl create -f filename"
func KubeCreate(filename string) error {
	return kubectlAction(Create, filename)
}

// KubeApply runs "kubectl apply -f filename"
func KubeApply(filename string) error {
	return kubectlAction(Apply, filename)
}

// KubeDelete runs "kubectl delete -f filename"
func KubeDelete(filename string) error {
	return kubectlAction(Delete, filename)
}

func kubectlAction(action string, filename string) error {
	cmd := exec.Command("kubectl", action, "-f", filename)
	cOut, err := cmd.CombinedOutput()
	if err != nil {
		log.Errorf("Failed to exec: < %s >.", cmd.String())
		log.Errorf("Exec logs: < %s >. Got error: %s.", string(cOut), err)
		return err
	}
	log.Info(strings.TrimSuffix(string(cOut), "\n"))
	return nil
}
