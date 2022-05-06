package plugin

import (
	"fmt"

	"github.com/spf13/viper"

	"github.com/devstream-io/devstream/pkg/util/log"
)

// Validate a plugin.
// 1. Render template files
// 2. Validate need validate files
func Validate() error {
	name := viper.GetString("name")
	if name == "" {
		return fmt.Errorf("the name must be not \"\", you can specify it by --name flag")
	}
	log.Debugf("Got the name: %s.", name)

	p := NewPlugin(name)

	// 1. Render template files
	files, err := p.RenderTplFiles()
	if err != nil {
		log.Debugf("Failed to render template files: %s.", err)
		return err
	}
	log.Info("Render template files finished.")

	// 2. Validate need validate files
	if err := p.ValidateFiles(files); err != nil {
		log.Debugf("Failed to validate files: %s.", err)
		return err
	}
	log.Info("Validate all files finished.")

	return nil
}