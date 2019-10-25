package cmd

import (
  "os"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	rootCmd = &cobra.Command{
		Use:   "MonitoringBund",
		Short: "Universal client/server for MonitoringBund",
		Long: `MonitoringBund - is a universal application for building distributed metrics collection
and processing system.`,
	}
)

// Execute executes the root command.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&bctx.CfgFile, "config", "", "config file (default is $HOME/.zabbix-bund)")
	rootCmd.PersistentFlags().StringVarP(&bctx.Logverbose, "verbose", "v", "info", "Level for the logging (trace,debug,warning,info,fatal)")

}

func initConfig() {
	if bctx.CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(bctx.CfgFile)
	} else {
		// Find home directory.
    viper.SetConfigName("MonitoringBund")
    viper.AddConfigPath("/etc/")
    viper.AddConfigPath("$HOME/.monitoringbund")
    viper.AddConfigPath(".")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		log.Debug(fmt.Sprintf("Using config file:", viper.ConfigFileUsed()))
	}
  bund_log.Init_Log(bctx.Logverbose, bctx.Logoutput)
  is_raft, _ := rootCmd.PersistentFlags().GetBool("is_raft")
  is_rest, _ := rootCmd.PersistentFlags().GetBool("is_rest")
  log.Debug("Called commandpath", rootCmd.Name())
  bund.Init_Internal_Components(is_raft, is_rest)
  log.Debug("root_init process complete")
}
