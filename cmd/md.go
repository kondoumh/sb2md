package cmd

import (
	"strings"
	"regexp"
)

var (
	codeblock bool
	table bool
	renderTableHeader bool
	hatenaMarkdown bool
)

var (
	rgxCodeBlock = regexp.MustCompile(`^code:([^.]*)(\.([^.]*))?`)
	rgxTable = regexp.MustCompile(`^table:(.*)$`)
	rgxHeading = regexp.MustCompile(`^\[(\*+)\s([^\]]+)\]`)
	rgxIndent = regexp.MustCompile(`^(\s+)([^\s].+)`)
	rgxStrong = regexp.MustCompile(`\[(\*+)\s(.+)\]`)
	rgxLink = regexp.MustCompile(`\[https?:\/\/[^\s]*\s[^\]]*]`)
	rgxLinkInside = regexp.MustCompile(`\[(https?:\/\/[^\s]*)\s([^\]]*)]`)
	rgxGazo = regexp.MustCompile(`\[https:\/\/gyazo.com\/[^\]]*\]`)
	rgxGazoInside = regexp.MustCompile(`\[(https:\/\/gyazo.com\/[^\]]*)\]`)
)

// ToMd convert lines to Markdown
func ToMd(lines []string, hatena bool) string {
	result := ""
	hatenaMarkdown = hatena
	for idx, line := range lines {
		if idx == 0 {
			result = line + "\n=================\n"
		} else {
			result += convert(line) + "\n"
		}
	}
	return result
}

func convert(line string) string {
	result := "";
	if codeblock {
		if !strings.HasPrefix(line, " ") {
			result = "```\n"
			codeblock = false
		} else {
			result = line
			return result
		}
	} else if table {
		if !strings.HasPrefix(line, " ") {
			table = false;
			renderTableHeader = false;
		} else {
			str := strings.Replace(line, "\t", "$\t", -1)
			str = strings.Trim(str, " ")
			tr := strings.Split(str, "$\t")
			result = "| " + strings.Join(tr, " | ") + " |"
			if !renderTableHeader {
				result += "\n" + strings.Repeat("|:--", len(tr)) + "|"
				renderTableHeader = true
			}
		}
		return result
	}
	if !codeblock && !table {
		if rgxCodeBlock.Match([]byte(line)) {
			codeblock = true
			ar := rgxCodeBlock.FindStringSubmatch(line)
			result = ar[1] + ar[2] + "\n```" + ar[3]
			return result
		}
		if rgxTable.Match([]byte(line)) {
			table = true
			ar := rgxTable.FindStringSubmatch(line)
			result = ar[1] + "\n"
			return result
		}
		if rgxHeading.Match([]byte(line)) {
			ar := rgxHeading.FindStringSubmatch(line)
			result += strings.Repeat("#", decideLevel(len(ar[1]))) + " " + ar[2]
		} else if rgxIndent.Match([]byte(line)) {
			ar := rgxIndent.FindStringSubmatch(line)
			indent := strings.Repeat("  ", len(ar[1]) - 1)
			str := replaceMdLink(ar[2])
			str = replaceGazoImmage(str)
			str = replaceStrong(str)
			result += indent + "- " + str
		} else {
			str := replaceMdLink(line)
			str = replaceGazoImmage(str)
			str = replaceStrong(str)
			result += str
		}
	}
	return result
}

func decideLevel(len int) int {
	if len >= 4 {
		return 1
	}
	return 5 - len
}

func replaceMdLink(str string) string {
	result := str
	if rgxLink.Match([]byte(str)) {
		links := rgxLink.FindAllString(str, -1)
		for _, link := range links {
			result = strings.Replace(result, link, toMdLink(link), -1)
		}
	}
	return result
}

func toMdLink(link string) string {
	ar := rgxLinkInside.FindStringSubmatch(link)
	if hatenaMarkdown {
		return "[" + ar[1] + ":embed:cite]"
	}
	return "[" + ar[2] + "](" + ar[1] + ")"
}

func replaceGazoImmage(str string) string {
	result := str
	if rgxGazo.Match([]byte(str)) {
		links := rgxGazo.FindAllString(str, -1)
		for _, link := range links {
			result = strings.Replace(result, link, toGazoLink(link), -1)
		}
	}
	return result
}

func toGazoLink(link string) string {
	ar := rgxGazoInside.FindStringSubmatch(link)
	return "![](" + ar[1] + ".png)"
}

func replaceStrong(str string) string {
	result := str
	if rgxStrong.Match([]byte(str)) {
		strongs := rgxStrong.FindAllString(str, -1)
		for _, strong := range strongs {
			ar := rgxStrong.FindStringSubmatch(strong)
			result = strings.Replace(result, strong, "**" + ar[2] + "**", -1)
		}
	}
	return result
}