/*
Copyright © 2024 shynome <shynome@gmail.com>
*/
package cmd

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/shynome/err0/try"
	"github.com/spf13/cobra"
	"github.com/vscode-lcode/hub/cmd/hub"
)

var args struct {
	addr   string
	domain string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lcode-hub",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, _args []string) {
		if _, err := net.DialTimeout("tcp", args.addr, time.Second); err == nil {
			fmt.Println("服务已启动")
			return
		}
		l := try.To1(net.Listen("tcp", args.addr))
		port := l.Addr().(*net.TCPAddr).Port
		hostTpl := fmt.Sprintf("%%s.%s:%d", args.domain, port)
		hub := hub.New(args.domain, hostTpl)
		http.Serve(l, hub)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.hub.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&args.addr, "listen", "l", "127.0.0.1:4349", "lcode-hub listen addr")
	rootCmd.Flags().StringVar(&args.domain, "domain", "lo.shynome.com", "webdav domain")
}
