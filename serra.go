// Package main provides a typing test
package main

import (
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/noqqe/serra/src/serra"
)

//[--count:1 <card>:[mmq/1] <set>:<nil> add:true cards:false remove:false set:false sets:false stats:false update:false]

var opts struct {
	Add     bool     `docopt:"add"`
	Remove  bool     `docopt:"remove"`
	Cards   bool     `docopt:"cards"`
	Set     bool     `docopt:"set"`
	Sets    bool     `docopt:"sets"`
	Stats   bool     `docopt:"stats"`
	Update  bool     `docopt:"update"`
	Card    []string `docopt:"<card>"`
	SetCode string   `docopt:"<setcode>"`
	Count   int64    `docopt:"--count"`
}

// Main Loop
func main() {

	usage := `Serra

Usage:
  serra add <card>... [--count=<number>]
  serra remove <card>...
  serra cards
  serra set <setcode>
  serra sets
  serra update
  serra stats

Options:
  -h --help					Show this screen.
	--count=<number>	Count of card to add.  [default: 1].
  --version					Show version.
`

	args, _ := docopt.ParseDoc(usage)
	err := args.Bind(&opts)
	if err != nil {
		fmt.Println(err)
	}

	if opts.Add {
		serra.Add(opts.Card, opts.Count)
	} else if opts.Remove {
		serra.Remove(opts.Card)
	} else if opts.Cards {
		serra.Cards()
	} else if opts.Sets {
		serra.Sets()
	} else if opts.Set {
		serra.ShowSet(opts.SetCode)
	} else if opts.Update {
		serra.Update()
	} else if opts.Stats {
		serra.Stats()
	}

}
