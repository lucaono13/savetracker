package backend

import (
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

func readHtmlFromFile(webPage string) (string, error) {
	content, err := os.ReadFile(webPage)
	// content, err := ioutil.ReadFile(webPage)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ParseTransfers(filePath string, teamName string) ([]Transfer, []Transfer, error) {

	text, err := readHtmlFromFile(filePath)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}

	// Create Tokenizer to parse through HTML file
	tkn := html.NewTokenizer(strings.NewReader(text))

	// Initialize Slice, Struct and variables
	inTransfers := []Transfer{}
	outTransfers := []Transfer{}
	transferRow := []string{}
	transfer := Transfer{Loan: 0, Free: 0}
	var isTd, isTh bool
	var isInOrOut string
	var n, headerCol, col int

	// Iterate through HTML file
	for {
		tt := tkn.Next()

		switch {
		case tt == html.ErrorToken:
			// If error is EOF, then return
			if tkn.Err().Error() == "EOF" {
				Logger.Info().Timestamp().Msg("Transfers successfully parsed.")
				return inTransfers, outTransfers, nil
			}
			// fmt.Println(tkn.Err())
			Logger.Error().Timestamp().Msg(tkn.Err().Error())
			return nil, nil, err
		// Check what start tag is being used
		case tt == html.StartTagToken:

			t := tkn.Token()
			isTd = t.Data == "td"
			isTh = t.Data == "th"
			if t.Data == "tr" {
				col = 0
			}

		case tt == html.TextToken:

			t := tkn.Token()

			// Check if its a header column as well as if the data to be parsed is the transfers in or out section
			// Transfer in section: "From"
			// Transfer out section: "To"
			if isTh {
				headerCol++
				if headerCol == 3 {
					isInOrOut = t.Data
				}
				if headerCol == 4 {
					headerCol = 0
				}

			}
			// Get data from each cell
			if isTd {
				col++
				data := strings.TrimSpace(t.Data)
				transferRow = append(transferRow, data)
				switch col {
				// First column is always date
				case 1:
					if data == "" {
						break
					}
					// Parse date
					transfer.Date = data
				// Second column is always Player name
				case 2:
					transfer.Player = data
				// Third column is always team
				case 3:
					transfer.Team = data
				// Fourth column is always fee associated with transfer fee/type of transfer if Loan or Free
				case 4:
					var numberSize, fee string
					if strings.Contains(data, "Loan") {
						transfer.Loan = 1
					} else {
						transfer.Loan = 0
					}
					// transfer.Loan = strings.Contains(data, "Loan")

					// If data is just Free
					if data == "Free" {
						transfer.Fee = 0
						transfer.Free = 1
						break
					}
					if data == "Undisclosed" {
						transfer.Fee = 0
						break
					}
					// If data is just Loan or an empty String break switch
					if data == "Loan" || data == "" {
						break
					}
					// Parse Loan fee
					if transfer.Loan == 1 && strings.Contains(data, "Total Fee") {
						fee = strings.ReplaceAll(data, "Loan - $", "")
						fee = strings.ReplaceAll(fee, " Total Fee", "")
						numberSize = string(fee[len(fee)-1])
						floatFee, err := strconv.ParseFloat(fee[0:len(fee)-1], 64)
						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
							// log.Fatal().Err(err)
						}
						if numberSize == "K" {
							transfer.Fee = int(floatFee * 1000)
						} else if numberSize == "M" {
							transfer.Fee = int(floatFee * 1000000)
						} else {
							transfer.Fee = int(floatFee)
						}
						break
					}
					// Parse Normal Fee
					fee = strings.Split(data, " ")[0]
					fee = strings.ReplaceAll(fee, "$", "")
					numberSize = string(fee[len(fee)-1])
					floatFee, err := strconv.ParseFloat(fee[0:len(fee)-1], 64)
					if err != nil {
						Logger.Error().Timestamp().Msg(err.Error())
						return nil, nil, err
						// log.Fatal().Err(err)
					}
					if numberSize == "K" {
						transfer.Fee = int(floatFee * 1000)
					} else if numberSize == "M" {
						transfer.Fee = int(floatFee * 1000000)
					} else {
						transfer.Fee = int(floatFee)
					}

					// Parsing Potential Fee
					if strings.Contains(data, "(") {
						fee = strings.ReplaceAll(data, "(", "")
						fee = strings.ReplaceAll(fee, ")", "")
						fee = strings.Split(fee, " ")[1]
						numberSize := string(fee[len(fee)-1])

						fee = strings.ReplaceAll(fee, "$", "")
						floatFee, err := strconv.ParseFloat(fee[0:len(fee)-1], 64)

						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
							// log.Fatal().Err(err)
						}

						if numberSize == "K" {
							transfer.PotentialFee.Int64 = int64(floatFee * 1000)

						} else if numberSize == "M" {
							transfer.PotentialFee.Int64 = int64(floatFee * 1000000)
						} else {
							transfer.PotentialFee.Int64 = int64(floatFee)
						}
						transfer.PotentialFee.Valid = true
					}

				}
				// Sets if the transfer is a transfer in or out
				transfer.InTransfer = (isInOrOut == "From")
				n++
			}
			// End of table row
			if (isTd || isTh) && n%4 == 0 {
				// fmt.Printf("%+v", transfer)
				var rowLen int
				// Makes sure row is not empty
				for _, v := range transferRow {
					rowLen += len(v)
				}
				// Add to slice to be exported
				if len(transferRow) > 0 && rowLen > 0 {
					// If the transfer out team does not equal Team's name, then append to transfers in slice
					// Else if the transfer team in is not the team name, then append to transfers out slice
					if isInOrOut == "From" && !strings.Contains(transfer.Team, teamName) {
						inTransfers = append(inTransfers, transfer)
					} else if isInOrOut == "To" && !strings.Contains(transfer.Team, teamName) {
						outTransfers = append(outTransfers, transfer)
					}
				}
				// fmt.Println()
				transfer = Transfer{Loan: 0, Free: 0}
				transferRow = nil
			}

			isTd = false
			isTh = false

		}
	}
}

