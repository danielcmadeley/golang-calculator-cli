package cmd

import (
	"calculator-generator/internal"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	config  internal.CalculatorConfig
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "calculator-generator",
	Short: "Generate customizable Python calculators",
	Long: `Calculator Generator is a CLI tool that creates customizable Python calculators.

You can generate either basic or scientific calculators with various features
including memory, history, plotting, and advanced mathematical functions.

Examples:
  calculator-generator generate --type basic --output calculator.py
  calculator-generator generate --type scientific --features trigonometric,logarithmic
  calculator-generator interactive`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.calculator-generator.yaml)")
	rootCmd.PersistentFlags().StringP("output", "o", "calculator.py", "output file path")
	rootCmd.PersistentFlags().StringP("author", "a", "Calculator Generator", "author name")
	rootCmd.PersistentFlags().Bool("verbose", false, "verbose output")

	// Bind flags to viper
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".calculator-generator" (without extension).
		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".calculator-generator")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil && viper.GetBool("verbose") {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Initialize default config
	config = internal.GetDefaultConfig()
}

// getConfigFromFlags updates config with command line flags
func getConfigFromFlags() {
	config.OutputFile = viper.GetString("output")
	config.Author = viper.GetString("author")
}
