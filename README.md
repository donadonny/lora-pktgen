# LoRaWAN Packet Generator (GW Simulator)

[![License](https://img.shields.io/badge/license-Apache%20v2.0-blue.svg)](LICENSE)
[![Build Status](https://travis-ci.org/mainflux/mainflux-cli.svg?branch=master)](https://travis-ci.org/mainflux/mainflux-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/Mainflux/mainflux-cli)](https://goreportcard.com/report/github.com/Mainflux/mainflux-cli)
[![Join the chat at https://gitter.im/Mainflux/mainflux](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

LoRaWAN packet generator is a command line tool for generation of UDP packets that can be sent from the PC host to the LoRa Network Server. It simulates LoRaWAN gateway and sends the UDP packages defined by the "Gateway to Server Interface Protocol" defined in [Semtech document ANNWS.01.2.1.W.SYS](https://www.google.fr/url?sa=t&rct=j&q=&esrc=s&source=web&cd=1&cad=rja&uact=8&ved=0ahUKEwjg44KhsN7SAhXpYZoKHVduAJEQFggaMAA&url=https%3A%2F%2Fwww.thethingsnetwork.org%2Fforum%2Fuploads%2Fdefault%2Foriginal%2F1X%2F4fbda86583605f4aa24dcedaab874ca5a1572825.pdf&usg=AFQjCNFfztbcaVB002yqLD3393nCDuJiaA&sig2=jkHvDwmrzKg7ePCSM25UOA&bvm=bv.149760088,d.bGs). Protocol is also described [here](https://github.com/Lora-net/packet_forwarder/blob/master/PROTOCOL.TXT).

Basically, `lora-pktgen` acts as a LoRa node + LoRa GW, i.e. it generates a package from a user-defined payload and adds service information in JSON that are injected by Packet Forwarder during the packet traversing through GW. This way LoRa Network Server get a properly shaped UDP packet, like the one that would come from a real HW.

`lora-pktgen` is used to test the LoRa Network Server deployments and integrations in the absence of expensive HW (gateways and nodes) and tedious HW network set-up.


### Installation
#### Prerequisite
If not set already, please set your `GOPATH` and `GOBIN` environment variables. For example:
```bash
mkdir -p ~/go
export GOPATH=~/go
export GOBIN=$GOPATH/bin
# It's often useful to add $GOBIN to $PATH
export PATH=$PATH:$GOBIN
```

#### Get the code
Use [`go`](https://golang.org/cmd/go/) tool to "get" (i.e. fetch and build) `lora-pktgen` package:
```bash
go get github.com/mainflux/lora-pktgen
```

This will download the code to `$GOPATH/src/github.com/mainflux/lora-pktgen` directory,
and then compile it and install the binary in `$GOBIN` directory.

Now you can run the program with:
```
mainflux-cli
```
if `$GOBIN` is in `$PATH` (otherwise use `$GOBIN/lora-pktgen`)

### Documentation
Development documentation can be found [here](http://mainflux.io/).

### Community
#### Mailing lists
[mainflux](https://groups.google.com/forum/#!forum/mainflux) Google group.

#### IRC
[Mainflux Gitter](https://gitter.im/Mainflux/mainflux?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

#### Twitter
[@mainflux](https://twitter.com/mainflux)

#### Authors
This tool is crafted by @drasko and @manuIO for the benefit of mankind.

### License
[Apache License, version 2.0](LICENSE)
