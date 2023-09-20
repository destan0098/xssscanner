package main

import (
	"flag"
	"fmt"
	"github.com/TwiN/go-color"
	"os"
	"strings"

	"github.com/destan0098/xssscanner/Core/GetScan"
	"github.com/destan0098/xssscanner/Core/PostScan"
)

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

	fmt.Println(color.Colorize(color.Green, "* WelCome To Xss Scanner *"))
	fmt.Println(color.Colorize(color.Green, "To Scan Please Enter Like Below "))
	fmt.Println(color.Colorize(color.Red, "This App Just For Training "))
	fmt.Println(color.Colorize(color.Green, `xssscanner -url="https://Test.com" -method="GET" or "POST" -parameter="search"`))
	fmt.Println(color.Colorize(color.Red, `If You have more than 1 Parameter In POST method separate with & `))
	fmt.Println(color.Colorize(color.Red, `For Dumping All Results To a File, add -o Parameter  `))
	fmt.Println(color.Colorize(color.Red, `Like xssscanner -o "outputfile" -url="https://Test.com" -method="GET" or "POST" -parameter="search"`))
	fmt.Println(color.Colorize(color.Yellow, "******************************************************************************"))

	// Define command-line flags for URL, method, parameters, and output file.
	website := flag.String("url", "https://test.com", "Website URL To Scan")
	method := flag.String("method", "GET", "HTTP Method (GET or POST)")
	parameters := flag.String("parameter", "q", "Parameters To Scan")
	fileoutput := flag.String("o", "out.txt", "Output File")

	flag.Parse()

	// Check if a valid website URL is provided.
	if *website == "https://test.com" {
		fmt.Println(color.Colorize(color.Red, `Please Enter Website Address By -url Argument Like -url="https://test.com"`))
		os.Exit(1)
	} else if strings.HasPrefix(*website, "http://") || strings.HasPrefix(*website, "https://") {
		// Start scanning the specified website.
		fmt.Println(color.Colorize(color.Yellow, "[+] Now scanning ") + color.Colorize(color.Yellow, *website) + color.Colorize(color.Yellow, " Website with ") + color.Colorize(color.Green, *method) + color.Colorize(color.Yellow, " Method"))
		fmt.Println(color.Colorize(color.Yellow, "[+] Parameters To Scan -> ") + color.Colorize(color.Yellow, *parameters))

		// Check the HTTP method and perform the appropriate scan.
		if *method == "POST" || *method == "post" {
			parametr := strings.Split(*parameters, "&")
			PostScan.Scan(*website, parametr, *fileoutput)
		} else if *method == "GET" || *method == "get" {
			parametr := strings.Split(*parameters, "&")
			GetScan.Scan(*website, parametr, *fileoutput)
		} else {
			fmt.Println(color.Colorize(color.Red, "[-] Your Method Is Invalid "))
			fmt.Println(color.Colorize(color.Red, "[-] Please Choose POST Or GET "))
		}
	}
}
