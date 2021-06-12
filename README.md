<h1 align="center">
  <br>
  <a href="https://github.com/a3kSec/purl"><img src="https://i.ibb.co/smc89DZ/purl-icon.png" width="170" height="170" alt="purl"></a>
  <br>
  purl
  <br>
</h1>

<h2 align="center">proxy urls</h2>
<h4 align="center">a simple standalone script to proxy full urls from stdin through your favorite proxy tool at the "speed of light" for analysis</h4>

<p align="center">
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
  <a href="https://github.com/a3kSec/purl/releases"><img src="https://img.shields.io/badge/release-v0.9-blue"></a>
</p>

Made this with the intent to learn go and because I wanted it but couldn't really find a standalone tool that does this. Any advice and contribution is welcomed.

## Contents
- [Usage](#usage)
- [Important](#warningwarningimportantwarningwarning)
- [Installation](#installation)

## Usage:
Examples:

```bash
$ echo http://example.com | purl -p http://127.0.0.1:8080
$ cat resolved.txt | purl -p http://127.0.0.1:8080
```

To display the help for the tool use the `-help` flag:

```bash
$ purl -help
```

| Flag | Description | Example |
|------|-------------|---------|
| `-p` | specify http proxy | `cat resovled.txt \| purl -p http://127.0.0.1:8080` |
| `-h` | specify header/s to include in request. supports duplicate flags | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -h "X-Forwarded-For: example.com" -h "X-Wife: wifeyyyy"` |
| `-H` | specify the path of the list of headers to use | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -H headers.txt` |
| `-b` | specify cookie VALUE to include in request | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -b "cookie value"` |
| `-B` | specify file path that contains the cookie VALUE to include in request. Reads up to 1mb of data | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -B cookie.txt` |
| `-c` | set the concurrency. Default value 1. Important information about using this can be found below | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -c 5` |
| `-t` | set the timeout in milliseconds. Default 10000 | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -t 100000` |

### :warning::warning:IMPORTANT:warning::warning:
Default concurrency value is set to 1. This is a safegaurd against any proxy tools that have live auditing features.\
Setting this value too high will cause any proxy tools with these features to consume alot if not all of your system resources in the blink of an eye if there are no resource limits set for these tools.
So please use with caution.\
I recommend playing with the concurrency setting and turning on/off any live auditing/passive crawl features to see what you're comfortable with.

## Installation:
```
$ go get github.com/a3kSec/purl
```

## Not necessary but if you wish to support

<a href="https://www.buymeacoffee.com/a3kSec" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>