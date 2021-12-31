package main

import (
	"fmt"
	"os"

	consulApi "github.com/hashicorp/consul/api"
)

func main() {

	/*	nomadClient, err := nomadApi.NewClient(&nomadApi.Config{})
		if err != nil {
			fmt.Println("Failed to connect to Nomad")
			os.Exit(1)
		}
		nomadQueryOpts := nomadApi.QueryOptions{}
		spew.Dump(nomadClient.Allocations().List(&nomadQueryOpts))
	*/
	consulClient, err := consulApi.NewClient(&consulApi.Config{})
	if err != nil {
		fmt.Println("Failed to connect to Nomad")
		os.Exit(1)
	}
	serviceNames := getConsulServices(consulClient)
	for svcName, tags := range serviceNames {
		fmt.Println(svcName, tags)
	}
	//spew.Dump()
}

func getConsulServices(consulClient *consulApi.Client) map[string][]string {
	consulQueryOpts := consulApi.QueryOptions{}
	serviceNames, _, err := consulClient.Catalog().Services(&consulQueryOpts)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return serviceNames
}
