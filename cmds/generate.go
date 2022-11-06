package cmds

import (
	"fmt"
	"strings"

	"github.com/jurgen-kluft/go-pass/clipboard"
	"github.com/jurgen-kluft/go-pass/password"
)

type GenerateCmd struct {
	Symbols  int     `optional:"" short:"s" name:"symbols" default:"2" help:"Password should contain 'n' symbols"`
	Numbers  int     `optional:"" short:"n" name:"numbers" default:"2" help:"Password should contain 'n' numbers"`
	Mixed    bool    `optional:"" short:"m" help:"Password may contain lower and upper case characters"`
	Unique   bool    `optional:"" short:"u" help:"Password may only contain unique characters (no repeat)"`
	Clip     bool    `optional:"" short:"c" help:"Copy generated password to clipboard"`
	InPlace  bool    `optional:"" short:"i" name:"in-place" help:"Set generated password as new password for 'pass-name'"`
	Form     *string `optional:"" short:"f" help:"The form to use, e.g. 'xxnx-xn-xnxn' (a=alpha, x=lower, X=upper, n=number, s=symbol, ingores arg 'length')"`
	Length   int     `optional:"" short:"l" default:"10" help:"The length of the password to generate"`
	PassName string  `arg:"" optional:"" help:"If provided, set the generated password as the new password for 'pass-name'"`
}

func (a *GenerateCmd) Run(globals *Globals) error {

	// If form is supplied determine the length of the password from the following allowed characters:
	// - 'x' for lowercase letters
	// - 'X' for uppercase letters
	// - 'n' for numbers
	// - 'a' for letter or number
	// - 's' for symbols
	// The following characters are allowed as separators:
	// - '-' for a dash
	// - '_' for an underscore

	if a.Form != nil {

		numAlphaNums := 0
		numUpperChars := 0
		numLowerChars := 0
		numNumbers := 0
		numSymbols := 0
		for _, c := range *a.Form {
			switch c {
			case 'a':
				numAlphaNums++
			case 'x':
				numLowerChars++
			case 'X':
				numUpperChars++
			case 'n':
				numNumbers++
			case 's':
				numSymbols++
			}
		}

		numLowerChars = numLowerChars + (numAlphaNums - (numAlphaNums / 3))
		a.Numbers = numNumbers + (numAlphaNums / 3)
		a.Symbols = numSymbols
		a.Length = numLowerChars + numUpperChars + a.Numbers + a.Symbols
	}

	fmt.Printf("Symbols: %d\n", a.Symbols)
	fmt.Printf("Numbers: %d\n", a.Numbers)
	fmt.Printf("In place: %v\n", a.InPlace)
	if a.Form != nil {
		fmt.Printf("Form: %s\n", *a.Form)
	}
	fmt.Printf("Length: %d\n", a.Length)
	fmt.Printf("Pass name: %s\n", a.PassName)
	fmt.Printf("Copy to clipboard: %v\n", a.Clip)

	// Generate(length, numDigits, numSymbols int, noUpper, allowRepeat bool) (string, error) {
	res, err := password.Generate(a.Length, a.Numbers, a.Symbols, !a.Mixed, !a.Unique)
	if err != nil {
		return err
	}

	if a.Form != nil {

		// Utility functions to determine if it is a number or symbol
		isNumber := func(c rune) bool { return c >= '0' && c <= '9' }
		isSymbol := func(c rune) bool {
			return !isNumber(c) && !strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", c)
		}

		// Split the password into the different parts (chars, numbers, symbols)
		symbols := []rune{}
		numbers := []rune{}
		chars := []rune{}
		for _, c := range res {
			if isNumber(c) {
				numbers = append(numbers, c)
			} else if isSymbol(c) {
				symbols = append(symbols, c)
			} else {
				chars = append(chars, c)
			}
		}

		// Construct the final password using Form
		password := ""
		for _, c := range *a.Form {
			switch c {
			case 'x':
				password += strings.ToLower(string(chars[0]))
				chars = chars[1:]
			case 'X':
				password += strings.ToUpper(string(chars[0]))
				chars = chars[1:]
			case 'n':
				password += string(numbers[0])
				numbers = numbers[1:]
			case 's':
				password += string(symbols[0])
				symbols = symbols[1:]
			case '-', '_', '/', '\\', '#':
				password += string(c)
			case 'a':
				password += string(res[0])
				res = res[1:]
			default:

			}
		}
		res = password
	}

	if a.Clip {
		clipboard.Write(res)
	}

	fmt.Println(res)

	return nil
}
