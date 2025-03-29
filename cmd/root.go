package cmd

import (
	"fmt"
	"os"
	"text/template"

	"github.com/0x9n0p/og/generator"
	"github.com/0x9n0p/og/payload"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "og",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		tmpl, err := template.ParseFiles(args[0] + ".tmpl")
		if err != nil {
			return fmt.Errorf("parse template: %w", err)
		}

		var pyld payload.Payload
		if err := pyld.Load(args[0] + ".json"); err != nil {
			return fmt.Errorf("load payload: %w", err)
		}

		generator := generator.Generator{
			Template:    tmpl,
			Payload:     pyld,
			Destination: os.Stdout,
		}

		if err := generator.Generate(cmd.Context()); err != nil {
			return fmt.Errorf("generate: %w", err)
		}

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.og.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
