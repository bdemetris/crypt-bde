package main

import (
	"fmt"
	"os"

	"github.com/bdemetris/crypt-bde/bde"
	"github.com/bdemetris/crypt-bde/config"
	"github.com/bdemetris/crypt-bde/crypt"
	"github.com/bdemetris/crypt-bde/version"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func createRotateKeyCmd(conf *config.Config) *cobra.Command {
	var rotateKeyCmd = &cobra.Command{
		Use:   "rotatekey",
		Short: "create a new bitlocker key",
		Long:  `create a new bitlocker key. if a key is already present, replace it`,
		Run: func(cmd *cobra.Command, args []string) {
			if !conf.Loaded() {
				fatal(errors.New("config file not loaded. Must specify --config flag"))
			}
			if err := bde.RotateKey(); err != nil {
				fatal(err)
			}
			if err := crypt.SendCheckin(conf); err != nil {
				fatal(err)
			}
		},
	}

	return rotateKeyCmd
}

func createVersionCmd() *cobra.Command {
	var fFull bool
	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number of crypt-bde",
		Long:  `Print the version number and build information of crypt-bde`,
		Run: func(cmd *cobra.Command, args []string) {
			if fFull {
				version.PrintFull()
				return
			}
			version.Print()
		},
	}

	versionCmd.PersistentFlags().BoolVar(&fFull, "full", false, "print full version information")

	return versionCmd
}

func createRootCmd(conf *config.Config) *cobra.Command {
	// rootCmd represents the base command when called without any subcommands
	var rootCmd = &cobra.Command{
		Use:   "crypt-bde",
		Short: "crypt-bde manages bitlocker on windows",
		Long: `crypt-bde manages local disk encryption, and escrows the secrets to crypt

Complete documentation is available at https://github.com/bdemetris/crypt-bde/.`,
		Run: func(cmd *cobra.Command, args []string) {
			if !conf.Loaded() {
				fatal(errors.New("config file not loaded. Must specify --config flag"))
			}
			if err := crypt.SendCheckin(conf); err != nil {
				fatal(err)
			}
		},
	}

	rootCmd.PersistentFlags().StringP("config", "c", "", "Path to a configuration file")

	return rootCmd
}

func loadConfig(cmd *cobra.Command, conf *config.Config) error {
	configFile := cmd.PersistentFlags().Lookup("config").Value.String()
	if configFile == "" {
		// no config file specified. Must not be required: (example: cryptbde version).
		return nil
	}
	loaded, err := config.New(configFile)
	// copy the value of loaded to conf.
	// because go arguments are passed by value (copied) we cannot replace the entire
	// struct here. Only modify its values.
	*conf = *loaded
	return errors.Wrapf(err, "loading config file %s", configFile)
}

func fatal(err error) {
	fmt.Printf("cryptbde did not complete: %s\n", err)
	os.Exit(1)
}

func main() {
	// create root command and load config.
	conf := new(config.Config)
	rootCmd := createRootCmd(conf)
	rootCmd.ParseFlags(os.Args)
	if err := loadConfig(rootCmd, conf); err != nil {
		fatal(err)
	}

	// add additional commands to root
	rootCmd.AddCommand(createVersionCmd())
	rootCmd.AddCommand(createRotateKeyCmd(conf))

	// run the root command.
	if err := rootCmd.Execute(); err != nil {
		fatal(err)
	}
}
