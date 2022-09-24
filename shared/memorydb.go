package shared

type MemoryDb map[string]ParsedLog

var repo MemoryDb

func NewMemoryDb() MemoryDb {
	if repo == nil {
		repo = make(MemoryDb)
	}
	return repo
}
