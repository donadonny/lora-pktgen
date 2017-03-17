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

	"github.com/mainflux/lora-pktgen/cmd"
	"github.com/spf13/cobra"
)

func main() {

	var s string

	UDPHost := "0.0.0.0"
	UDPPort := 1680

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
			// Set UDP server address
			cmd.SetServerAddr(UDPHost, UDPPort)
		},
	}

	// Root Commands
	rootCmd.AddCommand(cmdStatus)

	// Root Flags
	rootCmd.PersistentFlags().StringVarP(
		&UDPHost, "host", "m", UDPHost, "UDP Host address")
	rootCmd.PersistentFlags().IntVarP(
		&UDPPort, "port", "p", UDPPort, "UDP Host Port")

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
