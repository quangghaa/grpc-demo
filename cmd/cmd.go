package cmd

import (
	api "github.com/quangghaa/grpc-demo/server/demo"
	"github.com/quangghaa/grpc-demo/service/demo"
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
		ps := demo.NewPingService()
		ps.Start(port)
	},
}

var RegisterCmd = &cobra.Command{
	Use:   "register",
	Short: "Start register service",
	Run: func(cmd *cobra.Command, args []string) {
		// port, _ := cmd.Flags().GetInt("port")
		// demo.Start(port)
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

var CheckCmd = &cobra.Command{
	Use:   "check",
	Short: "Check current connections",
	Run: func(cmd *cobra.Command, args []string) {
		// connObj := demo.NewConnectionService()
		// connObj.Check()
	},
}

var ScanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and remove connections",
	Run: func(cmd *cobra.Command, args []string) {
		// connObj := demo.NewConnectionService()
		// connObj.Scan()
	},
}

func init() {
	PingCmd.Flags().Int("port", 8002, "port")
	ApiCmd.Flags().Int("port", 8080, "port")

	//add command
	RootCmd.AddCommand(PingCmd)
	RootCmd.AddCommand(ApiCmd)
	RootCmd.AddCommand(CheckCmd)
	RootCmd.AddCommand(ScanCmd)
}
