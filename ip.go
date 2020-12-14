package cymru

import (
	"fmt"
	"net"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ReplyIP struct {
	ASN         uint64 `json:"asn"`
	Prefix      string `json:"prefix"`
	CountryCode string `json:"countrycode"`
	Authority   string `json:"authority"`
	Since       string `json:"since"`

	mask int
}

func QueryIP(q string) (ret ReplyIP, err error) {
	ip := net.ParseIP(q)
	if ip == nil {
		err = fmt.Errorf("%s is not a valid IP address", q)
		return
	}

	var res []string
	if ip.To4() == nil {
		// IPv6
		res, err = query6(q)
	} else {
		// IPv4
		res, err = query4(q)
	}
	if err != nil {
		return
	}

	if len(res) == 0 {
		err = fmt.Errorf("No valid reply returned")
		return
	}

	var r ReplyIP
	var x []ReplyIP
	for _, v := range res {
		r, err = parseReplyIP(v)
		if err != nil {
			return
		}
		x = append(x, r)
	}

	if len(x) > 1 {
		sort.Slice(x, func(i, j int) bool {
			return x[i].mask > x[j].mask
		})
	}

	ret = x[0]

	return
}

func query4(q string) ([]string, error) {
	x := strings.Split(q, ".")
	y := fmt.Sprintf("%s.%s.%s.%s.origin.asn.cymru.com", x[3], x[2], x[1], x[0])
	return net.LookupTXT(y)
}

func query6(q string) (ret []string, err error) {
	err = fmt.Errorf("Not yet implemented!")
	return
}

func parseReplyIP(in string) (ret ReplyIP, err error) {
	x := strings.Split(in, " | ")
	if len(x) != 5 {
		err = fmt.Errorf("Bad format input")
		return
	}

	ret.ASN, err = strconv.ParseUint(x[0], 10, 64)
	if err != nil {
		return
	}

	_, n, err := net.ParseCIDR(x[1])
	if err != nil {
		return
	}
	ret.Prefix = x[1]
	ret.mask, _ = n.Mask.Size()

	if len(x[2]) != 2 {
		err = fmt.Errorf("Invalid country code format")
		return
	}
	ret.CountryCode = x[2]

	if len(x[3]) == 0 {
		err = fmt.Errorf("No authority specified")
		return
	}
	ret.Authority = x[3]

	_, err = time.Parse("2006-01-02", x[4])
	if err != nil {
		return
	}
	ret.Since = x[4]

	return
}
