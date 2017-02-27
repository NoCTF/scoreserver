package main

import (
	"log"
	"os"

	flags "github.com/jessevdk/go-flags"
)

type options struct {
	Dir string `short:"d" long:"directory" required:"true" description:"directory name to process"`
}

func main() {
	os.Exit(_main())
}

func _main() int {
	var opts options
	if _, err := flags.Parse(&opts); err != nil {
		log.Printf("%s", err)
		return 1
	}

	p := processor{}
	p.Dir = opts.Dir
	if err := p.Do(); err != nil {
		log.Printf("%s", err)
		return 1
	}
	return 0
}
