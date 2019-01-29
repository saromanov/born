package notifications

import structs "github.com/saromanov/born/structs/v1"

// Notifications provides sending of notification after build
type Notifications interface {
	Do(*structs.Notification)
}