func ParseSchedule(filePath string) ([]Match, error) {
	text, err := readHtmlFromFile(filePath)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, err
	}
	tkn := html.NewTokenizer(strings.NewReader(text))
	// matches := []Match{}
	match := Match{ExtraTime: 0, Penalties: 0}
	colNames := [6]string{"Date", "Opposition", "Time", "Venue", "Stadium", "Competition"}
	// var isTd bool
	// var col int
	// var homeTeam string
	var (
		matches  []Match
		isTd     bool
		col      int
		homeTeam string
	)

	// Weird Go shenanigans going on, please ignore
	_ = fmt.Sprint(homeTeam)
	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			if tkn.Err().Error() == "EOF" {
				Logger.Info().Timestamp().Msg("Schedule successfully parsed.")
				homeTeam = ""
				return matches, nil
			}
			Logger.Error().Timestamp().Msg(tkn.Err().Error())
			return nil, err
			// return nil, tkn.Err()
		case tt == html.StartTagToken:
			t := tkn.Token()
			// fmt.Println(t.Data, col)
			isTd = t.Data == "td"
			// isTh = t.Data == "th"
			if t.Data == "tr" {
				col = 0
			}
		case tt == html.TextToken:
			t := tkn.Token()
			if isTd {
				col++
				data := strings.TrimSpace(t.Data)
				switch col {
				case 1, 2, 3, 4, 5, 6:
					if data == "" {
						break
					}
					reflect.ValueOf(&match).Elem().FieldByName(colNames[col-1]).SetString(data)
				case 7:
					if data == "" {
						break
					}
					// homeTeam = data
				case 8:
					if data == "" {

						break
					}
					if strings.Contains(data, "e") {
						// Extra time edge case (shown in schedule file as 1-2e or e2-1)
						match.ExtraTime = 1
						data = strings.ReplaceAll(data, "e", "")
					} else if strings.Contains(data, "p") {
						match.Penalties = 1
					}
					goals := strings.Split(data, "-")
					// fmt.Println("[Match]", match)

					// Edge case for penalties (shown in schedule results as 1-1p or p1-1)
					// Remove the p in the string if it is in there for int parsing
					var whoWonPens string = ""
					if strings.Contains(goals[0], "p") {
						whoWonPens = "H"
						goals[0] = strings.ReplaceAll(goals[0], "p", "")
					} else if strings.Contains(goals[1], "p") {
						whoWonPens = "A"
						goals[1] = strings.ReplaceAll(goals[1], "p", "")
					}
					homeGoals, err := strconv.ParseInt(strings.TrimSpace(goals[0]), 10, 0)
					if err != nil {
						Logger.Error().Timestamp().Msg(err.Error())
						return nil, err
						// log.Fatal().Msg(err.Error())
					}
					awayGoals, err := strconv.ParseInt(strings.TrimSpace(goals[1]), 10, 0)
					if err != nil {
						Logger.Error().Timestamp().Msg(err.Error())
						return nil, err
						// log.Fatal().Msg(err.Error())
					}
					// match.HomeGoals = int(homeGoals)
					// match.AwayGoals = int(awayGoals)
					switch match.Venue {
					case "H":
						match.GoalsAgainst = int(awayGoals)
						match.GoalsFor = int(homeGoals)
						if homeGoals > awayGoals {
							match.Result = "W"
						} else if awayGoals > homeGoals {
							match.Result = "L"
						} else {
							if match.Penalties == 1 {
								if whoWonPens == "A" {
									match.Result = "L"
								} else {
									match.Result = "W"
								}
							} else {
								match.Result = "D"
							}
						}
					case "A":
						match.GoalsAgainst = int(homeGoals)
						match.GoalsFor = int(awayGoals)
						if homeGoals > awayGoals {
							match.Result = "L"
						} else if awayGoals > homeGoals {
							match.Result = "W"
						} else {
							if match.Penalties == 1 {
								if whoWonPens == "A" {
									match.Result = "W"
								} else {
									match.Result = "L"
								}
							} else {
								match.Result = "D"
							}
							// match.Result = "D"
						}
					case "N":
						// NOTE: for neutral games, goals in the home side are the goals scored by the team that that is capturing the schedule
						match.GoalsAgainst = int(awayGoals)
						match.GoalsFor = int(homeGoals)
						if match.Penalties == 0 {
							if homeGoals > awayGoals {
								match.Result = "W"
							} else if awayGoals > homeGoals {
								match.Result = "L"
							}
						} else {
							if whoWonPens == "H" {
								match.Result = "W"
							} else {
								match.Result = "L"
							}
						}
					}

				}
			}
			// fmt.Println(match)
			if col == 8 && len(match.Opposition) != 0 {
				// fmt.Println(match)
				matches = append(matches, match)
				match = Match{ExtraTime: 0, Penalties: 0}
				col = 0
			}

			isTd = false
		}
	}
	// return matches, nil
}

