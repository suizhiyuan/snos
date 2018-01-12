package base

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/suizhiyuan/snos/base/fulladdress"
)

var logger = log.New(os.Stdout, "snos.base", log.Ldate|log.Ltime|log.Lshortfile)

// Service is a describe of gather of function , generaly a service mean a project
type Service struct {
	Name      string
	Version   string
	Functions []Function
}

// Function is a single Api in a service , a Function may have multipure parameter and at least one result
type Function struct {
	Name       string
	Parameters map[string]CommonType
	Result     map[string]CommonType
}

// CommonType is a struct among Servcie , the same hashCode mean same type
type CommonType struct {
	HashCode string
	Fields   map[string]string
}

const (
	nilHandlerError = errors.New("handler can not be nil")
)

var handlerStorage func(string, []byte) []byte
var locaPathStorate []fulladdress.FullAddresss

func init() {

}

// Startup is called to start up service
func Startup(service Service, host string, port int, localpath []fulladdress.FullAddresss, handler func(string, []byte) []byte) (err error) {
	if handler == nil {
		return nilHandlerError
	}
	localHandler = handler

	logger.Print("start up service")

	http.HandleFunc("/call", callHandler)
	http.HandleFunc("/document", documentHendler)

	e := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)

	logger.Print("server stopd")

	if e != nil {
		logger.Fatalf("server error : %v", e)
		return err
	}
	return nil
}

// Call a remote function
func Call(serviceName string, functionName string, parameter []byte) (result []byte) {
	return nil
}

// CallSpecific Call a remote function whith specitic instance
func CallSpecific(serviceName string, instance string, functionName string, parameter []byte) (result []byte) {
	return nil
}

func callHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.Post {

	}
}

func documentHendler(w http.ResponseWriter, req *http.Request) {

}
