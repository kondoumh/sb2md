# sb2md

CLI to convert Scrapbox page to Markdown.

## Installing

```
go get -u github.com/kondoumh/sb2md
```

or 

```
curl -LO https://github.com/kondoumh/sb2md/releases/download/<version>/sb2md-<platform>-amd64.tar.gzff
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
