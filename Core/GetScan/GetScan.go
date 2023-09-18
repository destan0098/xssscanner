package GetScan

import (
	"bufio"
	"fmt"
	"github.com/TwiN/go-color"

	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"
)

func errormanager(err error) {
	fmt.Println("[-] " + err.Error())
}

var fo *os.File
var path string

func Scan(website string, parameter []string, outputfile bool) {
	Payloadfile := "Core/GetScan/XssPayloads.txt"
	Payload, err := os.Open(Payloadfile)
	if err != nil {
		errormanager(err)
		fmt.Println(color.Colorize(color.Red, "[-] Can't Open Payload File"))
		fmt.Println(color.Colorize(color.Red, "[-] Please Check Payload File From "+Payloadfile))
		fmt.Println(color.Colorize(color.BlueBackground, "[+] We Create XssPayload.txt File"))
		fmt.Println(color.Colorize(color.BlueBackground, "[+] Please Put your xss Payload in this File"))
		fo, err = os.OpenFile(Payloadfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {

			errormanager(err)
			fmt.Println(color.Colorize(color.Red, "[-] Error In Permission To Open File"))
		}
		os.Exit(1)
	} else {

		PayloadBuf := bufio.NewScanner(Payload)
		PayloadBuf.Split(bufio.ScanLines)
		for PayloadBuf.Scan() {
			Payl := PayloadBuf.Text()
			params := url.QueryEscape(Payl)
			if strings.HasSuffix(website, "&") {
				if len(parameter) == 1 {
					path = fmt.Sprintf(website+parameter[0]+"=%s", params)
				} else if len(parameter) == 2 {
					path = fmt.Sprintf(website+parameter[0]+"=%s&"+parameter[1]+"=%s", params, params)
				} else if len(parameter) == 3 {
					path = fmt.Sprintf(website+parameter[0]+"=%s&"+parameter[1]+"=%s&"+parameter[2]+"=%s", params, params, params)
				} else {
					fmt.Println(color.Colorize(color.Red, "[-] We Just Accept 2 Parameters"))
					os.Exit(1)
				}

			} else {
				if len(parameter) == 1 {
					path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s", params)
				} else if len(parameter) == 2 {
					path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s&"+parameter[1]+"=%s", params, params)
				} else if len(parameter) == 3 {

					path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s&"+parameter[1]+"=%s&"+parameter[2]+"=%s", params, params, params)
				} else {
					fmt.Println(color.Colorize(color.Red, "[-] We Just Accept 2 Parameters"))
					os.Exit(1)
				}

			}

			//
			resp, erer := http.Get(path)
			if erer != nil {
				errormanager(erer)
				continue
			}
			if resp.StatusCode == 200 {
				var Payl2, Payl3 string
				Payl2 = strings.ReplaceAll(Payl, "(", "[(]")
				Payl3 = strings.ReplaceAll(Payl2, ")", "[)]")

				content, _ := ioutil.ReadAll(resp.Body)
				content2 := string(content)

				var re = regexp.MustCompile(fmt.Sprintf(`(?mi)\W*((?i)%s(?-i))\W*`, Payl3))
				//	var re = regexp.MustCompile(`(?mi)\W*((?i)` + Payl3 + `(?-i))\W*`)
				parts := re.FindAllStringSubmatch(content2, -1)

				for i := range parts {

					if i > 0 {
						output := "[+] XSS Find with " + Payl
						output2 := fmt.Sprintln(color.Colorize(color.Yellow, output))
						fmt.Println(output2)
						fmt.Println(color.Colorize(color.Yellow, "[+] Click Below Link To Test This "))
						outputlink := fmt.Sprintln(path)
						fmt.Println(outputlink)

						fmt.Println(color.Colorize(color.Green, "******************************************************"))

						if outputfile {
							currentTime := time.Now()
							name := fmt.Sprintf("%v", currentTime.Format("2006-01-02-15-01"))

							fo, err = os.OpenFile(name+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
							if err != nil {

								errormanager(err)
								fmt.Println(color.Colorize(color.Red, "[-] Error In Permission To Open File"))
							}
							defer func(fo *os.File) {
								err := fo.Close()
								if err != nil {

								}
							}(fo)

							outlast := output + "\n" + outputlink + "\n" + "*******************************************************************\n"

							_, err = fmt.Fprint(fo, outlast)
							if err != nil {
								errormanager(err)
							}
						}
						break
					}

				}

			} else if resp.StatusCode == 403 {
				fmt.Println(color.Colorize(color.Red, "[-] WAF Blocked This Payload"))

			} else {
				fmt.Println(color.Colorize(color.Red, resp.StatusCode))
				//	fmt.Println(path)
				fmt.Println(color.Colorize(color.Red, "[-] Error In Page Load!!!"))
			}
		}
	}

}
