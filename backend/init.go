package backend

func StartBackend() {
	StartLogger()
	GetStoriesFromFile()
	OpenDB()
}
