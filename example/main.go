package main

import (
	"fmt"

	"github.com/marz619/gbfs"
)

var tobikeshare = "https://tor.publicbikesystem.net/ube/gbfs/v1/"

func main() {
	gbfs := gbfs.New(tobikeshare)

	d, err := gbfs.Discovery()
	if err != nil {
		panic(err)
	}

	fmt.Println(d.Languages())
	for _, l := range d.Languages() {
		for _, f := range d.Feeds(l) {
			fmt.Printf("%s: %s\n", l, f)
		}
	}
}
