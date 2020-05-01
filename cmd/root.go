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
	Use:   "sb2md <target>",
	Short: "CLI to convert Scrapbox page to Markdown",
	Long: usage(),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			target = args[0]
		}
	},
}

// Execute will fetch page and generate markdown
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if target != "" {
		fmt.Printf("target : %s\n", target)
		genMD()
	} else {
		help, _ := rootCmd.Flags().GetBool("help")
		if !help {
			rootCmd.Usage()
		}
	}
}

func genMD() error {
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
	rootCmd.PersistentFlags().BoolP("hatena", "n", false, "Use Hatena blog embd")
}

func usage() string {
	return `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`
}