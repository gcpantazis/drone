package deploy

import (
	"os"
	"fmt"
	"github.com/drone/drone/pkg/build/buildfile"
)

type Dokku struct {
	Target string `yaml:"target,omitempty"`
	Force  bool   `yaml:"force,omitempty"`
	Branch string `yaml:"branch,omitempty"`
}

func (d *Dokku) Write(f *buildfile.Buildfile) {

	if d.Branch == os.Getenv("DRONE_BRANCH") {
	
		f.WriteCmdSilent("COMMIT=$(git rev-parse HEAD)")
		f.WriteCmd("git checkout " + d.Branch)

		f.WriteCmd(fmt.Sprintf("git remote add dokku %s", d.Target))
		f.WriteCmd("git push dokku HEAD:master")
	}
}