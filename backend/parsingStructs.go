package backend

import (
	"database/sql"
)

type Team struct {
	TeamID    sql.NullInt64  `db:"teamID" json:"teamID"`
	TeamName  string         `db:"teamName" json:"teamName"`
	Country   string         `db:"country" json:"country"`
	ShortName sql.NullString `db:"shortName" json:"shortName"`
}

type Transfer struct {
	TransferID   int           `db:"transferID" json:"transferID"`
	SeasonID     int           `db:"seasonID" json:"seasonID"`
	Date         string        `db:"date" json:"date"`
	Player       string        `db:"playerName" json:"player"`
	Team         string        `db:"teamName" json:"team"`
	Fee          int           `db:"fee" json:"fee"`
	PotentialFee sql.NullInt64 `db:"potentialFee" json:"potentialFee"`
	InTransfer   bool          `db:"transferIn" json:"inTransfer"`
	Loan         int           `db:"loan" json:"loan"`
	Free         int           `db:"free" json:"free"`
	// PlayerID     sql.NullInt64 `db:"playerID"`
}

type Match struct {
	SeasonID     int        `db:"seasonID" json:"seasonID"`
	SaveID       NullInt64  `db:"saveID" json:"saveID"`
	Season       NullString `db:"year" json:"year"`
	Date         string     `db:"date" json:"date"`
	Opposition   string     `db:"opponentName" json:"opposition"`
	Time         string
	Venue        string `db:"venue" json:"venue"`
	Stadium      string `db:"stadium" json:"stadium"`
	Competition  string `db:"competition" json:"competition"`
	GoalsAgainst int    `db:"goalsAgainst" json:"goalsAgainst"`
	GoalsFor     int    `db:"goalsFor" json:"goalsFor"`
	Result       string `db:"result" json:"result"`
	ExtraTime    int    `db:"extraTime" json:"extraTime"`
	Penalties    int    `db:"penalties" json:"penalties"`
}

type PlayerInfo struct {
	PlayerID int    `db:"playerID" json:"playerID"`
	SaveID   int    `db:"saveID" json:"saveID"`
	Name     string `db:"playerName" json:"playerName"`
	UID      int    `db:"uniqueID" json:"uniqueID"`
	Nat      string `db:"nationality" json:"nationality"`
	// SecNat string `db:"secondNationality"`
	SecNat   sql.NullString `db:"secondNationality" json:"secondNationality"`
	DOB      string         `db:"birthdate" json:"birthdate"`
	Position string         `db:"position" json:"position"`
}

type PlayerSeason struct {
	// Name string
	UID int
	// Season         string
	PlayerSeasonID int `db:"playerSeasonID" json:"playerSeasonID"`
	PlayerStats
	PlayerTechAttr
	PlayerMentalAttr
	PlayerPhysAttr
	PlayerGKAttr
}

type PlayerStats struct {
	Apps     int     `db:"starts" json:"starts"`
	Subs     int     `db:"subs" json:"subs"`
	Gls      int     `db:"goals" json:"goals"`
	Ast      int     `db:"assists" json:"assists"`
	PoM      int     `db:"playerOfTheMatch" json:"playerOfTheMatch"`
	PasP     int     `db:"passPerc" json:"passPerc"`
	Yel      int     `db:"yellowCards" json:"yellowCards"`
	Red      int     `db:"redCards" json:"redCards"`
	AvgRat   float32 `db:"avgRating" json:"avgRating"`
	Gwin     int     `db:"winPerc" json:"winPerc"`
	Mins     int     `db:"minutes" json:"minutes"`
	Shutouts int     `db:"shutouts" json:"shutouts"`
	Sv       int     `db:"savePerc" json:"savePerc"`
}

type PlayerTechAttr struct {
	Cor int `db:"cor" json:"cor"`
	Cro int `db:"cro" json:"cro"`
	Dri int `db:"dri" json:"dri"`
	Fin int `db:"fin" json:"fin"`
	Fir int `db:"fir" json:"fir"`
	Fre int `db:"fre" json:"fre"`
	Hea int `db:"hea" json:"hea"`
	Lon int `db:"lon" json:"lon"`
	LTh int `db:"lth" json:"lth"`
	Mar int `db:"mar" json:"mar"`
	Pas int `db:"pas" json:"pas"`
	Pen int `db:"pen" json:"pen"`
	Tck int `db:"tck" json:"tck"`
	Tec int `db:"tec" json:"tec"`
}

type PlayerMentalAttr struct {
	Agg int `db:"agg" json:"agg"`
	Ant int `db:"ant" json:"ant"`
	Bra int `db:"bra" json:"bra"`
	Cmp int `db:"cmp" json:"cmp"`
	Cnt int `db:"cnt" json:"cnt"`
	Dec int `db:"dec" json:"dec"`
	Det int `db:"det" json:"det"`
	Fla int `db:"fla" json:"fla"`
	Ldr int `db:"ldr" json:"ldr"`
	OtB int `db:"otb" json:"otb"`
	Pos int `db:"pos" json:"pos"`
	Tea int `db:"tea" json:"tea"`
	Vis int `db:"vis" json:"vis"`
	Wor int `db:"wor" json:"wor"`
}

type PlayerPhysAttr struct {
	Acc int `db:"acc" json:"acc"`
	Agi int `db:"agi" json:"agi"`
	Bal int `db:"bal" json:"bal"`
	Jum int `db:"jum" json:"jum"`
	Nat int `db:"nat" json:"nat"`
	Pac int `db:"pac" json:"pac"`
	Sta int `db:"sta" json:"sta"`
	Str int `db:"str" json:"str"`
}

type PlayerGKAttr struct {
	Aer int `db:"aer" json:"aer"`
	Cmd int `db:"cmd" json:"cmd"`
	Com int `db:"com" json:"com"`
	Ecc int `db:"ecc" json:"ecc"`
	Han int `db:"han" json:"han"`
	Kic int `db:"kic" json:"kic"`
	OvO int `db:"ovo" json:"ovo"`
	Pun int `db:"pun" json:"pun"`
	Ref int `db:"ref" json:"ref"`
	TRO int `db:"tro" json:"tro"`
	Thr int `db:"thr" json:"thr"`
}
