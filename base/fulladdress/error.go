package fulladdress

import "fmt"

// ParseAddrError return by ParseFullAddresss while input string illegal
type ParseAddrError struct {
	input    string
	innerErr error
	message  string
}

// Error impl interface builtin/error
func (err ParseAddrError) Error() string {
	return fmt.Sprintf("parse addr error , intput string (%s)", err.input)
}

func getUnrecognizeError(str string) (err ParseAddrError) {
	return ParseAddrError{
		input:    str,
		innerErr: nil,
		message:  `can not recognize string`,
	}
}

func getParseIPFailError(str string) (err ParseAddrError) {
	return ParseAddrError{
		input:    str,
		innerErr: nil,
		message:  `parse ip fail`,
	}
}

func getParsePortFailError(str string, inner error) (err ParseAddrError) {
	return ParseAddrError{
		input:    str,
		innerErr: inner,
		message:  `parse port fail`,
	}
}

func getPortOutOfRangeError(str string) (err ParseAddrError) {
	return ParseAddrError{
		input:    str,
		innerErr: nil,
		message:  "port value out of range 0-65535",
	}
}
