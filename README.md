[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/cymru)](https://goreportcard.com/report/github.com/mskrha/cymru)

## cymru

### Description
Golang library for accessing [Team Cymru](https://team-cymru.com/) [IP to ASN Mapping Service](https://team-cymru.com/community-services/ip-asn-mapping/) using DNS.

### Important notes
* This project is at the very beginning, so there can be bugs, feel free to report them.
* IPv6 is not yet implemented.

### Installation
`go get github.com/mskrha/cymru`

### Example usage for IPv4
```go
package main

import (
	"fmt"

	"custom/cymru"
)

func main() {
	res, err := cymru.QueryIP("8.8.8.8")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
```

### Example usage for AS number
```go
package main

import (
	"fmt"

	"custom/cymru"
)

func main() {
	res, err := cymru.QueryASN("112")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", res)
}
```
