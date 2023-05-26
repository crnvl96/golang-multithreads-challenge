package main

import (
	"fmt"
	"os"
	"time"

	callapicep "github.com/crnvl96/go-channels/callAPICEP"
	callviacep "github.com/crnvl96/go-channels/callViaCEP"
)

func main() {
	apiCEPCh := make(chan callapicep.ApiCEP)
	viaCEPCh := make(chan callviacep.ViaCEP)

	go func() {
		result, err := callapicep.Call("14015-110")
		if err != nil {
			fmt.Printf("Error calling API CEP: %v", err)
			os.Exit(1)
		}
		apiCEPCh <- result
	}()

	go func() {
		result, err := callviacep.Call("14015110")
		if err != nil {
			fmt.Printf("Error calling VIA CEP: %v", err)
			os.Exit(1)
		}
		viaCEPCh <- result
	}()

	select {
	case r := <-apiCEPCh:
		fmt.Printf("API CEP: %v", r)
	case r := <-viaCEPCh:
		fmt.Printf("VIA CEP: %v", r)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
