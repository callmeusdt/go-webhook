package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-webhook/common"
	"go-webhook/utils"
	"os"
	"path/filepath"
)

// systemdCmd represents the systemd command
var systemdCmd = &cobra.Command{
	Use:   "systemd",
	Short: "print systemd configuration content for the webhook service",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value
		host := cmd.Flag("host").Value
		port := cmd.Flag("port").Value
		secret := cmd.Flag("secret").Value.String()
		if secret == "" {
			secret = utils.RandomString(32)
		}
		command := cmd.Flag("command").Value

		binPath, _ := os.Executable()
		binName := filepath.Base(binPath)
		binDir := filepath.Dir(binPath)
		serviceName := fmt.Sprintf("%s-webhook", name)

		fmt.Printf(`# For Debian
# /lib/systemd/system/%s.service
# systemctl enable %s

[Unit]
Description=%s webhook for github
After=network.target

[Service]
Type=simple
ExecStart=%s/%s serve --name %s --host %s --port %s --secret %s --command %s
#Restart=always
User=root
WorkingDirectory=%s
# RuntimeDirectoryMode=2755

[Install]
WantedBy=multi-user.target
Alias=%s.service`,
			serviceName,
			serviceName,
			name,
			binDir, binName, name, host, port, secret, command,
			binDir,
			serviceName)
	},
}

func init() {
	rootCmd.AddCommand(systemdCmd)

	systemdCmd.Flags().StringP("name", "n", "main", "webhook name")
	systemdCmd.Flags().StringP("host", "H", common.ServeHost, "serve host")
	systemdCmd.Flags().IntP("port", "P", common.ServePort, "serve port")
	systemdCmd.Flags().StringP("secret", "s", "", "webhook password")
	systemdCmd.Flags().StringP("command", "c", "", "command to run after trigger")
}
