package main

import (
	"flag"
	"fmt"
	"hash/fnv"
	"os"
)

func getHash(text string) string {
	h := fnv.New32a()
	h.Write([]byte(text))
	return fmt.Sprintf("%x", h.Sum32())
}

func main() {
	ip := flag.String("ip", "", "ip address")
	host := flag.String("host", "", "host address")
	flag.Parse()
	if *ip == "" || *host == "" {
		fmt.Fprint(os.Stderr, "--ip and/or --host flags are not provided\n")
		os.Exit(1)
	}
	fmt.Printf(`etcdctl set /skydns/local/cluster/svc/default/%s/%s '{"host":"%s","priority":10,"weight":10,"ttl":30,"targetstrip":0}'%s`,
		*host,
		getHash(
			fmt.Sprintf(`{"host":"%s","priority":10,"weight":10,"ttl":30,"targetstrip":0}`, *ip)),
		*ip,
		"\n")
}
