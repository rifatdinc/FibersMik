package routeros

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func Envrioment() map[string]string {
	nasuser := os.Getenv("nasuser")
	naspass := os.Getenv("naspassword")

	return map[string]string{"nasuser": nasuser, "naspasswd": naspass}
}

func Nas(address string) {
	username := Envrioment()["nasuser"]
	password := Envrioment()["naspasswd"]
	c, err := dial(address, username, password)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer c.Close()

	if *async {
		c.Async()
	}
	res, err := c.Run("/ppp/active/print")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	wg.Add(len(res.Re))

	for i, v := range res.Re {
		fmt.Println(v.Map["address"], "##Disi", i)

		go func(i int, v map[string]string) {
			fmt.Println(v["address"], "coroutine ici", i)
			defer wg.Done()
			if Router_Nmap(v["address"], "8728") == "open" {
				Client(v["address"]+":8728", username, "password")
			}
		}(i, v.Map)

	}
	wg.Wait()
	fmt.Println("End---------------")
}
