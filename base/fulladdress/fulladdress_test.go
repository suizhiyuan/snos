package fulladdress

import (
	"errors"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFullAddresss(t *testing.T) {
	testTable := []struct {
		message string
		input   string
		output  *FullAddresss
		err     error
	}{
		{
			message: "ipv4 test1",
			input:   "192.168.1.1:8080",
			output: &FullAddresss{
				IPAddr: net.ParseIP("192.168.1.1"),
				Port:   8080,
			},
			err: nil,
		},
		{
			message: "ipv4 test2",
			input:   "10.0.0.10:22",
			output: &FullAddresss{
				IPAddr: net.ParseIP("10.0.0.10"),
				Port:   22,
			},
			err: nil,
		},
		{
			message: "ipv4 test3",
			input:   "127.0.0.1:300",
			output: &FullAddresss{
				IPAddr: net.ParseIP("127.0.0.1"),
				Port:   300,
			},
			err: nil,
		},
		{
			message: "ipv6 test1",
			input:   "[::1]:445",
			output: &FullAddresss{
				IPAddr: net.ParseIP("::1"),
				Port:   445,
			},
			err: nil,
		},
		{
			message: "ipv6 test2",
			input:   "[2001:db8:85a3::8a2e:370:7334]:8080",
			output: &FullAddresss{
				IPAddr: net.ParseIP("2001:db8:85a3::8a2e:370:7334"),
				Port:   8080,
			},
			err: nil,
		},
		{
			message: "ipv6 test3",
			input:   "[2001:0db8:85a3:0101:0303:8a2e:0370:7334]:30000",
			output: &FullAddresss{
				IPAddr: net.ParseIP("2001:0db8:85a3:0101:0303:8a2e:0370:7334"),
				Port:   30000,
			},
			err: nil,
		},
		{
			message: "illegal input",
			input:   "127.0.1:65536",
			output:  nil,
			err:     getUnrecognizeError("127.0.1:65536"),
		},
		{
			message: "illegal ipv4",
			input:   "300.0.1.100:65536",
			output:  nil,
			err:     getParseIPFailError("300.0.1.100:65536"),
		},
		{
			message: "illegal ipv6",
			input:   "[2001:db8:85a3::8a2e:370:ghgh]:80000",
			output:  nil,
			err:     getUnrecognizeError("[2001:db8:85a3::8a2e:370:ghgh]:80000"),
		},
		{
			message: "ipv4 port error",
			input:   "127.0.0.1:65536",
			output:  nil,
			err:     getPortOutOfRangeError("127.0.0.1:65536"),
		},
		{
			message: "ipv6 prot error",
			input:   "[::1]:80000",
			output:  nil,
			err:     getPortOutOfRangeError("[::1]:80000"),
		},
	}
	for _, test := range testTable {
		result, e := NewFullAddresss(test.input)
		assert.Equal(t, test.output, result, "fail result at %s", test.message)
		assert.Equal(t, test.err, e, "fail error at %s", test.message)
	}
	assert.Equal(t, 1, 1)
}

func TestString(t *testing.T) {
	testTable := []struct {
		message string
		input   *FullAddresss
		output  string
	}{
		{
			message: "ipv4 test",
			input: &FullAddresss{
				IPAddr: net.ParseIP("192.168.1.5"),
				Port:   443,
			},
			output: "192.168.1.5:443",
		},
		{
			message: "ipv6 test",
			input: &FullAddresss{
				IPAddr: net.ParseIP("2001:db8:85a3::8a2e:370:7334"),
				Port:   8080,
			},
			output: "[2001:db8:85a3::8a2e:370:7334]:8080",
		},
	}

	for _, test := range testTable {
		var innerInput fmt.Stringer = test.input
		result := innerInput.String()
		assert.Equal(t, test.output, result, "fail result at %s", test.message)
	}
}

func TestErrorString(t *testing.T) {
	testTable := []struct {
		message string
		input   ParseAddrError
		output  string
	}{
		{
			message: "nil inner error",
			input:   getUnrecognizeError("123.123.123.123.123:123"),
			output:  "parse (123.123.123.123.123:123) with error : can not recognize string",
		},
		{
			message: "with inner error",
			input:   getParsePortFailError("123,123,123", errors.New("inner error")),
			output:  "parse (123,123,123) with error : parse port fail, with inner error inner error",
		},
	}
	for _, test := range testTable {
		var innerError error = test.input
		result := innerError.Error()
		assert.Equal(t, test.output, result, "fail result at %s", test.message)
	}
}
