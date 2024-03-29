package backend

const (
	AddSaveQ     = `INSERT INTO saves (saveName, managerName, gameVersion) VALUES (?, ?, ?)`
	CreateTables = `
	CREATE TABLE IF NOT EXISTS "saves" (
		"saveID"	INTEGER NOT NULL UNIQUE,
		"managerName"	TEXT NOT NULL,
		"gameVersion"	TEXT NOT NULL,
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
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		PRIMARY KEY("resultID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "trophies" (
		"trophyID"	INTEGER NOT NULL UNIQUE,
		"trophyName"	TEXT NOT NULL,
		"trophyImage"	INTEGER,
		PRIMARY KEY("trophyID" AUTOINCREMENT)
	);
	CREATE TABLE IF NOT EXISTS "trophiesWon" (
		"trophyWonID"	INTEGER NOT NULL UNIQUE,
		"seasonID"	INTEGER NOT NULL,
		"trophyID"	TEXT NOT NULL,
		PRIMARY KEY("trophyWonID" AUTOINCREMENT),
		FOREIGN KEY("seasonID") REFERENCES "seasons"("seasonID") ON DELETE CASCADE,
		FOREIGN KEY("trophyID") REFERENCES "trophies"("trophyID") ON DELETE CASCADE
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
	AllSaves          = `SELECT * FROM saves`
	NewSave           = `INSERT INTO saves (saveName, managerName, gameVersion, currency) VALUES (:saveName, :managerName, :gameVersion, :currency)`
	SingleSave        = `SELECT * FROM saves where saveID=?`
	SingleSaveImage   = `SELECT saveImage FROM saves where saveID=?`
	SaveImageUpdate   = `UPDATE saves SET saveImage=? WHERE saveID=?`
	TrophyImageUpdate = `UPDATE trophies SET trophyImage=? WHERE trophyID=?`
	AllTeams          = `SELECT * FROM teams`
	NewStats          = `INSERT INTO playerStats (playerSeasonID, minutes, starts, goals, assists, yellowCards, redCards, avgRating, subs, playerOfTheMatch, passPerc, winPerc, shutouts, savePerc) 
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
	NewTeam         = `INSERT INTO teams (teamName, shortName, country) VALUES (:teamName, :shortName, :country)`
	NewSeason       = `INSERT INTO seasons (teamID, saveID, year) VALUES (?, ?, ?)`
	NewPlayerSeason = `INSERT INTO playerSeason (playerID, seasonID) VALUES (?, ?)`
	AddPlayer1Nat   = `INSERT INTO players (saveID, playerName, position, birthdate, uniqueID, nationality) VALUES (:saveID, :playerName, :position, :birthdate, :uniqueID, :nationality)`
	AddPlayer2Nat   = `INSERT INTO players (saveID, playerName, position, birthdate, uniqueID, nationality, secondNationality) VALUES (:saveID, :playerName, :position, :birthdate, :uniqueID, :nationality, :secondNationality)`
	SavePlayers     = `SELECT * FROM players WHERE saveID=?`
	AllPlayers      = `SELECT * FROM players`
	PlayerSeasons   = `SELECT * FROM playerSeason`
	// SinglePlayer    = `SELECT * FROM players WHERE playerID=?`
	// OnePlayerSeasons = `SELECT * FROM playerSeason WHERE `
	NewTrophyWon = `INSERT INTO trophiesWon (seasonID, trophyID) VALUES (?, ?)`
	NewTrophy    = `INSERT INTO trophies (trophyName) VALUES (?)`
	// AllTrophies  = `SELECT * FROM trophies`
	SaveResults = `
		SELECT saves.saveID, results.date, seasons.year, results.competition, results.stadium, results.venue, 
			results.opponentName, results.result, results.goalsFor, results.goalsAgainst, results.penalties, results.extraTime
		FROM results INNER JOIN seasons ON results.seasonID = seasons.seasonID 
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ?
	`
	GetCurrency   = `SELECT currency FROM saves WHERE saveID=?`
	SaveTransfers = `	
		SELECT transfers.date, transfers.playerName, transfers.fee, transfers.potentialFee, transfers.teamName, transfers.free, transfers.loan, seasons.year
		FROM transfers 
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID 
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID=? AND transfers.transferIn=?
	`
	SaveOutfieldPlayers = `
		SELECT players.playerID, players.playerName, players.position, teams.teamName, seasons.year, playerStats.minutes, playerStats.starts,
			playerStats.subs, playerStats.goals, playerStats.assists, playerStats.yellowCards, playerStats.redCards, playerStats.avgRating, 
			playerStats.playerOfTheMatch, playerStats.passPerc, playerStats.winPerc, playerStats.shutouts, playerStats.savePerc					
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID 
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID = ? AND players.position != "GK"
	`
	SaveGoaliePlayers = `
		SELECT players.playerID, players.playerName, players.position, teams.teamName, seasons.year, playerStats.minutes,
			playerStats.starts, playerStats.subs, playerStats.goals, playerStats.assists, playerStats.yellowCards, playerStats.redCards, 
			playerStats.avgRating, playerStats.playerOfTheMatch, playerStats.passPerc, playerStats.winPerc, playerStats.shutouts, playerStats.savePerc					
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID 
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID = ? AND players.position = "GK"
	`
	TotalsSaveOutfieldPlayers = `
		SELECT players.playerID, players.playerName, players.position, teams.teamName, COUNT(seasons.year) numYears,
			SUM(playerStats.minutes) totMins, SUM(playerStats.starts) totStarts, SUM(playerStats.subs) totSubs, SUM(playerStats.goals) totGoals, 
			SUM(playerStats.assists) totAssists, SUM(playerStats.yellowCards) totYellow, SUM(playerStats.redCards) totRed, AVG(playerStats.avgRating) avgRating,
			SUM(playerStats.playerOfTheMatch) totPOM, AVG(playerStats.passPerc) avgPass, AVG(playerStats.winPerc) avgWin,
			SUM(playerStats.shutouts) totShutouts, AVG(playerStats.savePerc) avgSaveP	
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID 
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID = ? AND players.position != "GK"
		GROUP BY players.playerID
	`
	TotalsSaveGoalies = `
		SELECT players.playerID, players.playerName, players.position, teams.teamName, COUNT(seasons.year) numYears,
			SUM(playerStats.minutes) totMins, SUM(playerStats.starts) totStarts, SUM(playerStats.subs) totSubs, SUM(playerStats.goals) totGoals, 
			SUM(playerStats.assists) totAssists, SUM(playerStats.yellowCards) totYellow, SUM(playerStats.redCards) totRed, AVG(playerStats.avgRating) avgRating,
			SUM(playerStats.playerOfTheMatch) totPOM, AVG(playerStats.passPerc) avgPass, AVG(playerStats.winPerc) avgWin, SUM(playerStats.shutouts) totShutouts,
			AVG(playerStats.savePerc) avgSaveP
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID 
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID = ? AND players.position = "GK"
		GROUP BY players.playerID
	`
	Top5Gls = `
		SELECT players.playerID, players.playerName, SUM(playerStats.goals) goals
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID=? 
		GROUP BY players.playerID 
		ORDER BY goals DESC
		LIMIT 5
	`
	Top5Asts = `
		SELECT players.playerID, players.playerName, SUM(playerStats.assists) assists
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID 
		WHERE saves.saveID=?
		GROUP BY players.playerID 
		ORDER BY assists DESC
		LIMIT 5
	`
	Top5Apps = `
		SELECT players.playerID, players.playerName, SUM(playerStats.starts) starts
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID 
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID 
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID=? 
		GROUP BY players.playerID 
		ORDER BY starts DESC
		LIMIT 5
	`
	TopRat = `
		SELECT players.playerID, players.playerName, AVG(playerStats.avgRating) avgRating
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE saves.saveID=? 
		GROUP BY players.playerID 
		ORDER BY avgRating DESC
		LIMIT 5
	`
	NumSaves         = `SELECT COUNT(1) FROM saves`
	NumSeasons       = `SELECT COUNT(1) FROM seasons`
	NumSeasonsToSave = `SELECT COUNT(1) FROM seasons INNER JOIN saves ON seasons.saveID = saves.saveID WHERE saves.saveID=?`
	MostTransfers    = `
		SELECT transfers.teamName, saves.currency, AVG(transfers.fee) avgFee, SUM(transfers.fee) totFee, COUNT(transfers.fee) numTransfers
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ?  AND transfers.loan = false AND transfers.free = false
		GROUP BY transfers.teamName
		ORDER BY numTransfers DESC
		LIMIT 5
	`
	MostLoans = `
		SELECT transfers.teamName, saves.currency,
			SUM(transfers.fee) totFee, COUNT(transfers.fee) numLoans
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ? AND transfers.loan = true AND transfers.free = false
		GROUP BY transfers.teamName
		ORDER BY numLoans DESC
		LIMIT 5
	`
	AvgTransfersOut = `
		SELECT AVG(transfers.fee)
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ? AND transfers.transferIn = false AND transfers.free = false 
	`
	AvgTransfersIn = `
		SELECT AVG(transfers.fee)
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ? AND transfers.transferIn = true AND transfers.free = false 
	`
	SaveTrophies = `
		SELECT trophies.trophyID, trophies.trophyName, trophies.trophyImage, seasons.year
		FROM trophiesWon
		INNER JOIN trophies ON trophiesWon.trophyID = trophies.trophyID
		INNER JOIN seasons ON trophiesWon.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ?
	`
	SinglePlayer = `
		SELECT 
			seasons.year,
			teams.teamName,
			playerSeason.playerSeasonID,
			playerAttributes.corners cor,
			playerAttributes.crossing cro,
			playerAttributes.dribbling dri,
			playerAttributes.finishing fin,
			playerAttributes.firstTouch fir,
			playerAttributes.freeKicks fre,
			playerAttributes.heading hea,
			playerAttributes.longShots lon,
			playerAttributes.longThrows lth,
			playerAttributes.marking mar,
			playerAttributes.passing pas,
			playerAttributes.penalties pen,
			playerAttributes.tackling tck,
			playerAttributes.technique tec,
			playerAttributes.aggression agg,
			playerAttributes.anticipation ant,
			playerAttributes.bravery bra,
			playerAttributes.composure cmp,
			playerAttributes.concentration cnt,
			playerAttributes.decisions dec,
			playerAttributes.determination det,
			playerAttributes.flair fla,
			playerAttributes.leadership ldr,
			playerAttributes.offTheBall otb,
			playerAttributes.positioning pos,
			playerAttributes.teamwork tea,
			playerAttributes.vision vis,
			playerAttributes.workRate wor,
			playerAttributes.acceleration acc,
			playerAttributes.agility agi,
			playerAttributes.balance bal,
			playerAttributes.jumpingReach jum,
			playerAttributes.naturalFitness nat,
			playerAttributes.pace pac,
			playerAttributes.stamina sta,
			playerAttributes.strength str,
			playerAttributes.aerialReach aer,
			playerAttributes.commandOfArea cmd,
			playerAttributes.communication com,
			playerAttributes.eccentricity ecc,
			playerAttributes.handling han,
			playerAttributes.kicking kic,
			playerAttributes.oneOnOnes ovo,
			playerAttributes.punchingTendency pun,
			playerAttributes.reflexes ref,
			playerAttributes.rushingOutTendency tro,
			playerAttributes.throwing thr,
			playerStats.minutes,
			playerStats.starts,
			playerStats.subs,
			playerStats.goals,
			playerStats.assists,
			playerStats.yellowCards,
			playerStats.redCards,
			playerStats.avgRating,
			playerStats.playerOfTheMatch,
			playerStats.passPerc,
			playerStats.winPerc,
			playerStats.shutouts,
			playerStats.savePerc
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN playerAttributes ON playerAttributes.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN seasons ON playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		INNER JOIN saves ON players.saveID = saves.saveID
		WHERE players.playerID = ?
	`
	SinglePlayerSumsAvgs = `
		SELECT
			players.playerID,
			saves.saveID,
			saves.saveName,
			COUNT(seasons.year) years,
			saves.gameVersion,
			teams.teamName,
			players.playerName,
			players.uniqueID,
			players.birthdate,
			players.nationality,
			players.secondNationality,
			players.position, 
			AVG(playerStats.minutes) avgMin,
			SUM(playerStats.minutes) totMin,
			AVG(playerStats.starts) avgStart,
			SUM(playerStats.starts) totStart,
			AVG(playerStats.subs) avgSubs,
			SUM(playerStats.subs) totSubs,
			AVG(playerStats.goals) avgGls,
			SUM(playerStats.goals) totGls,
			AVG(playerStats.assists) avgAst,
			SUM(playerStats.assists) totAst,
			AVG(playerStats.yellowCards) avgYel,
			SUM(playerStats.yellowCards) totYel,
			AVG(playerStats.redCards) avgRed,
			SUM(playerStats.redCards) totRed,
			AVG(playerStats.avgRating) avgRat,
			AVG(playerStats.playerOfTheMatch) avgPOM,
			SUM(playerStats.playerOfTheMatch) totPOM,
			AVG(playerStats.passPerc) avgPasP,
			AVG(playerStats.winPerc) avgWinP,
			AVG(playerStats.shutouts) avgShutouts,
			SUM(playerStats.shutouts) totShutouts,
			AVG(playerStats.savePerc) avgSaveP
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN playerAttributes ON playerAttributes.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN seasons ON playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		INNER JOIN saves ON players.saveID = saves.saveID
		WHERE players.playerID = ?
	`
	AllAvgTransfersOut = `
		SELECT AVG(transfers.fee)
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE transfers.transferIn = false AND transfers.free = false 
	`
	AllAvgTransfersIn = `
		SELECT AVG(transfers.fee)
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE transfers.transferIn = true AND transfers.free = false 
	`
	AllTrophies = `
		SELECT trophies.trophyID, saves.saveName, trophies.trophyName, trophies.trophyImage, seasons.year
		FROM trophiesWon
		INNER JOIN trophies ON trophiesWon.trophyID = trophies.trophyID
		INNER JOIN seasons ON trophiesWon.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
	`
	AllMostTransfers = `
		SELECT transfers.teamName, saves.currency,
			AVG(transfers.fee) avgFee, SUM(transfers.fee) totFee, COUNT(transfers.fee) numTransfers
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE transfers.loan = false AND transfers.free = false
		GROUP BY transfers.teamName
		ORDER BY numTransfers DESC
		LIMIT 5
	`
	AllMostLoans = `
		SELECT transfers.teamName, saves.currency,
			SUM(transfers.fee) totFee, COUNT(transfers.fee) numLoans
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE transfers.loan = true AND transfers.free = false
		GROUP BY transfers.teamName
		ORDER BY numLoans DESC
		LIMIT 5
	`
	AllTopRat = `
		SELECT players.playerID, players.playerName, AVG(playerStats.avgRating) avgRating
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		GROUP BY players.playerID
		ORDER BY avgRating DESC
		LIMIT 5
	`
	AllSaveResults = `
		SELECT saves.saveID, results.date, seasons.year, saves.saveName, saves.gameVersion, results.competition,
			results.stadium, results.venue, results.opponentName, results.result, results.goalsFor,
			results.goalsAgainst, results.penalties, results.extraTime
		FROM results
		INNER JOIN seasons ON results.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
	`
	AllSaveTransfers = `	
		SELECT transfers.date, transfers.playerName, transfers.fee, transfers.potentialFee,
			transfers.teamName, transfers.free, transfers.loan, seasons.year, saves.saveName, saves.currency
		FROM transfers
		INNER JOIN seasons ON transfers.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE transfers.transferIn=?
	`
	AllSaveOutfieldPlayers = `
		SELECT players.playerID, players.playerName, players.position, saves.saveName, teams.teamName,
			seasons.year, playerStats.minutes, playerStats.starts, playerStats.subs, playerStats.goals, 
			playerStats.assists, playerStats.yellowCards, playerStats.redCards, playerStats.avgRating, 
			playerStats.playerOfTheMatch, playerStats.passPerc, playerStats.winPerc, playerStats.shutouts,
			playerStats.savePerc			
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE players.position != "GK"
	`
	AllSaveGoaliePlayers = `
		SELECT players.playerID, players.playerName, players.position, saves.saveName,
			teams.teamName, seasons.year, playerStats.minutes, playerStats.starts, playerStats.subs,
			playerStats.goals, playerStats.assists, playerStats.yellowCards, playerStats.redCards, playerStats.avgRating,
			playerStats.playerOfTheMatch, playerStats.passPerc, playerStats.winPerc, playerStats.shutouts, playerStats.savePerc			
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE players.position = "GK"
	`
	AllTotalsSaveOutfieldPlayers = `
		SELECT players.playerID, players.playerName, players.position, saves.saveName, teams.teamName,
			COUNT(seasons.year) numYears, SUM(playerStats.minutes) totMins, SUM(playerStats.starts) totStarts, SUM(playerStats.subs) totSubs,
			SUM(playerStats.goals) totGoals, SUM(playerStats.assists) totAssists, SUM(playerStats.yellowCards) totYellow, SUM(playerStats.redCards) totRed,
			AVG(playerStats.avgRating) avgRating, SUM(playerStats.playerOfTheMatch) totPOM, AVG(playerStats.passPerc) avgPass, AVG(playerStats.winPerc) avgWin,
			SUM(playerStats.shutouts) totShutouts, AVG(playerStats.savePerc) avgSaveP
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE players.position != "GK"
		GROUP BY players.playerID
	`
	AllTotalsSaveGoalies = `
		SELECT players.playerID, players.playerName, players.position, saves.saveName, teams.teamName,
			COUNT(seasons.year) numYears, SUM(playerStats.minutes) totMins, SUM(playerStats.starts) totStarts, SUM(playerStats.subs) totSubs,
			SUM(playerStats.goals) totGoals, SUM(playerStats.assists) totAssists, SUM(playerStats.yellowCards) totYellow, SUM(playerStats.redCards) totRed,
			AVG(playerStats.avgRating) avgRating, SUM(playerStats.playerOfTheMatch) totPOM, AVG(playerStats.passPerc) avgPass, AVG(playerStats.winPerc) avgWin,
			SUM(playerStats.shutouts) totShutouts, AVG(playerStats.savePerc) avgSaveP
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		WHERE players.position = "GK"
		GROUP BY players.playerID
	`
	AllTop5Gls = `
		SELECT players.playerID, players.playerName, SUM(playerStats.goals) goals
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		GROUP BY players.playerID
		ORDER BY goals DESC
		LIMIT 5
	`
	AllTop5Asts = `
		SELECT players.playerID, players.playerName, SUM(playerStats.assists) assists
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		GROUP BY players.playerID
		ORDER BY assists DESC
		LIMIT 5
	`
	AllTop5Apps = `
		SELECT players.playerID, players.playerName, SUM(playerStats.starts) starts
		FROM playerStats
		INNER JOIN playerSeason ON playerStats.playerSeasonID = playerSeason.playerSeasonID
		INNER JOIN players ON playerSeason.playerID = players.playerID
		INNER JOIN saves ON players.saveID = saves.saveID
		INNER JOIN seasons on playerSeason.seasonID = seasons.seasonID
		INNER JOIN teams on seasons.teamID = teams.teamID
		GROUP BY players.playerID
		ORDER BY starts DESC
		LIMIT 5
	`
	DeleteSaveQ = `
		DELETE
		FROM saves
		WHERE saveID = ?
	`
	SavePlayerSeasons = `
		SELECT COUNT(1)
		FROM playerSeason
		INNER JOIN seasons ON playerSeason.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
		WHERE saves.saveID = ?
	`
	TotalPlayerSeasons = `
		SELECT COUNT(1)
		FROM playerSeason
		INNER JOIN seasons ON playerSeason.seasonID = seasons.seasonID
		INNER JOIN saves ON seasons.saveID = saves.saveID
	`
)
