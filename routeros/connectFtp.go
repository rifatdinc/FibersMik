package routeros

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

type ConnectFtp struct {
	Host    string
	Port    int
	User    string
	Pass    string
	Timeout int
	Debug   bool
}

func Connected_Ftp() []string {
	// Connect Ftp to RouterOS
	c, err := ftp.Dial("10.50.254.253:21", ftp.DialWithTimeout(5*time.Second))

	if err != nil {
		log.Fatal(err)
	}
	err = c.Login("admin", "mc4152")
	if err != nil {
		log.Fatal(err)
	}
	r, err := c.Retr("rifatDinc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	buf, err := ioutil.ReadAll(r)
	fmt.Println(buf, err)
	buff_result := string(buf)
	if stringInSlice("Poyraz", []string{buff_result}) || stringInSlice("poyraz", []string{buff_result}) {
		return []string{buff_result}
	} else {
		return []string{"Poyrazwifi Isminde bir access point yok"}
	}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
