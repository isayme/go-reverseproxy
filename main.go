package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var (
	target = flag.String("t", "", "target backend server URL")
	printV = flag.Bool("v", false, "print version")

	VERSION = "dev"
)

func main() {
	flag.Parse()

	if *printV {
		fmt.Printf("version: %s\n", VERSION)
		return
	}

	if *target == "" {
		flag.Usage()
		slog.Warn("target URL is required")
		os.Exit(-1)
	}

	targetURL, err := url.Parse(*target)
	if err != nil {
		slog.Error("invalud target URL", "target", *target, "err", err)
		os.Exit(-1)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	http.Handle("/", proxy)

	fmt.Printf("Reverse proxy listening on :8080, forwarding to %s\n", targetURL.Host)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("listen failed", "err", err)
		os.Exit(-1)
	}
}
