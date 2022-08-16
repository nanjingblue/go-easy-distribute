package log

import (
	"bytes"
	"fmt"
	"github.com/nanjingblue/go-easy-distribute/registry"
	stdlog "log"
	"net/http"
)

func SetClientLogger(serviceURl string, clientService registry.ServiceName) {
	stdlog.SetPrefix(fmt.Sprintf("[%v] - ", clientService))
	stdlog.SetFlags(0)
	stdlog.SetOutput(&clientLogger{url: serviceURl})
}

type clientLogger struct {
	url string
}

func (cl clientLogger) Write(data []byte) (int, error) {
	b := bytes.NewBuffer([]byte(data[:len(data)-1]))
	res, err := http.Post(cl.url+"/log", "text/plain", b)
	if err != nil {
		return 0, err
	}
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to send log message. service responded wrong: %v", res.StatusCode)
	}
	return len(data), nil
}