func ParseStats(path string, season string) ([]PlayerInfo, []PlayerSeason, error) {
	text, err := readHtmlFromFile(path)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return nil, nil, err
	}
	tkn := html.NewTokenizer(strings.NewReader(text))

	var (
		isTd, isTh   bool
		col          int
		allInfo      []PlayerInfo
		playerInfo   PlayerInfo
		allSeasons   []PlayerSeason
		playerSeason PlayerSeason
		colNames     []string
	)

	for {
		tt := tkn.Next()
		switch {
		case tt == html.ErrorToken:
			if tkn.Err().Error() == "EOF" {
				Logger.Info().Timestamp().Msg("Player Stats and Attributes successfully parsed.")
				return allInfo, allSeasons, nil
			}
			Logger.Error().Timestamp().Msg(tkn.Err().Error())
			return nil, nil, err
		case tt == html.StartTagToken:
			t := tkn.Token()
			isTd = t.Data == "td"
			isTh = t.Data == "th"
			if t.Data == "tr" {
				col = 0
			}
		case tt == html.TextToken:
			t := tkn.Token()
			if isTh {
				data := strings.ReplaceAll(t.Data, " ", "")
				data = strings.ReplaceAll(data, "%", "")
				if data == "1v1" {
					colNames = append(colNames, "OvO")
				} else {
					colNames = append(colNames, data)
				}
			}

			if isTd {
				col++
				data := strings.TrimSpace(t.Data)
				switch col {
				case 1, 3, 6:
					reflect.ValueOf(&playerInfo).Elem().FieldByName(colNames[col-1]).SetString(data)
				case 2:
					// Player's Unique ID
					uid, err := strconv.ParseInt(data, 10, 0)
					if err != nil {
						Logger.Error().Timestamp().Msg(err.Error())
						return nil, nil, err
					}
					playerInfo.UID = int(uid)
				case 4:
					// If they have a 2nd Nationality
					if len(data) == 3 {
						playerInfo.SecNat.String = data
						playerInfo.SecNat.Valid = true
					}
				case 8, 9, 10, 12, 13, 16, 17:
					// Getting the player's season stat
					if data != "-" {
						data = strings.ReplaceAll(data, ",", "")
						stat, err := strconv.ParseInt(data, 10, 0)
						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
						}
						reflect.ValueOf(&playerSeason).Elem().FieldByName(colNames[col-1]).SetInt(stat)
					}
				case 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32,
					33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46,
					47, 48, 49, 50, 51, 52, 53, 54,
					55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65:
					// Getting the players technical attributes
					attr, err := strconv.ParseInt(data, 10, 0)
					if err != nil {
						Logger.Error().Timestamp().Msg(err.Error())
						return nil, nil, err
					}
					reflect.ValueOf(&playerSeason).Elem().FieldByName(colNames[col-1]).SetInt(attr)
				case 5:
					// Getting player's birthday (just the date no age)
					r := regexp.MustCompile(`[(][0-9]{2} years old[)]`)
					dob := r.ReplaceAllString(data, "")
					playerInfo.DOB = dob
				case 7:
					// Gets the player's appearances (both starting and subs)
					if data != "-" {
						apps := strings.Split(data, "(")
						apps[0] = strings.TrimSpace(apps[0])
						starts, err := strconv.ParseInt(apps[0], 10, 0)
						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
						}
						playerSeason.Apps = int(starts)
						if len(apps) > 1 {
							apps[1] = strings.TrimSpace(apps[1])
							subs, err := strconv.ParseInt(strings.ReplaceAll(apps[1], ")", ""), 10, 0)
							if err != nil {
								Logger.Error().Timestamp().Msg(err.Error())
								return nil, nil, err
								// log.Fatal().Msg(err.Error())
							}
							// playerSeason.Stats.Subs = int(subs)
							playerSeason.Subs = int(subs)
						}
					}

				case 11, 15, 18:
					// Gets the stats that are percentages: passing, win, and save
					if data != "-" {
						data := strings.ReplaceAll(data, "%", "")
						perc, err := strconv.ParseInt(data, 10, 0)
						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
						}
						if col == 11 {
							reflect.ValueOf(&playerSeason).Elem().FieldByName("PasP").SetInt(perc)
						} else {
							reflect.ValueOf(&playerSeason).Elem().FieldByName(colNames[col-1]).SetInt(perc)
						}
					}
				case 14:
					// Gets the only float data: average rating
					if data != "-" {
						rating, err := strconv.ParseFloat(data, 32)
						if err != nil {
							Logger.Error().Timestamp().Msg(err.Error())
							return nil, nil, err
						}
						playerSeason.AvgRat = float32(rating)
					}
				}
			}

			// At end of column add to slices and reset
			if col == 65 {
				col = 0
				// playerSeason.Season = season
				// playerSeason.Name = playerInfo.Name
				playerSeason.UID = playerInfo.UID
				allInfo = append(allInfo, playerInfo)
				allSeasons = append(allSeasons, playerSeason)
				playerInfo = PlayerInfo{}
				playerSeason = PlayerSeason{}
			}

			isTd = false
			isTh = false
		}
	}
}
