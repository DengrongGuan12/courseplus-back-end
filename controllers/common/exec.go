package common

import (
	"encoding/json"
	"fmt"
	"gogs.mebox.wiki/course-plus/courseplus-back-end/controllers/base"
	"os/exec"
)

type ExecController struct {
	base.BaseController
}

type GitAuth struct {
	Secret string
}

// @router /execGitSh [post]
func (c *ExecController) ExecGitSh() {
	// c.Data["data"] = "sdfsdf"
	// c.ServeJSON()
	var gitAuth GitAuth
	tmpSecret := "sdfsk234@ksdfdjf"
	json.Unmarshal(c.Ctx.Input.RequestBody, &gitAuth)
	if gitAuth.Secret == tmpSecret {
		// fmt.Println("auth success")
		cmd := exec.Command("/bin/sh", "/var/www/go_workspace/src/gogs.mebox.wiki/course-plus/courseplus-back-end/git.sh")
		_, err := cmd.Output()
		if err != nil {
			return
		}
		// fmt.Println(string(bytes))
	} else {
		// fmt.Println("auth fail")
	}
}

// @router /execWhoAmI
func (c *ExecController) ExecWhoAnI() {
	cmd := exec.Command("whoami")
	bytes, err := cmd.Output()
	if err != nil {
		fmt.Println("cmd.Output: ", err)
		return
	}

	fmt.Println(string(bytes))
}
