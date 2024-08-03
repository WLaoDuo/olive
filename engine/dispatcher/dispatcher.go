// Package dispatcher handles bout events.
package dispatcher

import "github.com/WLaoDuo/olive/engine/enum"

type Dispatcher interface {
	Dispatch(event *Event) error
	DispatcherType() enum.DispatcherTypeID
	DispatchTypes() []enum.EventTypeID
}
