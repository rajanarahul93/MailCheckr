package main;

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter the domain name to check: ")
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n ")
	for scanner.Scan(){
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err!= nil{
		log.Fatal("Error reading input: ", err)
	}
}

func checkDomain(domain string){
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	//Check for MX Records
	mxRecords, err:= net.LookupMX(domain)
	if err != nil{
		log.Printf("Error looking up MX records for %s: %v", domain, err)
	}
	if len(mxRecords) > 0 {
 		hasMX = true
	} else {
		hasMX = false
	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil{
		log.Printf("Error looking up TXT records for %s: %v", domain, err)
	} 
	for _, record := range txtRecords{
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}

	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil{
		log.Printf("Error looking up DMARC records for %s: %v", domain, err)
	}

	for _, record:= range dmarcRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}