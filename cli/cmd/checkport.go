/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"devops/utils"
	"strings"

	"github.com/spf13/cobra"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		if host == "" {
			cmd.Println("Informe o host a ser validade")
		}

		ports, _ := cmd.Flags().GetString("p")
		if ports == "" {
			cmd.Println("Informe a lista de portas.")
		}

		portList := strings.Split(ports, ",")
		utils.CheckPort(host, portList)

	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)

	checkportCmd.PersistentFlags().String("h", "", "Host a ser verificado")
	checkportCmd.PersistentFlags().String("p", "", "Lista de Portas. Ex: 80,433")

}
