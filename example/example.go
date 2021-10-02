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

	fmt.Println("last updated:", d.LastUpdatedRFC3339())
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

		fmt.Println("\tlanguage:", l)
		fmt.Println("\tsystem ID:", si.Data.SystemID)
		fmt.Println("\tsystem name:", si.Data.Name)
		fmt.Println("\tsystem timezone:", si.Data.Timezone)
		fmt.Println("\tsystem last updated:", si.LastUpdated)
		fmt.Println("\tsystem TTL", si.TTL)
	}
}
