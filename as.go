package cymru

import (
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
)

type ReplyASN struct {
	ASN         uint64 `json:"asn"`
	CountryCode string `json:"countrycode"`
	Authority   string `json:"authority"`
	Since       string `json:"since"`
	Description string `json:"description"`
}

func QueryASN(q string) (ret ReplyASN, err error) {
	as, err := strconv.ParseUint(q, 10, 64)
	if err != nil {
		return
	}

	res, err := net.LookupTXT(fmt.Sprintf("AS%d.asn.cymru.com", as))
	if err != nil {
		return
	}

	switch len(res) {
	case 0:
		err = fmt.Errorf("No valid reply returned")
	case 1:
		ret, err = parseReplyASN(res[0])
	default:
		err = fmt.Errorf("Invalid reply returned")
	}

	return
}

func parseReplyASN(in string) (ret ReplyASN, err error) {
	x := strings.Split(in, " | ")
	if len(x) != 5 {
		err = fmt.Errorf("Bad format input")
		return
	}

	ret.ASN, err = strconv.ParseUint(x[0], 10, 64)
	if err != nil {
		return
	}

	if len(x[1]) != 2 {
		err = fmt.Errorf("Invalid country code format")
		return
	}
	ret.CountryCode = x[1]

	if len(x[2]) == 0 {
		err = fmt.Errorf("No authority specified")
		return
	}
	ret.Authority = x[2]

	_, err = time.Parse("2006-01-02", x[3])
	if err != nil {
		return
	}
	ret.Since = x[3]

	if len(x[4]) == 0 {
		err = fmt.Errorf("No description specified")
		return
	}
	ret.Description = x[4]

	return
}
