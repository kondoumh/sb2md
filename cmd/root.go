package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/mamezou-tech/sbgraph/pkg/api"
	"github.com/spf13/cobra"
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

var (
	targetProject string
	targetPage    string
	rgxTarget     = regexp.MustCompile(`([^\/]+)/([^\/]+)`)
)

var rootCmd = &cobra.Command{
	Use:   "sb2md <project-name>/<page-title>",
	Short: "CLI to convert Scrapbox page to Markdown",
	Long: LongUsage(`
		sb2md is a CLI for outputting Scrapbox pages in Markdown format.
		Fetches the page data, converts it to Markdown format, and outputs it to standard output.
	`),

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			arg := strings.Join(args, " ")
			if rgxTarget.Match([]byte(arg)) {
				ar := rgxTarget.FindStringSubmatch(arg)
				targetProject = strings.Replace(ar[1], " ", "-", -1)
				targetPage = ar[2]
			}
		}
	},
}

// Execute sets flags
func Execute() {
	err := rootCmd.Execute()
	CheckErr(err)

	if targetProject != "" && targetPage != "" {
		genMd()
	} else {
		help, _ := rootCmd.Flags().GetBool("help")
		if !help {
			rootCmd.Usage()
			os.Exit(1)
		}
	}
}

func genMd() {
	hatena, _ := rootCmd.PersistentFlags().GetBool("hatena")
	bytes, err := api.FetchPage(targetProject, targetPage)
	CheckErr(err)

	var pg page
	err = json.Unmarshal(bytes, &pg)
	CheckErr(err)

	var lines []string
	for _, line := range pg.Lines {
		lines = append(lines, line.Text)
	}
	result := ToMd(lines, hatena)
	fmt.Println(result)
}

func init() {
	rootCmd.PersistentFlags().BoolP("hatena", "n", false, "Generate links in Hatena blog format")
}
