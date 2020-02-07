package cmd

import (
	"fmt"

	"github.com/jimrazmus/dir2consul/files"
	"github.com/jimrazmus/dir2consul/kv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(cliAddCmd)
}

var cliAddCmd = &cobra.Command{
	Use:   "add",
	Short: "only adds new kv pairs",
	Long:  `adds blah blah blah`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add command")
		AddNewKeyValues(args)
	},
}

// AddNewKeyValues inserts KV pairs from disk that do not already exist in Consul
func AddNewKeyValues(args []string) error {
	// Get Keys from Disk
	diskList := kv.NewList()
	err := files.LoadKeyValuesFromDisk(viper.GetString("dir"), diskList, viper.GetBool("expand"), viper.GetStringSlice("fileTypes"), viper.GetStringSlice("ignoreDirs"))
	if err != nil {
		return err
	}
	diskKeys := diskList.Keys()
	fmt.Println(diskKeys)

	// Get Keys from Consul

	// Subtract Consul Keys from Disk Keys

	// Add any remaining Disk Keys to Consul

	return nil
}
