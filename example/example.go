package main

import (
	"fmt"

	"github.com/marz619/gbfs-go"
)

var tobikeshare = "https://tor.publicbikesystem.net/ube/gbfs/v1/"

func main() {
	gbfs := gbfs.New(tobikeshare)

	// discover
	d, err := gbfs.GBFS()
	if err != nil {
		panic(err)
	}

	fmt.Println("languages:", d.Languages())

	for _, l := range d.Languages() {
		fmt.Printf("\n%s feeds:\n", l)
		for _, f := range d.IterFeeds(l) {
			fmt.Printf("\t%s\n", f)
		}
	}

	fmt.Println("\nsystem information:")
	// system info in all languages
	for _, l := range d.Languages() {
		si, err := d.SystemInformation(l)
		if err != nil {
			panic(err)
		}

		fmt.Println("\t", l)
		fmt.Println("\t", si.Data.SystemID)
		fmt.Println("\t", si.Data.Name)
		fmt.Println("\t", si.Data.Timezone)
		fmt.Println("\t", si.LastUpdated)
		fmt.Println("\t", si.TTL)
	}
}
