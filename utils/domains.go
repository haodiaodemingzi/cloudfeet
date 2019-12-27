package utils

import (
	"fmt"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// DomainToGFWConf ...
func DomainToGFWConf(topDomain string) string {
	gfwIPsetName := "ss_spec_dst_fw"
	return fmt.Sprintf(
		"server=/%s/127.0.0.1#8053\nipset=/%s/%s\n", topDomain, topDomain, gfwIPsetName)
}

// WriteGFWListFile ...
func ParseTopDomain(domain string) string {
	domain = strings.Trim(domain, "`~!@#$%^&*()_-+={}[]|\\:;'\",<.>/?")
	domain = strings.TrimSpace(domain)

	topDomain, err := publicsuffix.EffectiveTLDPlusOne(domain)
	if err != nil {
		topDomain = publicsuffix.List.PublicSuffix(domain)
	}
	return topDomain
}
