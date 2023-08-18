package backend

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/adrg/xdg"
)

var (
	stories      *[]*Story
	version      *Version
	curVersion   string = "v1.0.0"
	curDBVersion string = "v1.0.0"
	releaseLink  string = "https://api.github.com/repos/lucaono13/savetracker/releases/latest"
)

type Story struct {
	SaveID int    `json:"saveID"`
	Story  string `json:"story"`
}

type Version struct {
	Version   string `json:"version"`
	DBVersion string `json:"dbVersion"`
}

type GithubReleases struct {
	URL string `json:"html_url"`
	Tag string `json:"tag_name"`
}

func MakePlayersMap(players []PlayerInfo) map[int]int {
	playersMap := make(map[int]int)
	for _, player := range players {
		playersMap[player.UID] = int(player.PlayerID)
	}
	return playersMap
}

func MakeTrophiesMap(trophies []Trophy) map[string]int {
	trophiesMap := make(map[string]int)
	for _, trophy := range trophies {
		trophiesMap[trophy.CompetitionName] = int(trophy.TrophyID.Int64)
	}
	return trophiesMap
}

func GetStoriesFromFile() {
	storyFilePath, err := xdg.DataFile("Save Tracker/stories.json")
	if err != nil {
		Logger.Error().Msg(err.Error())
		return
	}
	if _, err := os.Stat(storyFilePath); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(storyFilePath)
			if err != nil {
				Logger.Error().Msg(err.Error())
				return
			}
		} else {
			Logger.Error().Msg(err.Error())
			return
		}
	}
	file, err := os.ReadFile(storyFilePath)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &stories)
		if err != nil {
			Logger.Error().Msg(err.Error())
			return
		}
	}

}

func WriteStoriesToFile(stories *[]*Story) error {
	storyFilePath, err := xdg.DataFile("Save Tracker/stories.json")
	if err != nil {
		Logger.Error().Msg(err.Error())
		return err
	}
	jsonedStories, err := json.Marshal(stories)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return err
	}
	err = os.WriteFile(storyFilePath, jsonedStories, 0666)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return err
	}
	return nil
}

func GetSaveStory(saveID int) Story {
	var saveStory Story = Story{SaveID: saveID, Story: ""}
	for _, story := range *stories {
		if story.SaveID == saveID {
			saveStory.Story = story.Story
		}
	}
	return saveStory
}

func UpdateSaveStory(updatedStory Story) error {
	fmt.Println(updatedStory)
	var updated bool = false
	for _, story := range *stories {
		if story.SaveID == updatedStory.SaveID {
			*story = updatedStory
			updated = true

		}
	}
	if !updated {
		*stories = append(*stories, &updatedStory)
	}
	fmt.Println(stories, updated)
	return WriteStoriesToFile(stories)
}

func GetConfig() {
	configFilePath, err := xdg.ConfigFile("Save Tracker/versions.json")
	if err != nil {
		Logger.Error().Msg(err.Error())
		return
	}
	if _, err := os.Stat(configFilePath); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(configFilePath)
			if err != nil {
				Logger.Error().Msg(err.Error())
				return
			}
			content, err := json.Marshal(
				Version{
					Version:   curVersion,
					DBVersion: curDBVersion,
				},
			)
			if err != nil {
				Logger.Error().Msg(err.Error())
				return
			}
			err = os.WriteFile(configFilePath, content, 0666)
			if err != nil {
				Logger.Error().Msg(err.Error())
				return
			}
		} else {
			Logger.Error().Msg(err.Error())
			return
		}
	}
	file, err := os.ReadFile(configFilePath)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return
	}
	err = json.Unmarshal(file, &version)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return
	}
}

func GetVersion() string {
	return version.Version
}

func GetDBVersion() string {
	return version.DBVersion
}

// Check Github version
func CheckVersion() (bool, string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", releaseLink, nil)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return false, "", err
	}
	req.Header.Add("Accept", "application/vnd.github+json")
	req.Header.Add("X-Github-Api-Version", "2022-11-28")

	resp, err := client.Do(req)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return false, "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return false, "", err
	}

	githubVersion := GithubReleases{}
	err = json.Unmarshal(body, &githubVersion)
	if err != nil {
		Logger.Error().Msg(err.Error())
		return false, "", nil
	}
	if version.Version == githubVersion.Tag {
		return false, "", nil
	} else {
		return true, githubVersion.URL, nil
	}
}
