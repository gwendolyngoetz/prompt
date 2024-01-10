package cmd

import (
	"fmt"
	"gwendolyngoetz/prompt/cmd/prompt"
	"gwendolyngoetz/prompt/cmd/shell"
	"gwendolyngoetz/prompt/cmd/tmux"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Version = "development"

var (
	cfgFile     string
	showVersion bool
)

var rootCmd = &cobra.Command{
	Use:   "prompt",
	Short: "A brief description of your application",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			fmt.Println(Version)
			os.Exit(0)
		}
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

func addCommands() {
	rootCmd.AddCommand(prompt.Command)
	rootCmd.AddCommand(shell.Command)
	rootCmd.AddCommand(tmux.Command)
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $XDG_CONFIG_HOME/prompt/prompt.yaml)")
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version")
	addCommands()
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath("$XDG_CONFIG_HOME/prompt")
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName("prompt")
	}

	viper.AutomaticEnv()
	// viper.SetEnvPrefix("PROMPT")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Unable to load config file:", viper.ConfigFileUsed())
	}
}
