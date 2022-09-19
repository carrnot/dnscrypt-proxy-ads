package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	upstream = "https://raw.githubusercontent.com/bigdargon/hostsVN/master/option/hostsVN-surge-rule.conf"
)

const (
	prefixKeyword = "DOMAIN-KEYWORD,"
	prefixSuffix  = "DOMAIN-SUFFIX,"
)

func main() {
	resp, err := http.Get(upstream)
	if err != nil {
		log.Fatalln(err)
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	str := strings.Builder{}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "#") {
			continue
		}
		if strings.HasPrefix(txt, prefixKeyword) {
			str.WriteString("*" + strings.TrimPrefix(txt, prefixKeyword) + "*\n")
			continue
		}
		if strings.HasPrefix(txt, prefixSuffix) {
			str.WriteString(strings.TrimPrefix(txt, prefixSuffix)+"\n")
			continue
		}
	}
	
	fmt.Println(str.String())
}
