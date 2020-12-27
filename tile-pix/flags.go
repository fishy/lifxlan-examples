package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"go.yhsif.com/flagutils"
	"go.yhsif.com/lifxlan"
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

	kelvin = flag.Int(
		"kelvin",
		8000,
		"The Kelvin value of the color, in range of [2500, 9000].",
	)

	noskip = flag.Bool(
		"noskip",
		false,
		"Do not skip failed frames, retry them after the next interval.",
	)

	turnon = flag.Bool(
		"turnon",
		false,
		"Turn on the device if it's not already on, and turn it off afterwards when neither -loop nor -still are set.",
	)

	loop  flagutils.OneOf
	still flagutils.OneOf

	target lifxlan.Target
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(
			flag.CommandLine.Output(),
			"Usage:\n\tcat path/to/image.jpg | %s [args]\n\nArgs:\n",
			os.Args[0],
		)
		flag.PrintDefaults()
	}

	flag.Var(
		&loop,
		"loop",
		"After fully shown the picture, loop over instead of reverting to the original colors. Unsets -still.",
	)
	flag.Var(
		&still,
		"still",
		"Shrink the picture to fully fit inside the tile boundaries and display still instead of scrolling. Unsets -loop, ignores -interval.",
	)
	flagutils.GroupOneOf(&loop, &still)

	flag.Var(
		&target,
		"target",
		"The MAC address of the target tile device. Empty value means any (first) tile device.",
	)

	flag.Parse()
}
