package routeros

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"gopkg.in/routeros.v2"
)

type Dict map[string]interface{}
type Convers struct {
	Regex []string
}

var (
	async  = flag.Bool("async", false, "Use async code")
	useTLS = flag.Bool("tls", false, "Use TLS")
)

func dial(address, username, password string) (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(address, username, password, nil)
	}
	return routeros.Dial(address, username, password)
}

func command(command string, address, username, password string) routeros.Reply {
	flag.Parse()

	c, err := dial(address, username, password)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}

	r, err := c.Run(command)
	if err != nil {
		log.Fatal(err)
	}
	return *r
}

func Loops(commands, user, passwd *string, ip []string) []string {
	// a slice create and for loop and dictorarny append
	var result []string
	var dict []map[string]string
	fmt.Println(dict)
	for _, v := range ip {
		if Router_Nmap(v, "8728") == "open" {

			scx := command(*commands, ""+v+":8728", *user, *passwd)
			for _, v := range scx.Re {
				jsonStr, err := json.Marshal(v)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(jsonStr))
				result = append(result, string(jsonStr))
			}

		} else {
			log.Fatal("Port not working")
		}
	}
	fmt.Println(result)
	return result
}

func Mikregistration(user, passwd *string, ip []string) map[string]string {
	commnd := "/interface/wireless/registration-table/print"
	// var fin_result []string
	var results map[string]string
	for _, IP := range ip {
		if Router_Nmap(IP, "8728") == "open" {
			result := command(commnd, ""+IP+":8728", *user, *passwd)
			results = result.Re[0].Map
		}
	}
	return results
}
