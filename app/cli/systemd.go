package cli

import (
	"fmt"
	"hook/app"
)

func DumpSystemd() {
	name := fmt.Sprintf("%s_webhook", app.OPTIONS.Config)

	fmt.Println("# For Debian")
	fmt.Printf("# /lib/systemd/system/%s.service\n", name)
	fmt.Printf("# systemctl enable %s\n\n", name)
	fmt.Printf("[Unit]\nDescription=%s for github\nAfter=network.target\n\n"+
		"[Service]\nType=simple\nExecStart=/opt/webhook/hook -c %s\n#Restart=always\nUser=root\nWorkingDirectory=/opt/webhook\n# RuntimeDirectoryMode=2755\n\n"+
		"[Install]\nWantedBy=multi-user.target\nAlias=%s.service\n", name, app.OPTIONS.Config, name)

}
