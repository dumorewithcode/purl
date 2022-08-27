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

<br />

## Usage:
Examples:

```bash
$ cat urls.txt | purl -p http://127.0.0.1:8080 -c 50
$ cat domains.txt | httprobe --prefer-https -c 50 | waybackurls | purl -p http://127.0.0.1:8080
$ cat resolved.txt | gau -b png,jpg,gif | purl -p http://127.0.0.1:8080
```
<br />

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
| `-c` | set the concurrency. Default value 1. Use this setting to increase the rate at which the urls are sent to your proxy tool. | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -c 5` |
| `-t` | set the timeout in milliseconds. Default 10000. If you are getting any timeout errors, increase this | `cat resovled.txt \| purl -p http://127.0.0.1:8080 -t 100000` |
<br />

```
$ go get -u github.com/a3kSec/purl
```
<br />