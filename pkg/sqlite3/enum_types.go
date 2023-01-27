package sqlite3

type UriType int

const (
	RwModeWithCreate UriType = iota
	RoMode
	RwMode
)

func (s UriType) String() string {
	switch s {
	case RwModeWithCreate:
		return "?cache=shared&mode=rwc"
	case RoMode:
		return "?cache=shared&mode=ro"
	case RwMode:
		return "?cache=shared&mode=rw"
	}
	return "?cache=shared&mode=rw"
}

type CheckMode int

const (
	NoMatter CheckMode = iota
	Version
)

func (s CheckMode) String() string {
	switch s {
	case Version:
		return "версия проверятся"
	case NoMatter:
		return "версия не проверятся"
	}
	return "не проверятся"
}
