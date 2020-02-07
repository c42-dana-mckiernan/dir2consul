package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "dir2consul",
		Short: "manages Consul Key-Value pairs from files",
		Long:  "a longer description",
	}
)

// Execute ...
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetEnvPrefix("dir2consul")
	viper.SetDefault("consulPrefix", "axiomatic/dir2consul")
	viper.SetDefault("dir", ".")
	viper.SetDefault("expandFiles", false)
	viper.SetDefault("fileTypes", []string{".hcl", ".ini", ".json", ".properties", ".toml", ".txt", ".yml", ".yaml"})
	viper.SetDefault("ignoreDirs", []string{".git"})
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
