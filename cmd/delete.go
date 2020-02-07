package cmd

import (
	"fmt"

	"github.com/jimrazmus/dir2consul/files"
	"github.com/jimrazmus/dir2consul/kv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(cliDeleteCmd)
}

var cliDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "deletes kv pairs not found in the files",
	Long:  `deletes blah blah blah`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("the delete command")
		DeleteExtraKeyValues(args)
	},
}

// DeleteExtraKeyValues deletes KV pairs from Consul that do exist on Disk
func DeleteExtraKeyValues(args []string) error {
	// Get Keys from Disk
	diskList := kv.NewList()
	err := files.LoadKeyValuesFromDisk(viper.GetString("dir"), diskList, viper.GetBool("expand"), viper.GetStringSlice("fileTypes"), viper.GetStringSlice("ignoreDirs"))
	if err != nil {
		return err
	}
	diskKeys := diskList.Keys()
	fmt.Println(diskKeys)

	// Get Keys from Consul

	// Subtract Disk Keys from Consul Keys

	// Delete any remaining Consul Keys from Consul

	return nil
}
