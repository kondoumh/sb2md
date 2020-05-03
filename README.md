# sb2md

![Go](https://github.com/kondoumh/sb2md/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kondoumh/sb2md)](https://goreportcard.com/report/github.com/kondoumh/sb2md)

sb2md is a CLI for outputting Scrapbox pages in Markdown format.
Fetches the page data, converts it to Markdown format, and outputs it to standard output.

## Installing

```
go get -u github.com/kondoumh/sb2md
```

or 

```
curl -LO https://github.com/kondoumh/sb2md/releases/download/<version>/sb2md-<platform>-amd64.tar.gz
tar xvf sb2md-<platform>-amd64.tar.gz
sudo mv sb2md /usr/local/bin
```

## Usage

Specify the Scrapbox page with the project name and title.

```
sb2md <project name>/<page title> [flags]
```

Markdown data is output to standard output.

## Flags

### -n, --hatena
Generates embedded links in Hatena blog format.

### -h, --help
Print help
