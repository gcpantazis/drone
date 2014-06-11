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

		// set the git user and email based on the individual
		// that made the commit.
		f.WriteCmdSilent("git config --global user.name $(git --no-pager log -1 --pretty=format:'%an')")
		f.WriteCmdSilent("git config --global user.email $(git --no-pager log -1 --pretty=format:'%ae')")

		f.WriteCmd("git checkout " + d.Branch)

		// f.WriteCmd("git reset $COMMIT --hard")

		// f.WriteCmd("rm -rf .git")
		// f.WriteCmd("git init")
		// f.WriteCmd("git add -A")		
		// f.WriteCmd("git commit -m 'package build'")

		f.WriteCmd(fmt.Sprintf("git remote add dokku %s", d.Target))
		f.WriteCmd("git push dokku HEAD:master")
	}
}