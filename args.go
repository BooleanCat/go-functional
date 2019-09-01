package main

import flags "github.com/jessevdk/go-flags"

type PositionalArgs struct {
	TypeName string
}

type Args struct {
	ImportPath string         `short:"i" long:"import-path"`
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

func isErrHelp(err error) bool {
	if err == nil {
		return false
	}

	flagErr, ok := err.(*flags.Error)
	if !ok {
		return false
	}

	return flagErr.Type == flags.ErrHelp
}
