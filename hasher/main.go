package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
)

func getHash(text string) string {
	h := fnv.New32a()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum32())
}

func main() {
	ip := flag.String("ip", "", "ip address")
	host := flag.String("host", "", "host address")
	domain := flag.String("domain", "default.svc.skydns.local", "domain name")
	prefix := flag.String("prefix", "/skydns", "SkyDNS etcd path prefix")
	flag.Parse()
	if *ip == "" || *host == "" {
		fmt.Fprint(os.Stderr, "--ip and/or --host flags are not provided\n")
		os.Exit(1)
	}
	var domainReversed []string
	for _, dom := range strings.Split(*domain, ".") {
		domainReversed = append([]string{dom}, domainReversed...)
	}
	fmt.Printf(`etcdctl set %s/%s/%s/%s '{"host":"%s","priority":10,"weight":10,"ttl":30,"targetstrip":0}'%s`,
		*prefix,
		strings.Join(domainReversed, "/"),
		*host,
		getHash(
			fmt.Sprintf(`{"host":"%s","priority":10,"weight":10,"ttl":30,"targetstrip":0}`, *ip)),
		*ip,
		"\n")
}
