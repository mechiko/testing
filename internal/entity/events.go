package entity

type EventMonitor interface {
	Attach(Subcriber, EventInt) (bool, error)
	Detach(Subcriber, EventInt) (bool, error)
	Notify(EventInt, string) (bool, error)
}

type Subcriber interface {
	NameSubcriber() string
	UpdateSubcriber(val string, evt EventInt)
}

type EventInt int

const (
	UTM_STATUS_CHANGED EventInt = 1 + iota
	REQUEST_REST_SEND
	REQUEST_REST_ERROR
	REQUEST_REST_RECEIVE
	REQUEST_REST_PROCCESSED
	NEED_UPDATE_STATUS_BAR
	NEED_TICKS_EVERY_SECOND
	TASKS_STATUS_UPDATE
	ALL
)

func (e EventInt) String() string {
	switch e {
	case UTM_STATUS_CHANGED:
		return "UTM_STATUS_CHANGED"
	case REQUEST_REST_SEND:
		return "REQUEST_REST_SEND"
	case REQUEST_REST_ERROR:
		return "REQUEST_REST_ERROR"
	case REQUEST_REST_RECEIVE:
		return "REQUEST_REST_RECEIVE"
	case REQUEST_REST_PROCCESSED:
		return "REQUEST_REST_PROCCESSED"
	case NEED_UPDATE_STATUS_BAR:
		return "NEED_UPDATE_STATUS_BAR"
	case NEED_TICKS_EVERY_SECOND:
		return "NEED_TICKS_EVERY_SECOND"
	case TASKS_STATUS_UPDATE:
		return "TASKS_STATUS_UPDATE"
	case ALL:
		return "ALL"
	}
	return ""
}
