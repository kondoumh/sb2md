package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"

	"github.com/mamezou-tech/sbgraph/pkg/api"
)

type page struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Views  int    `json:"views"`
	Linked int    `json:"linked"`
	Lines  []line `json:"lines"`
}

type line struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

var target string

var rootCmd = &cobra.Command{
	Use:   "sb2md <project-name>/<page-title>",
	Short: "CLI to convert Scrapbox page to Markdown",
	Long: LongUsage(`
		sb2md is a CLI for outputting Scrapbox pages in Markdown format.
		Fetches the page data, converts it to Markdown format, and outputs it to standard output.
	`),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			target = args[0]
		}
	},
}

// Execute sets flags
func Execute() {
	err := rootCmd.Execute()
	CheckErr(err)

	if target != "" {
		fmt.Printf("Contents of %s\n", target)
		genMd()
	} else {
		help, _ := rootCmd.Flags().GetBool("help")
		if !help {
			rootCmd.Usage()
			os.Exit(1)
		}
	}
}

func genMd() error {
	hatena, _ := rootCmd.PersistentFlags().GetBool("hatena")
	fmt.Printf("hatena: %t\n", hatena)
	bytes, _ := api.FetchPage("kondoumh", "Dev")
	var pg page
	json.Unmarshal(bytes, &pg)
	for _, line := range pg.Lines {
		fmt.Println(line.Text)
	}
	return nil
}

func init() {
	rootCmd.PersistentFlags().BoolP("hatena", "n", false, "Generate links in Hatena blog format")
}
