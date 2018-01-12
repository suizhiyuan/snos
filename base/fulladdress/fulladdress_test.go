package fulladdress

import (
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
			message: "ipv4 port error",
			input:   "127.0.0.1:65536",
			output:  nil,
			err:     getPortOutOfRangeError("127.0.0.1:65536"),
		},
		{
			message: "illegal input",
			input:   "127.0.1:65536",
			output:  nil,
			err:     getUnrecognizeError("127.0.1:65536"),
		},
	}
	for _, test := range testTable {
		result, e := NewFullAddresss(test.input)
		assert.Equal(t, test.output, result, "fail result at %s", test.message)
		assert.Equal(t, test.err, e, "fail error at %s", test.message)
	}
	assert.Equal(t, 1, 1)
}
