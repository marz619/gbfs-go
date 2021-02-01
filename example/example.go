package main

import (
	"fmt"

	"github.com/marz619/gbfs-go"
)

var tobikeshare = "https://tor.publicbikesystem.net/ube/gbfs/v1/"

func main() {
	gbfs := gbfs.New(tobikeshare)

	// discover
	d, err := gbfs.Discover()
	if err != nil {
		panic(err)
	}

	fmt.Println("languages:", d.Languages())

	for _, l := range d.Languages() {
		for _, f := range d.Feeds(l) {
			fmt.Printf("%s: %s\n", l, f)
		}
	}

	// system info
	for _, l := range d.Languages() {
		for _, f := range d.Feeds(l) {
			if f.Name == "system_information" {
				sysInfo, err := gbfs.SystemInformation(f.URL.String())
				if err != nil {
					panic(err)
				}
				fmt.Println(l, f.Name, sysInfo.Data.SystemID)
			}
		}
	}
}
