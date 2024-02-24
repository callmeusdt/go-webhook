package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-webhook/common"
	"go-webhook/http"
	"go-webhook/http/router"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "run a webhook server",
	Run: func(cmd *cobra.Command, args []string) {
		name := cmd.Flag("name").Value
		http.InitLogger(name.String())

		host := cmd.Flag("host").Value
		port := cmd.Flag("port").Value
		common.GithubSecret = cmd.Flag("secret").Value.String()
		common.HookCommand = cmd.Flag("command").Value.String()

		gin.SetMode(gin.ReleaseMode)
		http.SERVER = gin.Default()

		engine := router.SetupRouter()

		// 启动 HTTP 服务器
		log.Fatal(
			engine.Run(fmt.Sprintf("%s:%s", host, port)).Error(),
		)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringP("name", "n", "main", "webhook name")
	serveCmd.Flags().StringP("host", "H", common.ServeHost, "serve host")
	serveCmd.Flags().IntP("port", "P", common.ServePort, "serve port")
	serveCmd.Flags().StringP("secret", "s", "", "webhook password")
	serveCmd.Flags().StringP("command", "c", "", "command to run after trigger")
}
