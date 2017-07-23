package main

import (
	"flag"
	"fmt"
)

type FlagArgs struct {
	name      string
	key       string
	value     string
	delete    bool
	deleteAll bool
	metadata  map[string]string
}

func NewFlagArgs() FlagArgs {
	var nameFlag, keyFlag, valueFlag string
	var deleteFlag, allFlag bool
	var metadataFlags MultiFlag

	flag.StringVar(&nameFlag, "name", "", "Name of instance")
	flag.StringVar(&keyFlag, "key", "", "Key of metadata")
	flag.StringVar(&valueFlag, "val", "", "Value of metadata")

	flag.Var(&metadataFlags, "data", "Set metadata in the format: 'key=value'")

	flag.BoolVar(&deleteFlag, "delete", false, "Delete metadata key from an instance (use with -key or -all)")
	flag.BoolVar(&allFlag, "all", false, "Delete all metadata from an instance (use with -delete)")

	flag.Parse()

	return FlagArgs{
		name:      nameFlag,
		key:       keyFlag,
		value:     valueFlag,
		delete:    (deleteFlag && !allFlag),
		deleteAll: (deleteFlag && allFlag),
		metadata:  metadataFlags.Values,
	}
}

func (f FlagArgs) ValidateName() error {
	if f.name == "" {
		return fmt.Errorf("-name flag required\n")
	}
	return nil
}
