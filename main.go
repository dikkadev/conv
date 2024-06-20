package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	var (
		printDecimal bool
		printHex     bool
		printOctal   bool
		printBinary  bool
		noPrefix     bool
		noLabel      bool
	)

	pflag.BoolVar(&printDecimal, "dec", false, "Print decimal")
	pflag.BoolVar(&printHex, "hex", false, "Print hexadecimal")
	pflag.BoolVar(&printOctal, "oct", false, "Print octal")
	pflag.BoolVar(&printBinary, "bin", false, "Print binary")
	pflag.BoolVar(&noPrefix, "no-prefix", false, "Do not print prefixes")
	pflag.BoolVar(&noLabel, "no-label", false, "Do not print labels")

	pflag.Parse()

	args := pflag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: conv [options] <number> [<number> ...]")
		os.Exit(1)
	}

	for i, input := range args {
		var value int64
		var err error

		if strings.HasPrefix(input, "0x") {
			value, err = strconv.ParseInt(input[2:], 16, 64)
		} else if strings.HasPrefix(input, "0o") {
			value, err = strconv.ParseInt(input[2:], 8, 64)
		} else if strings.HasPrefix(input, "0b") {
			value, err = strconv.ParseInt(input[2:], 2, 64)
		} else {
			value, err = strconv.ParseInt(input, 10, 64)
		}

		if err != nil {
			fmt.Printf("Invalid number format: %s\n", input)
			os.Exit(1)
		}

		if !printDecimal && !printHex && !printOctal && !printBinary {
			printDecimal, printHex, printOctal, printBinary = true, true, true, true
		}

		if printDecimal {
			printValue(value, 10, "dec", noPrefix, noLabel)
		}
		if printHex {
			printValue(value, 16, "hex", noPrefix, noLabel)
		}
		if printOctal {
			printValue(value, 8, "oct", noPrefix, noLabel)
		}
		if printBinary {
			printValue(value, 2, "bin", noPrefix, noLabel)
		}

		if i < len(args)-1 {
			fmt.Println()
		}
	}
}

func printValue(value int64, base int, prefix string, noPrefix bool, noLabel bool) {
	var result string
	switch base {
	case 10:
		result = strconv.FormatInt(value, 10)
	case 16:
		result = strconv.FormatInt(value, 16)
		if !noPrefix {
			result = "0x" + result
		}
	case 8:
		result = strconv.FormatInt(value, 8)
		if !noPrefix {
			result = "0o" + result
		}
	case 2:
		result = strconv.FormatInt(value, 2)
		if !noPrefix {
			result = "0b" + result
		}
	}

	if noLabel {
		fmt.Println(result)
	} else {
		fmt.Printf("%s: %s\n", prefix, result)
	}
}
