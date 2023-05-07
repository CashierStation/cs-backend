package util

import (
	"flag"
	"os"
)

func IsFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

func IsProduction() bool {
	return os.Getenv("GIN_MODE") == "release"
}
