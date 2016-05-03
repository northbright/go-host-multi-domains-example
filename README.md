# go-host-multi-domains-example

[![Build Status](https://travis-ci.org/northbright/go-host-multi-domains-example.svg?branch=master)](https://travis-ci.org/northbright/go-host-multi-domains-example)

Example of hosting multiple domains in a single Golang HTTP server.

#### Host Multiple Domains in a Single Golang HTTP Server

To host multiple domains in a single Golang HTTP server, we need to use Host-specific pattern for [ServeMux](https://godoc.org/net/http#ServeMux):

>Patterns may optionally begin with a host name, restricting matches to URLs on that host only. Host-specific patterns take precedence over general patterns, so that a handler might register for the two patterns "/codesearch" and "codesearch.google.com/" without also taking over requests for "http://www.google.com/".

    // For Example:
    func main() {
        http.HandleFunc("a.com/", helloA)
        http.HandleFunc("sub.a.com/", helloSubA)
        http.HandleFunc("b.com/", helloB)

        if err := http.ListenAndServe(":80", nil); err != nil {
            log.Fatal(err)
        }
    }

#### Usage
* Add "a.com", "sub.a.com", "b.com" into `/etc/hosts`
* Restart network service or reboot.
* `sudo ./go-host-multi-domains-example`
* Run `curl a.com`, `curl sub.a.com`, `curl b.com` to test. 

#### License
* [MIT License](./LICENSE)

