<h1 align="center">
  <br>
  <a href="https://github.com/a3kSec/purl"><img src="https://i.ibb.co/smc89DZ/purl-icon.png" width="170" height="170" alt="purl"></a>
  <br>
  purl
  <br>
</h1>

<h2 align="center">proxy urls</h2>
<h4 align="center">A simple script to proxy full urls from stdin through your favorite http proxy tool very quickly for analysis</h4>

<p align="center">
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-orange"></a>
  <a href="https://github.com/a3kSec/purl/issues"><img src="https://img.shields.io/badge/contributions-welcome-blueviolet"></a>
  <a href="https://golang.org/"><img src="https://img.shields.io/badge/powered%20by-go-%2300ADD8"></a>
  <a href="https://goreportcard.com/report/github.com/a3kSec/purl"><img src="https://img.shields.io/badge/go%20report-A%2B-success"></a>
  <a href="https://github.com/a3kSec/purl#download"><img src="https://img.shields.io/badge/platform-any-ec2eb4.svg?style=flat"></a>
</p>

<p align="center">
  <a href="#usage">Usage</a> •
  <a href="#warningwarningimportantwarningwarning">Important</a> •
  <a href="#installation">Installation</a>
</p>
<br />

Contribution is much appreciated.
<br />
<br />

## Usage:
Examples:

```bash
$ cat urls.txt | purl -p http://127.0.0.1:8080 -c 50
$ cat domains.txt | httprobe --prefer-https -c 50 | waybackurls | purl -p http://127.0.0.1:8080
$ cat resolved.txt | gau -b png,jpg,gif | purl -p http://127.0.0.1:8080
$ cat urls.txt | httpx -follow-redirects -mc 403 -silent | purl -p http://127.0.0.1:8080 -h "X-Bug-Bounty: hunter" -h "X-Custom-IP-Authorization: 127.0.0.1" -H 403-bypass.txt -c 50
```
<br />

To display the help for the tool use the `-help` flag or pass any flag without a value:

```bash
$ purl -help
```

| Flag | Description | Example |
|------|-------------|---------|
| `-p` | specify http proxy | `cat resovled.txt \| purl -p http://127.0.0.1:8080` |
| `-h` | specify header/s to include in every request. Supports duplicate flags | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -h "X-Forwarded-For: example.com" -h "X-Wife: divorced"` |
| `-H` | specify file path that contains a list of headers to use. This will resend the same url(request) but with the next header in the file. Use `-h` flag for headers you want to include in every request | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -H headers.txt` |
| `-b` | specify cookie VALUE to include in request | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -b "cookie value"` |
| `-B` | specify file path that contains the cookie VALUE to include in request | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -B cookie.txt` |
| `-c` | set the concurrency. Default value 1. Use this setting to increase the rate at which the urls are sent to your proxy tool. Important information about using this can be found below | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -c 5` |
| `-t` | set the timeout in milliseconds. Default 10000. If you are getting any timeout errors, increase this | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -t 100000` |
<br />

### :warning::warning:IMPORTANT:warning::warning:
Default concurrency value is set to 1. This is a safegaurd against any proxy tools that have live auditing features.\
Setting this value too high will cause any proxy tools with these features to consume alot if not all of your system resources in the blink of an eye if there are no resource limits set for these tools.
So please use with caution.\
I recommend playing with the concurrency setting and turning on/off any live auditing/passive crawl features while monitoring the resource usage of your proxy tool to see what you're comfortable with.
<br />
<br />

## Installation:
```
$ go get -u github.com/a3kSec/purl
```
<br />

## Not necessary but if you wish to support

<a href="https://www.buymeacoffee.com/a3kSec"><img src="https://img.buymeacoffee.com/button-api/?text=Buy me a beer&emoji=🍺&slug=a3kSec&button_colour=00ff08&font_colour=000000&font_family=Cookie&outline_colour=000000&coffee_colour=FFDD00" alt="Buy Me A Beer" style="height: 41px !important;width: 174px !important;box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;-webkit-box-shadow: 0px 3px 2px 0px rgba(190, 190, 190, 0.5) !important;"></a>