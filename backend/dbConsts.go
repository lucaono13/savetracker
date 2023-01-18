package backend

const (
	AddSaveQ     = `INSERT INTO saves (saveName, managerName, gameVersion) VALUES (?, ?, ?)`
	CreateTables = `
	CREATE TABLE IF NOT EXISTS "competitions" (
		"competitionID"	INTEGER NOT NULL UNIQUE,
		"compName"	TEXT UNIQUE,
		"league"	INTEGER,
		"trophyName"	TEXT,
		PRIMARY KEY("competitionID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "saves" (
		"saveID"	INTEGER NOT NULL UNIQUE,
		"managerName"	TEXT NOT NULL,
		"gameVersion"	INTEGER NOT NULL,
		"saveName"	TEXT NOT NULL,
		PRIMARY KEY("saveID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "teams" (
		"teamID"	INTEGER NOT NULL UNIQUE,
		"teamName"	TEXT NOT NULL,
		"primaryColor"	TEXT NOT NULL,
		"secondaryColor"	TEXT NOT NULL,
		"tertiaryColor"	TEXT,
		"country"	TEXT NOT NULL,
		PRIMARY KEY("teamID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "teamsManaged" (
		"teamManagedID"	INTEGER NOT NULL UNIQUE,
		"saveID"	INTEGER NOT NULL,
		"teamID"	INTEGER NOT NULL,
		"startYear"	INTEGER NOT NULL,
		"endYear"	INTEGER,
		FOREIGN KEY("saveID") REFERENCES "saves"("saveID") ON DELETE CASCADE,
		PRIMARY KEY("teamManagedID" AUTOINCREMENT),
		FOREIGN KEY("teamID") REFERENCES "teams"("teamID") ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS "seasons" (
		"seasonID"	INTEGER NOT NULL UNIQUE,
		"teamManagedID"	INTEGER,
		"year"	INTEGER NOT NULL,
		FOREIGN KEY("teamManagedID") REFERENCES "teamsManaged"("teamManagedID") ON DELETE CASCADE,
		PRIMARY KEY("seasonID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "competitionsWon" (
		"competitionsWonID"	INTEGER NOT NULL,
		"seasonID"	INTEGER NOT NULL,
		"competitionID"	INTEGER NOT NULL,
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		FOREIGN KEY("competitionID") REFERENCES "competitions"("competitionID") ON DELETE CASCADE,
		PRIMARY KEY("competitionsWonID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "players" (
		"playerID"	INTEGER NOT NULL UNIQUE,
		"saveID"	INTEGER NOT NULL,
		"playerName"	TEXT,
		"position"	TEXT,
		"birthdate"	TEXT,
		"uniqueID"	INTEGER,
		FOREIGN KEY("saveID") REFERENCES "saves"("saveID") ON DELETE CASCADE,
		PRIMARY KEY("playerID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "playerSeason" (
		"playerSeasonID"	INTEGER NOT NULL UNIQUE,
		"playerID"	INTEGER NOT NULL,
		"seasonID"	INTEGER NOT NULL,
		FOREIGN KEY("playerID") REFERENCES "players"("playerID") ON DELETE CASCADE,
		PRIMARY KEY("playerSeasonID" AUTOINCREMENT),
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS "playerStats" (
		"playerStatID"	INTEGER NOT NULL UNIQUE,
		"playerSeasonID"	INTEGER NOT NULL,
		"minutes"	INTEGER NOT NULL,
		"appearances"	INTEGER,
		"goals"	INTEGER,
		"assists"	INTEGER,
		"yellowCards"	INTEGER,
		"redCards"	INTEGER,
		"avgRating"	REAL,
		FOREIGN KEY("playerSeasonID") REFERENCES "playerSeason"("playerSeasonID") ON DELETE CASCADE,
		PRIMARY KEY("playerStatID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "transfers" (
		"transferID"	INTEGER NOT NULL UNIQUE,
		"teamID"	INTEGER NOT NULL,
		"playerID"	INTEGER NOT NULL,
		"date"	TEXT NOT NULL,
		"fee"	INTEGER NOT NULL DEFAULT 0,
		"potentialFee"	INTEGER,
		"transferIn"	TEXT NOT NULL,
		"loan"	TEXT NOT NULL,
		"free"	TEXT,
		FOREIGN KEY("playerID") REFERENCES "players"("playerID") ON DELETE CASCADE,
		FOREIGN KEY("teamID") REFERENCES "teams"("teamID") ON DELETE CASCADE,
		PRIMARY KEY("transferID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "results" (
		"resultID"	INTEGER NOT NULL UNIQUE,
		"seasonID"	INTEGER NOT NULL,
		"opponentID"	INTEGER NOT NULL,
		"competitionID"	INTEGER NOT NULL,
		"awayScore"	INTEGER NOT NULL,
		"homeScore"	INTEGER NOT NULL,
		"date"	TEXT,
		"stadium"	TEXT,
		"venue"	TEXT,
		FOREIGN KEY("competitionID") REFERENCES "competitions"("competitionID") ON DELETE CASCADE,
		FOREIGN KEY("opponentID") REFERENCES "teams"("teamID") ON DELETE CASCADE,
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		PRIMARY KEY("resultID" AUTOINCREMENT)
	);
	`
	SingleSave = `SELECT * FROM saves where saveID=?`
)
