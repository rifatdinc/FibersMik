package macadres

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	g "github.com/gosnmp/gosnmp"
)

func RequestMac(mac string) string {
	fmt.Println("Requesting mac: " + mac)
	host := "https://api.macvendors.com/" + mac
	resp, err := http.Get(host)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	// snmp("10.50.254.253")
	return string(body)

}

func Snmp(ip string) string {
	g.Default.Target = "10.50.254.253"
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
	defer g.Default.Conn.Close()

	oids := []string{"1.3.6.1.2.1.2.2.1.6.2"}
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}
	macAdres := ""
	for _, variable := range result.Variables {

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		// fmt.Println(variable.Value)
		fmt.Printf("string: %s", string(variable.Value.([]byte)))
		for _, value := range variable.Value.([]byte) {
			fmt.Println(value)
			macAdres += ":" + string(byte(value))
		}
		// macAdres = strconv.Itoa(int(variable.Value))
		// fmt.Println(reflect.TypeOf(variable.Value))
		// fmt.Println(variable.Type.String())
		// switch variable.Type {
		// case g.OctetString:
		// 	fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		// default:
		// 	// ... or often you're just interested in numeric values.
		// 	// ToBigInt() will return the Value as a BigInt, for plugging
		// 	// into your calculations.
		// 	fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		// }
	}
	fmt.Println(macAdres)
	return macAdres
}
