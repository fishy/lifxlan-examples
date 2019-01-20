package main

import (
	"flag"
	"time"

	"github.com/fishy/go-flagutils"
	"github.com/fishy/lifxlan"
)

var (
	noack = flag.Bool(
		"noack",
		false,
		"Do not require ack for drawing API calls.",
	)

	discoverTimeout = flag.Duration(
		"discoverTimeout",
		time.Second*2,
		"Timeout for discover API calls.",
	)

	drawTimeout = flag.Duration(
		"drawTimeout",
		time.Millisecond*200,
		"Timeout for drawing API calls.",
	)

	interval = flag.Duration(
		"interval",
		time.Millisecond*1500,
		"Interval between 2 frames.",
	)

	broadcastHost = flag.String(
		"broadcastHost",
		"",
		`Broadcast IP (e.g. "192.168.1.255"). Empty value means "255.255.255.255", which should work in most networks.`,
	)

	generations = flag.Int(
		"generations",
		20,
		"Number of generations before stop. 0 means never stop (but still stops when the board is empty).",
	)

	loop = flag.Bool(
		"loop",
		false,
		"After reached the number of generations, reset the board and loop over instead of reverting to the original colors.",
	)

	kelvin = flag.Int(
		"kelvin",
		4000,
		"The Kelvin value of the color, in range of [2500, 9000].",
	)

	turnon = flag.Bool(
		"turnon",
		false,
		"Turn on the device if it's not already on, and turn it off afterwards (when no loop args specified).",
	)

	origColor = &flagutils.RGB{
		R: 0xff,
		G: 0xff,
		B: 0xff,
	}

	target lifxlan.Target
)

func init() {
	flag.Var(
		&target,
		"target",
		"The MAC address of the target tile device. Empty value means any (first) tile device.",
	)
	flag.Var(
		origColor,
		"color",
		"The hex color to use, in format of `\"rrggbb\"`.",
	)
	flag.Parse()
}
