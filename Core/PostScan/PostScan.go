package PostScan

import (
	"bufio"
	"crypto/tls"
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

// var path string
var data url.Values
var fo *os.File

func Scan(website string, parameter []string, outputfile bool) {
	Payloadfile := "Core/GetScan/XssPayloads.txt"
	Payload, err := os.Open(Payloadfile)
	if err != nil {
		errormanager(err)
		fmt.Println(color.Colorize(color.Red, "[-] Can't Open Payload File"))
		fmt.Println(color.Colorize(color.Red, "[-] Please Check Payload File From "+Payloadfile))
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
			params := Payl
			if len(parameter) == 1 {
				//	path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s", params)
				data = url.Values{parameter[0]: {params}}
			} else if len(parameter) == 2 {

				//path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s&"+parameter[1]+"=%s", params, params)
				data = url.Values{parameter[0]: {params}, parameter[1]: {params}}
			} else if len(parameter) == 3 {

				//path = fmt.Sprintf(website+"/?"+parameter[0]+"=%s&"+parameter[1]+"=%s", params, params)
				data = url.Values{parameter[0]: {params}, parameter[1]: {params}, parameter[2]: {params}}
			} else {
				fmt.Println(color.Colorize(color.Red, "[-] We Just Accept 2 Parameters"))
				os.Exit(1)
			}

			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{Transport: tr}
			resp, erer := client.PostForm(website, data)

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

				//	var re = regexp.MustCompile(fmt.Sprintf(`(?mi)\W*((?i)%s(?-i))\W*`, Payl))
				var re = regexp.MustCompile(fmt.Sprintf(`(?mi)\W*((?i)%s(?-i))\W*`, Payl3))
				parts := re.FindAllStringSubmatch(content2, -1)

				for i := range parts {

					if i > 0 {
						//	fmt.Println(content2)
						output := fmt.Sprintln(color.Colorize(color.Yellow, "[+] XSS Find with "+Payl+" Payload , Just Run This on your Browser"))
						fmt.Println(output)
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

							outlast := output + "\n" + "*************************************\n"

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
				fmt.Println(color.Colorize(color.Red, "[-] Error In Page Load!!!"))
			}
		}
	}

}
