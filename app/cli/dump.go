package cli

import (
	"fmt"
	"hook/app"
)

func Dump() {
	fmt.Println("# For Debian")
	fmt.Printf("# config file name should be './%s.yaml'\n", app.OPTIONS.Config)
	fmt.Printf("app:\n"+
		"  name: \"%s\"\n"+
		"  version: 0.1\n"+
		"  port: 58080\n"+
		"  secret: \"helloGithub\"\n\n"+
		"commands:\n"+
		"  - name: \"/usr/bin/git\"\n"+
		"    args:\n"+
		"      - \"pull\"\n", app.OPTIONS.Config)
}
