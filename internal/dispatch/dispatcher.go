package dispatch

import (
	"fmt"

	"github.com/dgdraganov/noti-fire/internal/model"
	"go.uber.org/zap"
)

type notificationDispatcher struct {
	drivers map[string]Driver
	logs    *zap.SugaredLogger
}

// NewNotificationDispatcher is a constructor function for the notificationDispatcher type
func NewNotificationDispatcher(logger *zap.SugaredLogger) *notificationDispatcher {
	return &notificationDispatcher{
		drivers: map[string]Driver{},
		logs:    logger,
	}
}

// RegisterDriver registers a new driver to the list of drivers
func (d *notificationDispatcher) RegisterDriver(name string, driver Driver) {
	_, ok := d.drivers[name]
	if ok {
		panic(fmt.Sprintf("driver with that name is already registered: %s", name))
	}

	d.drivers[name] = driver
}

// Dispatch releases the target message through all available drivers
func (d *notificationDispatcher) Dispatch(msg model.NotificationMessage) {
	for name, driver := range d.drivers {
		err := driver.Send(msg)
		if err != nil {
			d.logs.Errorf(
				"dispatching message failed",
				"driver", name,
				"error", err,
			)
			continue
		}
		d.logs.Infof(
			"successfully dispatched message",
			"driver", name,
		)
	}
}
