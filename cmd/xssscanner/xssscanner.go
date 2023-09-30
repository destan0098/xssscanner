package main

import (
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"github.com/destan0098/xssscanner/pkg/Core/GetScan"
	"github.com/destan0098/xssscanner/pkg/Core/PostScan"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var website, method, fileoutput, parameters string

// //////Develed For Test By Omid Z
func main() {

	fmt.Print(color.Colorize(color.Green, `


 __     __    _____    _____       _____     ____     ____        __      _      __      _    _____   ______    
(_ \   / _)  / ____\  / ____\     / ____\   / ___)   (    )      /  \    / )    /  \    / )  / ___/  (   __ \   
  \ \_/ /   ( (___   ( (___      ( (___    / /       / /\ \     / /\ \  / /    / /\ \  / /  ( (__     ) (__) )  
   \   /     \___ \   \___ \      \___ \  ( (       ( (__) )    ) ) ) ) ) )    ) ) ) ) ) )   ) __)   (    __/   
   / _ \         ) )      ) )         ) ) ( (        )    (    ( ( ( ( ( (    ( ( ( ( ( (   ( (       ) \ \  _  
 _/ / \ \_   ___/ /   ___/ /      ___/ /   \ \___   /  /\  \   / /  \ \/ /    / /  \ \/ /    \ \___  ( ( \ \_)) 
(__/   \__) /____/   /____/      /____/     \____) /__(  )__\ (_/    \__/    (_/    \__/      \____\  )_) \__/  
                                                                                                                


`))
	fmt.Println(color.Colorize(color.Green, ``))
	fmt.Println(color.Colorize(color.Green, "* WelCome To Xss Scanner *"))
	app := &cli.App{
		Name:  "xssscanner",
		Usage: "This Tool Check Website XSS Vulnerability And Just For Training",
		Commands: []*cli.Command{

			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "Scan command",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "url",
						Value:       "https://test.com",
						Usage:       "Url Address to scan",
						Aliases:     []string{"u"},
						Destination: &website,
					},
					&cli.StringFlag{
						Name:        "method",
						Value:       "GET",
						Usage:       "Method to scan(POST or GET)",
						Aliases:     []string{"m"},
						Destination: &method,
					},
					&cli.StringFlag{
						Name:        "parameter",
						Value:       "q",
						Usage:       "Parameter to scan",
						Aliases:     []string{"p"},
						Destination: &parameters,
					},
					&cli.StringFlag{
						Name:        "out",
						Value:       "out.txt",
						Usage:       "Output file address",
						Aliases:     []string{"o"},
						Destination: &fileoutput,
					},
				},
				Action: func(c *cli.Context) error {
					inputFile := c.String("url")
					if inputFile == "https://test.com" {

						fmt.Println(color.Colorize(color.Green, "Please Enter Website Address "))
						err := cli.ShowAppHelp(c)
						if err != nil {
							return err
						}
						return nil

					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	// Define command-line flags for URL, method, parameters, and output file.

	flag.Parse()

	// Check if a valid website URL is provided.
	if website == "https://test.com" {
		fmt.Println(color.Colorize(color.Red, `Please Enter Website Address By scan -url Argument Like scan -url="https://test.com"`))
		os.Exit(1)
	} else if strings.HasPrefix(website, "http://") || strings.HasPrefix(website, "https://") {
		// Start scanning the specified website.
		fmt.Println(color.Colorize(color.Yellow, "[+] Now scanning ") + color.Colorize(color.Yellow, website) + color.Colorize(color.Yellow, " Website with ") + color.Colorize(color.Green, method) + color.Colorize(color.Yellow, " Method"))
		fmt.Println(color.Colorize(color.Yellow, "[+] Parameters To Scan -> ") + color.Colorize(color.Yellow, parameters))

		// Check the HTTP method and perform the appropriate scan.
		if method == "POST" || method == "post" {
			parametr := strings.Split(parameters, "&")
			PostScan.Scan(website, parametr, fileoutput)
		} else if method == "GET" || method == "get" {
			parametr := strings.Split(parameters, "&")
			GetScan.Scan(website, parametr, fileoutput)
		} else {
			fmt.Println(color.Colorize(color.Red, "[-] Your Method Is Invalid "))
			fmt.Println(color.Colorize(color.Red, "[-] Please Choose POST Or GET "))
		}
	}
}
