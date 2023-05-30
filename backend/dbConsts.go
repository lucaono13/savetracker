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
		"currency"	TEXT NOT NULL,
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
		"teamName"	TEXT NOT NULL,
		"date"	TEXT NOT NULL,
		"fee"	INTEGER NOT NULL DEFAULT 0,
		"potentialFee"	INTEGER,
		"transferIn"	INTEGER NOT NULL,
		"loan"	INTEGER NOT NULL,
		"free"	INTEGER NOT NULL,
		"seasonID"	INTEGER NOT NULL,
		"playerName"	TEXT NOT NULL,
		PRIMARY KEY("transferID" AUTOINCREMENT),
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE
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
		"opponentName"	TEXT NOT NULL,
		"competition"	TEXT NOT NULL,
		"goalsAgainst"	INTEGER NOT NULL,
		"goalsFor"	INTEGER NOT NULL,
		"date"	TEXT NOT NULL,
		"stadium"	TEXT NOT NULL,
		"venue"	TEXT NOT NULL,
		"result"	TEXT NOT NULL,
		"penalties"	INTEGER NOT NULL,
		"extraTime"	INTEGER NOT NULL,
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID"),
		PRIMARY KEY("resultID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "trophies" (
		"trophyWonID"	INTEGER NOT NULL UNIQUE,
		"seasonID"	INTEGER NOT NULL,
		"competitionName"	TEXT NOT NULL,
		"trophyImage"	TEXT,
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
	NewSave         = `INSERT INTO saves (saveName, managerName, gameVersion, currency) VALUES (:saveName, :managerName, :gameVersion, :currency)`
	SingleSave      = `SELECT * FROM saves where saveID=?`
	SingleSaveImage = `SELECT saveImage FROM saves where saveID=?`
	SaveImageUpdate = `UPDATE saves SET saveImage=? WHERE saveID=?`
	AllTeams        = `SELECT * FROM teams`
	NewStats        = `INSERT INTO playerStats (playerSeasonID, minutes, starts, goals, assists, yellowCards, redCards, avgRating, subs, playerOfTheMatch, passPerc, winPerc, shutouts, savePerc) 
						VALUES (:playerSeasonID, :minutes, :starts, :goals, :assists, :yellowCards, :redCards, :avgRating, :subs, :playerOfTheMatch, :passPerc, :winPerc, :shutouts, :savePerc)`
	NewAttrs = `INSERT INTO playerAttributes (playerSeasonID, corners, crossing, dribbling, finishing, firstTouch, freeKicks, heading, longShots, longThrows, marking, passing, penalties, tackling, technique, 
						aggression, anticipation, bravery, composure, concentration, decisions, determination, flair, leadership, offTheBall, positioning, teamwork, vision, workRate,
						acceleration, agility, balance, jumpingReach, naturalFitness, pace, stamina, strength,
						aerialReach, commandOfArea, communication, eccentricity, handling, kicking, oneOnOnes, punchingTendency, reflexes, rushingOutTendency, throwing) VALUES
						(:playerSeasonID, :cor, :cro, :dri, :fin, :fir, :fre, :hea, :lon, :lth, :mar, :pas, :pen, :tck, :tec,
						:agg, :ant, :bra, :cmp, :cnt, :dec, :det, :fla, :ldr, :otb, :pos, :tea, :vis, :wor,
						:acc, :agi, :bal, :jum, :nat, :pac, :sta, :str,
						:aer, :cmd, :com, :ecc, :han, :kic, :ovo, :pun, :ref, :tro, :thr)`
	NewResult = `INSERT INTO results (seasonID, date, opponentName, venue, stadium, competition, goalsAgainst, goalsFor, result, extraTime, penalties) VALUES
					(:seasonID, :date, :opponentName, :venue, :stadium, :competition, :goalsAgainst, :goalsFor, :result, :extraTime, :penalties)`
	NewTransfer = `INSERT INTO transfers (seasonID, date, playerName, teamName, fee, potentialFee, transferIn, loan, free) VALUEs
					(:seasonID, :date, :playerName, :teamName, :fee, :potentialFee, :transferIn, :loan, :free)`
	NewTeam          = `INSERT INTO teams (teamName, shortName, country) VALUES (:teamName, :shortName, :country)`
	NewSeason        = `INSERT INTO seasons (teamID, saveID, year) VALUES (?, ?, ?)`
	NewPlayerSeason  = `INSERT INTO playerSeason (playerID, seasonID) VALUES (?, ?)`
	AddPlayer1Nat    = `INSERT INTO players (saveID, playerName, position, birthdate, uniqueID, nationality) VALUES (:saveID, :playerName, :position, :birthdate, :uniqueID, :nationality)`
	AddPlayer2Nat    = `INSERT INTO players (saveID, playerName, position, birthdate, uniqueID, nationality, secondNationality) VALUES (:saveID, :playerName, :position, :birthdate, :uniqueID, :nationality, :secondNationality)`
	SavePlayers      = `SELECT * FROM players WHERE saveID=?`
	AllPlayers       = `SELECT * FROM players`
	PlayerSeasons    = `SELECT * FROM playerSeason`
	SinglePlayer     = `SELECT * FROM players WHERE playerID=?`
	OnePlayerSeasons = `SELECT * FROM playerSeason WHERE `
	NewTrophy        = `INSERT INTO trophies VALUES (:seasonID, :competitionName)`
	SaveResults      = `SELECT saves.saveID,
							results.date,
							seasons.year,
							results.competition,
							results.stadium,
							results.venue,
							results.opponentName,
							results.result,
							results.goalsFor,
							results.goalsAgainst,
							results.penalties,
							results.extraTime
						FROM results
						INNER JOIN seasons ON results.seasonID = seasons.seasonID
						INNER JOIN saves ON seasons.saveID = saves.saveID
						WHERE saves.saveID = ?
						`
	GetCurrency   = `SELECT currency FROM saves WHERE saveID=?`
	SaveTransfers = `SELECT
							transfers.date,
							transfers.playerName,
							transfers.fee,
							transfers.potentialFee,
							transfers.teamName,
							transfers.free,
							transfers.loan,
							seasons.year
						FROM
							transfers
						INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
						INNER JOIN saves ON seasons.saveID = saves.saveID
						WHERE saves.saveID=? AND transfers.transferIn=?
						`
)
