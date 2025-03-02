package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"permission-cat/internal/word"
	"strings"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamel
	ModeUnderscoreToLowerCamel
	ModeCamelcaseToUnderscore
)

var DESCRIPTIONS = strings.Join([]string{
	"A number of word format transforming is supported: ",
	"1. Any -> Uppercase",
	"2. Any -> Lowercase",
	"3. Underscore -> UpperCamel",
	"4. Underscore -> LowerCamel",
	"5. Camel -> Underscore",
}, "\n")

var input string
var mode int8

var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "Word format transforming",
	Long:  DESCRIPTIONS,
	Run: func(cmd *cobra.Command, args []string) {
		var content string

		switch mode {
		case ModeUpper:
			content = word.ToUpper(input)
			break
		case ModeLower:
			content = word.ToLower(input)
			break
		case ModeUnderscoreToUpperCamel:
			content = word.UnderscoreToUpperCamelCase(input)
			break
		case ModeUnderscoreToLowerCamel:
			content = word.UnderscoreToLowerCamelCase(input)
			break
		case ModeCamelcaseToUnderscore:
			content = word.CamelCaseToUnderscore(input)
			break
		default:
			log.Fatalln("Unsupported mode. Use help word to view usages.")
		}

		log.Printf("Result: %s\n", content)
	},
}

func init() {
	wordCmd.Flags().StringVarP(&input, "word", "w", "", "Input word")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "Input transform mode")
}
