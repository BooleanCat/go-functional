package main

import flags "github.com/jessevdk/go-flags"

type PositionalArgs struct {
	TypeName string
}

type Args struct {
	Positional PositionalArgs `positional-args:"yes" required:"yes"`
}

func parseArgs() (Args, error) {
	var args Args
	_, err := flags.Parse(&args)
	if err != nil {
		return Args{}, err
	}
	return args, nil
}
