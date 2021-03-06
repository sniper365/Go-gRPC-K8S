package event

import (
	"reflect"

	"github.com/sniper365/Go-gRPC-K8S/hub"
)

var handler = hub.Handler{}

func GetHandler() hub.Handler {
	return handler
}

func RegisterEvent(event interface{}, eventFunc hub.EventFunc) {
	handler[reflect.TypeOf(event)] = eventFunc
}
