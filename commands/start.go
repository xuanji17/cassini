package commands

import (
	"github.com/spf13/cobra"
)

func addFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&conf, "config", "./config/config.conf", "config file path")
	cmd.Flags().StringVar(&logConf, "log", "./config/log.conf", "log config file path")
}

// NewStartCommand 创建 start/服务启动 命令
func NewStartCommand(run Runner, isKeepRunning bool) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start relay service",
		RunE: func(cmd *cobra.Command, args []string) error {
			return commandRunner(run, isKeepRunning)
		},
	}

	addFlags(cmd)
	return cmd
}
