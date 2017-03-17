/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mainflux/mainflux-cli/cmd"
	"github.com/spf13/cobra"

	"github.com/BurntSushi/toml"
)

// Config struct
type Config struct {
	// LoRa Network Server to send packets to
	UDPHost string
	UDPPort int
}

func main() {

	var s string

	var confFile = "config.toml"
	var conf Config

	conf.UDPHost = "0.0.0.0"
	conf.UDPPort = 1680

	// get current directory location
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	} else {
		confFile = pwd + "/" + confFile
	}

	// print mainflux-cli banner
	fmt.Println(banner)

	////
	// Status
	////
	var cmdStatus = &cobra.Command{
		Use:   "status",
		Short: "Server status",
		Long:  `Mainflux server health checkt.`,
		Run: func(cmdCobra *cobra.Command, args []string) {
			s = cmd.Status()
		},
	}

	var rootCmd = &cobra.Command{
		Use: "maninflux-cli",
		PersistentPreRun: func(cmdCobra *cobra.Command, args []string) {
			// Parse conf file
			if _, err := toml.DecodeFile(confFile, &conf); err != nil {
				fmt.Println("< ! > Using default configuration")
				fmt.Println("< ! > " + err.Error() + "\n")
			}
			// Set UDP server address
			cmd.SetServerAddr(conf.UDPHost, conf.UDPPort)
		},
	}

	// Root Commands
	rootCmd.AddCommand(cmdStatus)

	// Root Flags
	rootCmd.PersistentFlags().StringVarP(
		&confFile, "config", "c", confFile, "Config file path")
	rootCmd.PersistentFlags().StringVarP(
		&conf.UDPHost, "host", "m", conf.UDPHost, "HTTP Host address")
	rootCmd.PersistentFlags().IntVarP(
		&conf.UDPPort, "port", "p", conf.UDPPort, "HTTP Host Port")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}
}

var banner = `
┬  ┌─┐┬─┐┌─┐   ┌─┐┬┌─┌┬┐┌─┐┌─┐┌┐┌
│  │ │├┬┘├─┤───├─┘├┴┐ │ │ ┬├┤ │││
┴─┘└─┘┴└─┴ ┴   ┴  ┴ ┴ ┴ └─┘└─┘┘└┘
`
