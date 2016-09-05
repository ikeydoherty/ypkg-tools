package cmd

import "flag"

// CMD - structure for each of the supported commands for ytools
type CMD struct {
	Flags *flag.FlagSet
	Run   func()
}
