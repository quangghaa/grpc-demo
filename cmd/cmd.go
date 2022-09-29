package cmd

import (
	api "github.com/quangghaa/grpc-demo/server/register"
	pingService "github.com/quangghaa/grpc-demo/service/ping"
	registerService "github.com/quangghaa/grpc-demo/service/register"
	"github.com/spf13/cobra"
)

// ServeCmd represents the serve command
var RootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "Start ping service",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		pingService.Start(port)
	},
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Start register service",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		registerService.Start(port)
	},
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "Start api service",
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		api.GatewayServer(port)
	},
}

func init() {
	PingCmd.Flags().Int("port", 8001, "port")
	RegisterCmd.Flags().Int("port", 8002, "port")
	ApiCmd.Flags().Int("port", 8080, "port")

	//add command
	RootCmd.AddCommand(PingCmd)
	RootCmd.AddCommand(RegisterCmd)
	RootCmd.AddCommand(ApiCmd)
}
