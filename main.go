package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Domain | HasMX | HasSPF | SpfRecord | HasDMARC | DmarcRecord \n")

	fmt.Println("Write your email ..")
	for scanner.Scan() {

		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error: could'nt read form input :%v\n", err)
	}

}

func checkDomain(domain string) {

	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("MxRecords -> Error :%v\n", err)
	}

	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("TxtRecords -> Error :%v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc" + domain)

	if err != nil {
		log.Printf("DmarcRecords -> Error :%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=MARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("Domain: %v \n HasMX: %v \n HasSPF: %v \n SpfRecord: %v \n HasDMARC: %v \n DmarcRecord: %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}
