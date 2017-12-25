package main

import (
	"fmt"
	"github.com/frankh/rai-vanity/address"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "RaiBlocks Vanity Generator"
	app.Usage = "Generate wallet seeds with desirable public addresses"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "prefix, p",
			Usage: "Prefix to search for at the start of address",
		},
		cli.IntFlag{
			Name:  "count, n",
			Value: 1,
			Usage: "Number of valid addresses to generate before exiting, or 0 for infinite (default=1).",
		},
	}
	app.Action = func(c *cli.Context) {
		fmt.Println("Estimated number of iterations needed:", address.EstimatedIterations(c.String("prefix")))
		for i := 0; i < c.Int("count") || c.Int("count") == 0; i++ {
			seed, addr, err := address.GenerateVanityAddress(c.String("prefix"))
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
			fmt.Printf(`Found matching address!
Seed: %s
Address: %s

`, strings.ToUpper(seed), addr)
		}
	}
	app.Run(os.Args)
}
