package backend

const (
	AddSaveQ     = `INSERT INTO saves (saveName, managerName, gameVersion) VALUES (?, ?, ?)`
	CreateTables = `
	CREATE TABLE IF NOT EXISTS "saves" (
		"saveID"	INTEGER NOT NULL UNIQUE,
		"managerName"	TEXT NOT NULL,
		"gameVersion"	INTEGER NOT NULL,
		"saveName"	TEXT NOT NULL,
		"saveImage"	TEXT,
		PRIMARY KEY("saveID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "teams" (
		"teamID"	INTEGER NOT NULL UNIQUE,
		"teamName"	TEXT NOT NULL,
		"country"	TEXT,
		"shortName"	TEXT,
		PRIMARY KEY("teamID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "players" (
		"playerID"	INTEGER NOT NULL UNIQUE,
		"saveID"	INTEGER NOT NULL,
		"playerName"	TEXT NOT NULL,
		"position"	TEXT,
		"birthdate"	TEXT NOT NULL,
		"uniqueID"	INTEGER NOT NULL,
		"nationality"	TEXT,
		"secondNationality"	TEXT,
		PRIMARY KEY("playerID" AUTOINCREMENT),
		FOREIGN KEY("saveID") REFERENCES "saves"("saveID") ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS "seasons" (
		"seasonID"	INTEGER NOT NULL UNIQUE,
		"teamID"	INTEGER NOT NULL,
		"year"	TEXT NOT NULL,
		"saveID"	INTEGER NOT NULL,
		FOREIGN KEY("saveID") REFERENCES "saves"("saveID") ON DELETE CASCADE,
		FOREIGN KEY("teamID") REFERENCES "teams"("teamID") ON DELETE CASCADE,
		PRIMARY KEY("seasonID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "transfers" (
		"transferID"	INTEGER NOT NULL UNIQUE,
		"teamName"	TEXT,
		"playerID"	INTEGER NOT NULL,
		"date"	TEXT NOT NULL,
		"fee"	INTEGER NOT NULL DEFAULT 0,
		"potentialFee"	INTEGER,
		"transferIn"	INTEGER NOT NULL,
		"loan"	INTEGER NOT NULL,
		"free"	INTEGER,
		"saveID"	INTEGER NOT NULL,
		PRIMARY KEY("transferID" AUTOINCREMENT),
		FOREIGN KEY("playerID") REFERENCES "players"("playerID") ON DELETE CASCADE,
		FOREIGN KEY("saveID") REFERENCES "saves"("saveID") ON DELETE CASCADE
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
		"minutes"	INTEGER,
		"starts"	INTEGER,
		"goals"	INTEGER,
		"assists"	INTEGER,
		"yellowCards"	INTEGER,
		"redCards"	INTEGER,
		"avgRating"	REAL,
		"subs"	INTEGER,
		"playerOfTheMatch"	INTEGER,
		"passPerc"	INTEGER,
		"winPerc"	INTEGER,
		"shutouts"	INTEGER,
		"savePerc"	INTEGER,
		PRIMARY KEY("playerStatID" AUTOINCREMENT),
		FOREIGN KEY("playerSeasonID") REFERENCES "playerSeason"("playerSeasonID") ON DELETE CASCADE
	);
	CREATE TABLE IF NOT EXISTS "results" (
		"resultID"	INTEGER NOT NULL UNIQUE,
		"seasonID"	INTEGER NOT NULL,
		"opponentID"	INTEGER NOT NULL,
		"competition"	TEXT NOT NULL,
		"awayScore"	INTEGER NOT NULL,
		"homeScore"	INTEGER NOT NULL,
		"date"	TEXT NOT NULL,
		"stadium"	TEXT NOT NULL,
		"venue"	TEXT NOT NULL,
		"result"	TEXT NOT NULL,
		"penalties"	INTEGER NOT NULL,
		"extraTime"	INTEGER NOT NULL,
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		FOREIGN KEY("opponentID") REFERENCES "teams"("teamID") ON DELETE CASCADE,
		PRIMARY KEY("resultID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "trophies" (
		"trophyWonID"	INTEGER NOT NULL UNIQUE,
		"seasonID"	INTEGER NOT NULL,
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		PRIMARY KEY("trophyWonID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "playerAttributes" (
		"playerAttrID"	INTEGER NOT NULL UNIQUE,
		"playerSeasonID"	INTEGER NOT NULL,
		"corners"	INTEGER NOT NULL,
		"crossing"	INTEGER NOT NULL,
		"dribbling"	INTEGER NOT NULL,
		"finishing"	INTEGER NOT NULL,
		"firstTouch"	INTEGER NOT NULL,
		"freeKicks"	INTEGER NOT NULL,
		"heading"	INTEGER NOT NULL,
		"longShots"	INTEGER NOT NULL,
		"longThrows"	INTEGER NOT NULL,
		"marking"	INTEGER NOT NULL,
		"passing"	INTEGER NOT NULL,
		"penalties"	INTEGER NOT NULL,
		"tackling"	INTEGER NOT NULL,
		"technique"	INTEGER NOT NULL,
		"aggression"	INTEGER NOT NULL,
		"anticipation"	INTEGER NOT NULL,
		"bravery"	INTEGER NOT NULL,
		"composure"	INTEGER NOT NULL,
		"concentration"	INTEGER NOT NULL,
		"decisions"	INTEGER NOT NULL,
		"determination"	INTEGER NOT NULL,
		"flair"	INTEGER NOT NULL,
		"leadership"	INTEGER NOT NULL,
		"offTheBall"	INTEGER NOT NULL,
		"positioning"	INTEGER NOT NULL,
		"teamwork"	INTEGER NOT NULL,
		"vision"	INTEGER NOT NULL,
		"workRate"	INTEGER NOT NULL,
		"acceleration"	INTEGER NOT NULL,
		"agility"	INTEGER NOT NULL,
		"balance"	INTEGER NOT NULL,
		"jumpingReach"	INTEGER NOT NULL,
		"naturalFitness"	INTEGER NOT NULL,
		"pace"	INTEGER NOT NULL,
		"stamina"	INTEGER NOT NULL,
		"strength"	INTEGER NOT NULL,
		"aerialReach"	INTEGER NOT NULL,
		"commandOfArea"	INTEGER NOT NULL,
		"communication"	INTEGER NOT NULL,
		"eccentricity"	INTEGER NOT NULL,
		"handling"	INTEGER NOT NULL,
		"kicking"	INTEGER NOT NULL,
		"oneOnOnes"	INTEGER NOT NULL,
		"punchingTendency"	INTEGER NOT NULL,
		"reflexes"	INTEGER NOT NULL,
		"rushingOutTendency"	INTEGER NOT NULL,
		"throwing"	INTEGER NOT NULL,
		FOREIGN KEY("playerSeasonID") REFERENCES "playerSeason"("playerSeasonID") ON DELETE CASCADE,
		PRIMARY KEY("playerAttrID" AUTOINCREMENT)
	);
	`
	AllSaves        = `SELECT * FROM saves`
	NewSave         = `INSERT INTO saves (saveName, managerName, gameVersion) VALUES (:saveName, :managerName, :gameVersion)`
	SingleSave      = `SELECT * FROM saves where saveID=?`
	SingleSaveImage = `SELECT saveImage FROM saves where saveID=?`
	SaveImageUpdate = `UPDATE saves SET saveImage=? WHERE saveID=?`
	AllTeams        = `SELECT * FROM teams`
)
