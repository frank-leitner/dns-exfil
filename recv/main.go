package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var debug = flag.Bool("d", false, "enable debugging.")
	var qlog = flag.String("qlog", "", "path to the dns query log file.")
	var help = flag.Bool("help", false, "show help.")
	var marker = flag.String("marker", "jzp", "a unique marker to identify the file in the dns logs.")

	flag.Parse()

	if *help || len(os.Args) == 1 {
		flag.PrintDefaults()
		return
	}

	bytes, err := ioutil.ReadFile(*qlog)
	if err != nil {
		fmt.Print(err)
	}

	// By default bytes are read
	//fmt.Println(bytes)

	// Convert bytes to a string
	str := string(bytes)

	// Split string on newlines
	lines := strings.Split(str, "\n")

	// Map to hold payloads (unordered and cannot be sorted)
	payloadParts := make(map[int]string)

	// Get slice of uniq queries
	for _, line := range lines {
		if strings.Contains(line, *marker) {
			payload := strings.Split(line, " ")[6]
			strIndex := strings.Split(payload, ".")[0]
			index, _ := strconv.Atoi(strIndex)
			payloadParts[index] = payload
			if *debug {
				log.Printf("%d\n", index)
				log.Printf("%s\n", payload)
			}
		}
	}

	// Unique list of the chunk indexes in a slice
	keys := make([]int, 0)

	for key, _ := range payloadParts {
		keys = append(keys, key)
	}

	// Sort the chunk indexes
	sort.Ints(keys)

	payload := make([]byte, 0)

	for key := range keys {
		if *debug {
			log.Printf("%s\n", payloadParts[key])
		}
		part := payloadParts[key]

		numberOfLabels := strings.Split(part, ".")[2]

		if numberOfLabels == "1" {
			label1 := strings.Split(part, ".")[3]
			src := []byte(label1)
			dst := make([]byte, hex.DecodedLen(len(src)))
			bytesWritten, err := hex.Decode(dst, src)
			if err != nil {
				fmt.Printf("%d\n", bytesWritten)
				log.Fatal(err)
			}
			payload = append(payload, dst...)
		}

		if numberOfLabels == "2" {
			label1 := strings.Split(part, ".")[3]
			label2 := strings.Split(part, ".")[4]

			src1 := []byte(label1)
			dst1 := make([]byte, hex.DecodedLen(len(src1)))
			bytesWritten1, err1 := hex.Decode(dst1, src1)
			if err1 != nil {
				fmt.Printf("%d\n", bytesWritten1)
				log.Fatal(err1)
			}

			src2 := []byte(label2)
			dst2 := make([]byte, hex.DecodedLen(len(src2)))
			bytesWritten2, err2 := hex.Decode(dst2, src2)
			if err2 != nil {
				fmt.Printf("%d\n", bytesWritten2)
				log.Fatal(err2)
			}

			payload = append(payload, dst1...)
			payload = append(payload, dst2...)
		}

		if numberOfLabels == "3" {
			label1 := strings.Split(part, ".")[3]
			label2 := strings.Split(part, ".")[4]
			label3 := strings.Split(part, ".")[5]

			src1 := []byte(label1)
			dst1 := make([]byte, hex.DecodedLen(len(src1)))
			bytesWritten1, err1 := hex.Decode(dst1, src1)
			if err1 != nil {
				fmt.Printf("BytesWritten1: %d\n", bytesWritten1)
				fmt.Printf("src1 len: %d\n", len(src1))
				fmt.Printf("label1 len: %d\n", len(label1))
				fmt.Printf("label1: %s\n", label1)
				log.Fatal(err1)
			}

			src2 := []byte(label2)
			dst2 := make([]byte, hex.DecodedLen(len(src2)))
			bytesWritten2, err2 := hex.Decode(dst2, src2)
			if err2 != nil {
				fmt.Printf("%d\n", bytesWritten2)
				log.Fatal(err2)
			}

			src3 := []byte(label3)
			dst3 := make([]byte, hex.DecodedLen(len(src3)))
			bytesWritten3, err3 := hex.Decode(dst3, src3)
			if err3 != nil {
				fmt.Printf("%d\n", bytesWritten3)
				log.Fatal(err3)
			}

			payload = append(payload, dst1...)
			payload = append(payload, dst2...)
			payload = append(payload, dst3...)
		}
	}

	fmt.Printf("%s", payload)
	if *debug {
		log.Printf("%d\n", len(payload))
	}
}
