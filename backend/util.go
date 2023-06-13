package backend

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

var (
	stories *[]*Story
)

type Story struct {
	SaveID int    `json:"saveID"`
	Story  string `json:"story"`
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
		Logger.Error().Timestamp().Msg(err.Error())
		return
	}
	if _, err := os.Stat(storyFilePath); err != nil {
		if os.IsNotExist(err) {
			_, err := os.Create(storyFilePath)
			if err != nil {
				Logger.Error().Timestamp().Msg(err.Error())
				return
			}
		} else {
			Logger.Error().Timestamp().Msg(err.Error())
			return
		}
	}
	file, err := os.ReadFile(storyFilePath)
	// Logger.Info().Timestamp().Msg(string(file))
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return
	}
	if len(file) > 0 {
		err = json.Unmarshal(file, &stories)
		if err != nil {
			Logger.Error().Timestamp().Msg(err.Error())
			return
		}
	}

}

func WriteStoriesToFile(stories *[]*Story) error {
	// fmt.Println(stories)
	storyFilePath, err := xdg.DataFile("Save Tracker/stories.json")
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	jsonedStories, err := json.Marshal(stories)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	err = os.WriteFile(storyFilePath, jsonedStories, 0666)
	if err != nil {
		Logger.Error().Timestamp().Msg(err.Error())
		return err
	}
	return nil
}

func GetSaveStory(saveID int) Story {
	// fmt.Println(stories)
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
