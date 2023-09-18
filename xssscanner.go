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
	fmt.Println(color.Colorize(color.Green, `xssscanner -url="https://Test.com" -method="GET" or "POST" -parameter="seaech"`))
	fmt.Println(color.Colorize(color.Red, `If You have more than 1 Parameter In POST method separate with & `))
	fmt.Println(color.Colorize(color.Red, `For Dump All Result To File add -o Parameter  `))
	fmt.Println(color.Colorize(color.Red, `Like xssscanner -o "outputfile" -url="https://Test.com" -method="GET" or "POST" -parameter="seaech"`))

	fmt.Println(color.Colorize(color.Yellow, "******************************************************************************"))

	website := flag.String("url", "https://test.com", "WebSite Url To Scan")
	method := flag.String("method", "GET", "Method To Test")
	parameters := flag.String("parameter", "q", "Parameters To Scan")
	fileoutput := flag.String("o", "out.txt", "Parameters To Save Into File")
	flag.Parse()

	if *website == "https://test.com" {
		fmt.Println(color.Colorize(color.Red, `Please Enter Website Address By -url Argument Like -url="https://test.com"`))
		os.Exit(1)
	} else if strings.HasPrefix(*website, "http://") || strings.HasPrefix(*website, "https://") {
		fmt.Println(color.Colorize(color.Yellow, "[+] You Now scanning ") + color.Colorize(color.Yellow, *website) + color.Colorize(color.Yellow, " WebSite With ") + color.Colorize(color.Green, *method) + color.Colorize(color.Yellow, " Method"))
		fmt.Println(color.Colorize(color.Yellow, "[+] Parameters To Scan -> ") + color.Colorize(color.Yellow, *parameters))
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
