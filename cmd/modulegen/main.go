package main

import (
	"flag"
	"fmt"
	"os"

	"gin-starter/cmd/modulegen/gen"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gen [handler|repository|service] --schema=... --version=... --entity=...")
		os.Exit(1)
	}

	target := os.Args[1]
	fs := flag.NewFlagSet(target, flag.ExitOnError)
	schema := fs.String("schema", "", "Schema name")
	version := fs.String("version", "", "Version (e.g. v1)")
	entity := fs.String("entity", "", "Entity name")
	fs.Parse(os.Args[2:])

	switch target {
	case "handler":
		if err := gen.GenerateHandlers(*schema, *version, *entity); err != nil {
			panic(err)
		}
	case "repository":
		if err := gen.GenerateRepositories(*schema, *version, *entity); err != nil {
			panic(err)
		}
	case "service":
		if err := gen.GenerateServices(*schema, *version, *entity); err != nil {
			panic(err)
		}
	default:
		fmt.Println("Unknown target:", target)
		os.Exit(1)
	}
}
