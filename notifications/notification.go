package notifications

import "github.com/0prototype/gwall-master/entities"

type Notification interface {
	SendAll([]entities.User)
}
