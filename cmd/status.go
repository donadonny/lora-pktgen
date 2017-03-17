/**
 * Copyright (c) 2016 Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package cmd

// Status
func Status() string {
	// Exit on error
	if UDPConnErr != nil {
		return "Cannot connect to LoRa Network Server - UDP connection error."
	}

	return "OK"
}
