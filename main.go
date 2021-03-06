package main

import (
	"encoding/json"
	"fmt"
	"goweight/pkg"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

var (
	jsonOutput = kingpin.Flag("json", "Output json").Short('j').Bool()
)

func main() {
	kingpin.Version(fmt.Sprintf("%s (%s)", version, commit))
	kingpin.Parse()
	weight := pkg.NewGoWeight()
	work := weight.BuildCurrent()
	modules := weight.Process(work)

	if *jsonOutput {
		m, _ := json.Marshal(modules)
		fmt.Print(string(m))
	} else {
		for _, module := range modules {
			fmt.Printf("%8s %s\n", module.SizeHuman, module.Name)
		}
	}
}
