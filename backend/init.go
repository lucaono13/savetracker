package backend

// type Backend struct {
// }

func StartBackend(dbPath string) {
	StartLogger()
	OpenDB(dbPath)
}
