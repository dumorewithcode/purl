package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type requestData struct {
	client         http.Client
	cookie         string
	headers        []string
	headersFile    string
	headerFromFile string
	cookieFromFile string
}

type duplicateFlags []string

func (dupes *duplicateFlags) String() string {
	return ""
}

func (dupes *duplicateFlags) Set(value string) error {
	*dupes = append(*dupes, strings.TrimSpace(value))
	return nil
}

//getUserAgent() Stolen from gau. Thanks Corben :)
func getUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (iPad; CPU OS 8_1_3 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B466 Safari/600.1.4",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_2) AppleWebKit/600.3.18 (KHTML, like Gecko) Version/8.0.3 Safari/600.3.18",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.2; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; Win64; x64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/44.0.2403.157 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 8_1_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12B440 Safari/600.1.4",
		"Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; KFTT Build/IML74K) AppleWebKit/537.36 (KHTML, like Gecko) Silk/3.68 like Chrome/39.0.2171.93 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 8_2 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12D508 Safari/600.1.4",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Ubuntu Chromium/37.0.2062.94 Chrome/37.0.2062.94 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.85 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:40.0) Gecko/20100101 Firefox/40.0",
		"Mozilla/5.0 (iPad; CPU OS 8_4_1 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H321 Safari/600.1.4",
	}

	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(userAgents))

	userAgent := userAgents[index]

	return userAgent
}

func getCookieFromFile(file string) string {
	value := ""

	if file != "" {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		//read up to 1mb of data
		buffer := make([]byte, 1024*1024)

		for {
			str, err := reader.Read(buffer)
			if err != nil {
				if err != io.EOF {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
				break
			}
			value = string(buffer[0:str])
		}
	}
	return value
}

func (r requestData) makeRequest(url string) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	req.Header.Set("User-Agent", getUserAgent())

	if r.cookie != "" {
		req.Header.Set("Cookie", r.cookie)
	}

	if r.cookieFromFile != "" {
		req.Header.Set("Cookie", r.cookieFromFile)
	}

	if len(r.headers) > 0 {
		for _, header := range r.headers {
			if strings.Contains(header, ":") {
				//Split by first occurrence of colon(:), any other colons after that will be ignored
				name, value := header[:strings.IndexByte(header, ':')], header[strings.IndexByte(header, ':')+1:]
				req.Header.Set(strings.TrimSpace(name), strings.TrimSpace(value))
			}
		}
	}

	if r.headerFromFile != "" {
		if strings.Contains(r.headerFromFile, ":") {
			//Split by first occurrence of colon(:), any other colons after that will be ignored
			name, value := r.headerFromFile[:strings.IndexByte(r.headerFromFile, ':')], r.headerFromFile[strings.IndexByte(r.headerFromFile, ':')+1:]
			req.Header.Set(strings.TrimSpace(name), strings.TrimSpace(value))
		}
	}

	resp, err := r.client.Do(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
}

func main() {
	var (
		concurrency    int
		proxy          string
		cookie         string
		cookieFromFile string
		headers        duplicateFlags
		headersFile    string
		timeout        int
	)

	flag.Var(&headers, "h", "specify header to include in request")
	flag.IntVar(&concurrency, "c", 1, "set the concurrency")
	flag.StringVar(&proxy, "p", "", "specify http proxy")
	flag.StringVar(&cookie, "b", "", "specify cookie VALUE to include in request")
	flag.StringVar(&cookieFromFile, "B", "", "specify file that contains the cookie VALUE to include in request. Reads up to 1mb of data")
	flag.StringVar(&headersFile, "H", "", "specify list of headers. This sends each header 1 at a time to the same request")
	flag.IntVar(&timeout, "t", 10000, "set the timeout in milliseconds")

	flag.Parse()

	sc := bufio.NewScanner(os.Stdin)

	proxyUrl, err := url.Parse(proxy)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	// timeout Stolen from httprobe. Thanks Tom :)
	to := time.Duration(timeout * 1000000)

	client := &http.Client{
		Timeout: to,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: to,
			}).DialContext,
			IdleConnTimeout:       time.Second,
			TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
			Proxy:                 http.ProxyURL(proxyUrl),
			ResponseHeaderTimeout: (to / 2),
			ForceAttemptHTTP2:     true,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req := requestData{
		client:         *client,
		headers:        headers,
		headersFile:    headersFile,
		cookie:         cookie,
		cookieFromFile: getCookieFromFile(cookieFromFile),
	}

	//Create buffer(queue) and set the size to the concurrency value
	queue := make(chan bool, concurrency)
	defer close(queue)

	//For each url that is read from stdin, we want to add it to the buffer(queue)
	//If the concurrency value is 5, then 5 requests will be added to the buffer(queue)
	//This limits the throughput to 5 requests. In otherwords only 5 urls at a time will
	//be passed to the proxy tool
	for sc.Scan() {
		url := sc.Text()
		//For example we set the concurrency value to 5, so the buffer(queue) has 5 values
		//this means that the buffer(queue) is full and it will wait until a request
		//has given up its position in the buffer(queue)
		queue <- true
		go func(url string) {
			// 'defer func() { <-queue }()' will empty this request's position in the buffer(queue)
			// and send a new request to the buffer(queue)
			defer func() { <-queue }()

			//Open headers file for reading, and resend the request but with the next header in line
			if req.headersFile != "" {
				file, err := os.Open(req.headersFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(1)
				}
				defer file.Close()

				sc := bufio.NewScanner(file)

				//Send the same request for each header in the headers file(-H flag)
				for sc.Scan() {
					req.headerFromFile = sc.Text()
					req.makeRequest(url)
				}
				if err := sc.Err(); err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
				}
			}
			req.makeRequest(url)
		}(url)
	}

	//flush queue
	for i := 0; i < concurrency; i++ {
		queue <- true
	}
}
