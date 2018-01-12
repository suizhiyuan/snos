package fulladdress

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
)

var (
	regIPv4 *regexp.Regexp
	regIPv6 *regexp.Regexp
)

func init() {
	regIPv4, _ = regexp.Compile(`^((?:[0-9]{1,3}\.){3}[0-9]{1,3}):([0-9]{1,5})$`)
	regIPv6, _ = regexp.Compile(`^\[((?:(?:[a-fA-F0-9]{1,4})?:){2,7}(?:[a-fA-F0-9]{1,4}))\]:([0-9]{1,5})$`)
}

// FullAddresss include
type FullAddresss struct {
	IPAddr net.IP
	Port   uint16
}

// NewFullAddresss create FullAddresss from string
func NewFullAddresss(str string) (addr *FullAddresss, err error) {
	for _, reg := range []*regexp.Regexp{regIPv4, regIPv6} {
		match := reg.FindStringSubmatch(str)
		if len(match) != 3 {
			continue
		}
		ipAddr := net.ParseIP(match[1])
		if ipAddr == nil {
			return nil, getParseIPFailError(str)
		}
		port, e := parsePort(match[2])
		if e != nil {
			ei := e.(ParseAddrError)
			ei.input = str
			return nil, ei
		}
		addr = &FullAddresss{
			IPAddr: ipAddr,
			Port:   uint16(port),
		}
		return
	}
	return nil, getUnrecognizeError(str)
}

func (addr *FullAddresss) String() (str string) {
	if ipv4 := addr.IPAddr.To4(); ipv4 == nil {
		return fmt.Sprintf("[%s]:%d", addr.IPAddr.String(), addr.Port)
	}
	return fmt.Sprintf("%s:%d", addr.IPAddr.String(), addr.Port)
}

func parsePort(str string) (port uint16, err error) {
	interalPort, e := strconv.ParseInt(str, 10, 32)
	if e != nil {
		return 0, getParsePortFailError("", e)
	}
	if interalPort >= 65536 || interalPort < 0 {
		return 0, getPortOutOfRangeError("")
	}
	return uint16(interalPort), nil
}
