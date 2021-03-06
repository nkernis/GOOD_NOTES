// Copyright © 2018 NAME HERE nkernis, alexnovak 
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var noteDir string
const DEFAULT_NOTEPATH string = ".good_notes"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gn",
	Short: "Use to take many a good note",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initNoteDir)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.GOOD_NOTES.yaml)")
	rootCmd.PersistentFlags().StringVar(&noteDir, "note-dir", "", "note directory (default is $HOME/.gnote)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".GOOD_NOTES" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".GOOD_NOTES")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func initNoteDir() {
	if noteDir == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		path := home + "/" + DEFAULT_NOTEPATH
		if _, err := os.Stat(path); os.IsNotExist(err) {
			filemode := os.FileMode(448) // Translates to 0700	
			err := os.Mkdir(path, filemode)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		} else {
			finfo, err := os.Lstat(path)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			if !finfo.Mode().IsDir() {
				fmt.Printf(`ERROR: A file exists at the default GOOD_NOTES directory
location: %s, and no other directory was specified. Exiting
`, path)
				os.Exit(1)
			}
		}
	}
}
