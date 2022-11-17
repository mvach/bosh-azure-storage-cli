package blob

type ExistenceState int64

const (
	Exists      ExistenceState = 0
	NotExisting                = 1
	Unknown                    = 2
)
