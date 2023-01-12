package backend

const (
	AddSaveQ = `INSERT INTO saves (managerName, gameVersion) VALUES (?, ?)`
)
