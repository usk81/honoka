package commands

import (
    "github.com/spf13/cobra"
    "github.com/YusukeKomatsu/honoka"
    "github.com/davecgh/go-spew/spew"
)

var (
    listCmd = &cobra.Command{
        Use:   "list",
        Short: "Retrive cache index list",
        Long:  "Retrive cache index list (not include cache data). If you get cache, use get method",
        Run:   listCommand,
    }
)

func listCommand(cmd *cobra.Command, args []string) {
    cli, err := honoka.New()
    if err != nil {
        Exit(err)
    }

    list, err := cli.List()
    if err != nil {
        Exit(err)
    }
    spew.Dump(list);
}

func init() {
    RootCmd.AddCommand(listCmd)
}