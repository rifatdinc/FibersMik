package routeros

import (
	"flag"
	"fmt"
	"log"
	"rifatdinc/signalspeaker/telegram"
	"strconv"

	"gopkg.in/routeros.v2"
)

type Dict map[string]interface{}
type Convers struct {
	Regex []string
}

var (
	async  = flag.Bool("async", true, "Use async code")
	useTLS = flag.Bool("tls", false, "Use TLS")
)

func dial(address, username, password string) (*routeros.Client, error) {
	if *useTLS {
		return routeros.DialTLS(address, username, password, nil)
	}
	return routeros.Dial(address, username, password)
}

func Client(address, username, password string) {
	flag.Parse()

	c, err := dial(address, username, password)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}

	command(c, address)

}
func command(c *routeros.Client, address string) (bool, error) {

	eth, err := c.RunArgs([]string{"/interface/ethernet/monitor", "=numbers=ether1", "=once"})
	if err != nil {
		log.Fatal(err)
	}
	sysIdentity, err := c.RunArgs([]string{"/system/identity/print"})
	if err != nil {
		log.Fatal(err)
	}
	eth0, err := c.RunArgs([]string{"/interface/ethernet/print"})
	if err != nil {
		log.Fatal(err)
	}
	wirelessInfo, err := c.RunArgs([]string{"/interface/wireless/print"})
	if err != nil {
		log.Fatal(err)
	}
	detailWir, err := c.RunArgs([]string{"/interface/wireless/registration-table/print"})
	if err != nil {
		log.Fatal(err)
	}
	intGetall, err := c.RunArgs([]string{"/interface/getall"})
	if err != nil {
		log.Fatal(err)
	}

	identity := systemIdentity(sysIdentity)
	eth_running := ifEthdisable(eth0)
	interfaceGet(intGetall, address, identity)
	EthDetail(eth, identity, eth_running, address)
	detailWireless(detailWir, wirelessInfo, address, identity)

	return true, nil

}
func systemIdentity(clientName *routeros.Reply) string {
	return clientName.Re[0].Map["name"]
}

func ifEthdisable(eth0 *routeros.Reply) bool {
	fmt.Println()
	var x bool = true
	for _, v := range eth0.Re {
		if v.Map["disabled"] == "true" {
			x = false
		}
		if v.Map["running"] == "false" {
			x = false
		}
	}
	return x
}
func Wirelesz(wirelessInfo *routeros.Reply, address, identity string) {
	for _, v := range wirelessInfo.Re {

		if len(v.Map["tx-chains"]) < 3 {
			telegram.Telegram(address + " " + "Tx chains  are not marked" + " " + identity)
		}
		if len(v.Map["rx-chains"]) < 3 {
			telegram.Telegram(address + " " + "Rx chains  are not marked" + " " + identity)
		}
		if v.Map["channel-width"] != "20/40mhz-Ce" && v.Map["channel-width"] != "20/40mhz-XX" && v.Map["channel-width"] != "20/40mhz-eC" {
			telegram.Telegram(address + " " + "The Channel Setting of the device is 20 MHz" + " " + identity)
		}
		if v.Map["frequency-mode"] != "superchannel" {
			telegram.Telegram(address + " " + "The device is Not in Super Channel Frequency Mode." + " " + identity)
		}
		if v.Map["wireless-protocol"] != "nv2-nstreme-802.11" && v.Map["wireless-protocol"] != "nv2" {
			telegram.Telegram(address + " " + "The device is in 802.11n mode" + " " + identity)
		}
	}
}
func EthDetail(Ethinfo *routeros.Reply, identyt string, eth_running bool, address string) {

	if eth_running {
		for _, x := range Ethinfo.Re {
			if x.Map["Rate"] == "10Mbps" {
				telegram.Telegram(address + " " + "Ethernet is running at 10 mbps" + " " + identyt)
			}
		}
	} else {
		telegram.Telegram(address + " " + "Ethernet is disable or not running " + " " + identyt + " ")
	}
}

func detailWireless(detailWir, wirelessInfo *routeros.Reply, address, identiy string) {
	for _, v := range detailWir.Re {
		tx0, _ := strconv.Atoi(v.Map["tx-signal-strength-ch0"])
		tx1, _ := strconv.Atoi(v.Map["tx-signal-strength-ch1"])
		rx0, _ := strconv.Atoi(v.Map["signal-strength-ch0"])
		rx1, _ := strconv.Atoi(v.Map["signal-strength-ch1"])
		if tx0 < -72 || tx1 < -72 || rx0 < -72 || rx1 < -72 {
			telegram.Telegram(address + " " + identiy + " " + "The signal strength is weak" + " Tx0 =>" + v.Map["tx-signal-strength-ch0"] + " Tx1 =>" + v.Map["tx-signal-strength-ch1"] + " Rx0" + v.Map["signal-strength-ch0"] + " Rx1 =>" + v.Map["signal-strength-ch1"])
		}
		Wirelesz(wirelessInfo, address, identiy)

	}
}

func interfaceGet(intGetall *routeros.Reply, address, identity string) {
	for _, v := range intGetall.Re {
		res, _ := strconv.Atoi(v.Map["link-downs"])
		if res > 100 {
			telegram.Telegram(address + " " + "The interface is very broken" + " " + v.Map["name"] + " " + v.Map["link-downs"] + " " + identity)
		}
	}
}

// func Loops(commands []string, user, passwd *string, ip []string) []string {
// 	// a slice create and for loop and dictorarny append
// 	var result []string
// 	var dict []map[string]string
// 	fmt.Println(dict)
// 	for _, v := range ip {
// 		if Router_Nmap(v, "8728") == "open" {

// 			scx := command(commands, ""+v+":8728", *user, *passwd)
// 			for _, v := range scx.Re {
// 				jsonStr, err := json.Marshal(v)
// 				if err != nil {
// 					fmt.Println(err)
// 				}
// 				fmt.Println(string(jsonStr))
// 				result = append(result, string(jsonStr))
// 			}

// 		} else {
// 			log.Fatal("Port not working")
// 		}
// 	}
// 	return result
// }

// func Mikregistration(user, passwd *string, ip []string) []map[string]string {
// 	var Datarr []map[string]string
// 	for _, IP := range ip {
// 		if Router_Nmap(IP, "8728") == "open" {
// 			Datarr = detailWireless(IP, user, passwd, Datarr)
// 			Datarr = EthDetail(IP, user, passwd, Datarr)
// 			// Wirelesz(IP, user, passwd)
// 		}
// 	}
// 	return Datarr
// }

// func CheckMarked(v map[string]string, IP string, identyt string) bool {
// 	markedTrue := false
// 	if _, ok := v["tx-signal-strength-ch0"]; !ok {
// 		telegram.Telegram(IP + " " + "Tx chains 0 are note marked " + " " + identyt + " ")
// 		markedTrue = true
// 	}
// 	if _, ok := v["tx-signal-strength-ch1"]; !ok {
// 		telegram.Telegram(IP + " " + "Tx chains 1 are note marked " + " " + identyt + " ")
// 		markedTrue = true
// 	}
// 	if _, ok := v["signal-strength-ch0"]; !ok {
// 		telegram.Telegram(IP + " " + "Rx chains 0 are note marked " + " " + identyt + " ")
// 		markedTrue = true
// 	}
// 	if _, ok := v["signal-strength-ch1"]; !ok {
// 		telegram.Telegram(IP + " " + "Rx chains 1 are note marked " + " " + identyt + " ")
// 		markedTrue = true
// 	}
// 	return markedTrue

// }
