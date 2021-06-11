<h1 align="center">
  <br>
  <a href="https://github.com/a3kSec/purl"><img src="https://image.ibb.co/" alt="purl"></a>
  <br>
  purl
  <br>
</h1>

<h3 align="center">proxy urls</h3>
<h4 align="center">proxy urls(absolute) from stdin through your favorite proxy tool</h4>

<p align="center">
  <img src="https://img.shields.io/badge/license-MIT-_red.svg">
  <img src="https://img.shields.io/badge/release-v0.9-blue">
</p>

# Resources
- [Usage](#usage)
- [Important](#usage)
- [Installation](#installation)

## Usage:
Examples:

```bash
$ echo http://example.com | purl -p http://127.0.0.1:8080
$ cat resolved.txt | purl -p http://127.0.0.1:8080
```

To display the help for the tool use the `-h` flag:

```bash
$ purl -help
```

| Flag | Description | Example |
|------|-------------|---------|
| `-p` | specify http proxy | `cat resovled.txt \| purl -p http://127.0.0.1:8080` |
| `-h` | specify header/s to include in request. supports duplicate headers | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -h "X-Forwarded-For: example.com" -h "X-Wife: wifeyyyy"` |
| `-H` | specify the path of the list of headers to use | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -H headers.txt` |
| `-b` | specify cookie VALUE to include in request | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -b "cookie value"` |
| `-B` | specify file path that contains the cookie VALUE to include in request. Reads up to 1mb of data | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -B cookie.txt` |
| `-c` | set the concurrency. Default value 1. Important information about using this can be found below | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -c 5` |
| `-t` | set the timeout in milliseconds. Default 10000 | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -t 100000` |

### :warning::warning:IMPORTANT:warning::warning:
Default concurrency value is set to 1. This is a safegaurd against any proxy tools that have live auditing features.\
Setting this value too high will cause any tools with these features to consume alot if not all of your system resources if there are no resource limits set for these tools.
So please use with caution.\
I recommend playing with the concurrency setting and turning on/off any live auditing/passive crawl features to see what you're comfortable with.

## Installation:
```
$ go get github.com/a3kSec/purl
```

## If you wish

<a href="https://www.buymeacoffee.com/a3kSec" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;" ></a>