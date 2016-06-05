package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	protocol := flag.String("protocol", "tcp", "protocol ('tcp' (default), 'tcp4' (IPv4-only), 'tcp6' (IPv6-only), 'udp', 'udp4' (IPv4-only), 'udp6' (IPv6-only), 'ip', 'ip4' (IPv4-only), 'ip6' (IPv6-only), 'unix', 'unixgram' and 'unixpacket')")
	host := flag.String("host", "", "destination host")
	port := flag.String("port", "", "destination port")
	times := flag.Int("times", 3, "repeat this many times")

	flag.Parse()

	if *host == "" {
		flag.PrintDefaults()
		log.Fatal("host missing, exiting.")
	}

	if *port == "" {
		flag.PrintDefaults()
		log.Fatal("port missing, exiting.")
	}

	address := fmt.Sprintf("%s:%s", *host, *port)
	log.Printf("Trying %s %s %d times", *protocol, address, *times)

	for i := 0; i < *times; i++ {
		conn, err := net.Dial(*protocol, address)
		if err == nil {
			log.Println("Destination reachable.")
			conn.Close()
			os.Exit(0)
		}
		log.Print(".")
		time.Sleep(1 * time.Second)
	}
	log.Println("Destination not reachable.")
	os.Exit(1)
}
