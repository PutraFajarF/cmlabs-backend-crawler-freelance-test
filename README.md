# Go Web Crawler API

This is a simple API built in Golang to perform website crawling and save the results in HTML files. The API uses Chromedp, a headless browser library, to render JavaScript-based websites such as Single-Page Applications (SPA), Server-Side Rendered (SSR) websites, and Progressive Web Apps (PWA).

### Features

1. Crawling Single or Multiple Websites: The API allows you to crawl a single website or multiple websites simultaneously by providing the URLs as a comma-separated list.
2. Custom User-Agent: You can specify a custom User-Agent header in the request to mimic different browsers or devices.

### Installation

1. Clone the repository to your local machine:

```bash
git clone hhttps://github.com/PutraFajarF/cmlabs-backend-crawler-freelance-test.git

cd cmlabs-backend-crawler-freelance-test
```

2. Install the dependencies:
```bash
go get github.com/gorilla/mux
go get github.com/chromedp/chromedp
go get github.com/chromedp/cdproto/network
```

### Usage

1. Start the server

```bash
go run .
```

2. Perform a crawling request using your web browser or a tool like Postman:
```bash
GET http://localhost:8080/crawl?url=https://example.com
```
Replace https://example.com with the URL you want to crawl. You can also specify multiple URLs as a comma-separated list.

Optional parameters: 
- user_agent: Set a custom User-Agent header (Default: Chrome on Windows).

### Example
To crawl a single website:
```bash
GET http://localhost:8080/crawl?url=https://cmlabs.co
```

To crawl multiple websites:
```bash
GET http://localhost:8080/crawl?url=https://cmlabs.co,https://sequence.day
```

Using a User-Agent of Google Chrome on Windows:
```bash
GET http://localhost:8080/crawl?url=https://example.com&user_agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36
```

Using a User-Agent of Apple Safari on macOS:
```bash
GET http://localhost:8080/crawl?url=https://example.com&user_agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Safari/537.36
```

Using a User-Agent of Mozilla Firefox on Windows:
```bash
GET http://localhost:8080/crawl?url=https://example.com&user_agent=Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0
```

Using a User-Agent of Apple Iphone:
```bash
GET http://localhost:8080/crawl?url=https://example.com&user_agent=Mozilla/5.0 (iPhone; CPU iPhone OS 14_7 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15E148 Safari/604.1
```

### Output
After performing the crawling request, the API will generate HTML files for each URL crawled. The files will be named with the format output_<url>.html and will be saved in the same directory as the API.

The API will respond with the message "Crawling finished. Results are saved in HTML files." once the crawling process is completed.

### Note
- This API is designed to handle SPA, SSR, and PWA websites.
- Make sure to use a valid URL for crawling.
- The default User-Agent header mimics Chrome on Windows. You can change it by providing the `user_agent` parameter in the request.