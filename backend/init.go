package backend

func StartBackend() {
	StartLogger()
	GetConfig()
	GetStoriesFromFile()
	OpenDB()
}
