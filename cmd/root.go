// Copyright © 2018 mritd <mritd1234@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"github.com/mitchellh/go-homedir"
	"github.com/mritd/gitflow-toolkit/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "gitflow-toolkit",
	Short: "Git Flow 辅助工具",
	Long: `
一个用于 CI/CD 实施的 Git Flow 辅助工具`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// init config
	cobra.OnInitialize(initConfig)

	// add flags
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gitflow-toolkit.yaml)")

	// add sub cmd
	RootCmd.AddCommand(NewCi())
	RootCmd.AddCommand(NewCm())
	RootCmd.AddCommand(NewXMr())
	RootCmd.AddCommand(NewFeat())
	RootCmd.AddCommand(NewFix())
	RootCmd.AddCommand(NewDocs())
	RootCmd.AddCommand(NewStyle())
	RootCmd.AddCommand(NewRefactor())
	RootCmd.AddCommand(NewPerf())
	RootCmd.AddCommand(NewHotFix())
	RootCmd.AddCommand(NewTest())
	RootCmd.AddCommand(NewChore())
	RootCmd.AddCommand(NewInstall())
	RootCmd.AddCommand(NewUninstall())
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		util.CheckAndExit(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".gitflow-toolkit")
	}
	viper.AutomaticEnv()
	viper.ReadInConfig()
}
