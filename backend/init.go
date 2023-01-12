package backend

func StartBackend(dbPath string) {
	StartLogger()
	OpenDB(dbPath)
}
