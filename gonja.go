package gojinja2

import (
	"github.com/jmolinski/gosible-templates/config"
	"github.com/jmolinski/gosible-templates/exec"
	"github.com/jmolinski/gosible-templates/loaders"
)

var (
	// DefaultLoader is being used by the DefaultSet.
	DefaultLoader = loaders.MustNewFileSystemLoader("")

	// DefaultEnv is an environment created for quick/standalone template rendering.
	DefaultEnv = NewEnvironment(config.DefaultConfig, DefaultLoader)

	// Methods on the default set
	FromString = DefaultEnv.FromString
	FromBytes  = DefaultEnv.FromBytes
	FromFile   = DefaultEnv.FromFile
	FromCache  = DefaultEnv.FromCache

	// Globals for the default set
	Globals = DefaultEnv.Globals
)

// Must panics, if a Template couldn't successfully parsed. This is how you
// would use it:
//     var baseTemplate = gonja.Must(gonja.FromFile("templates/base.html"))
func Must(tpl *exec.Template, err error) *exec.Template {
	if err != nil {
		panic(err)
	}
	return tpl
}
