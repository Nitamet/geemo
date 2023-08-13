// Package lolbuild provides functions for loading builds from different sources
package lolbuild

import (
	"encoding/json"
	"fmt"
	"github.com/Nitamet/geemo/backend"
	"log"
	"net/http"
	"strings"
	"sync"
)

const buildDataVersion = 1

// URLs-related constants
const (
	buildCollectionsHost       = "https://geemo.app"
	dataDragonVersionsUrl      = "https://ddragon.leagueoflegends.com/api/versions.json"
	dataDragonRunesReforgedUrl = "http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/runesReforged.json"
	dataDragonAssetsUrl        = "https://ddragon.leagueoflegends.com/cdn/img/"
)

type Loader struct {
	runeTrees RuneTrees // All game rune trees
	version   string    // Latest version of the game
}

// RawRuneTree is a raw rune tree structure from data dragon
type RawRuneTree struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Icon  string `json:"icon"`
	Name  string `json:"name"`
	Slots []struct {
		Runes []struct {
			ID        int    `json:"id"`
			Key       string `json:"key"`
			Icon      string `json:"icon"`
			Name      string `json:"name"`
			ShortDesc string `json:"shortDesc"`
			LongDesc  string `json:"longDesc"`
		} `json:"runes"`
	} `json:"slots"`
}

// RuneTree is a rune tree structure ready to be used in the app
type RuneTree struct {
	Name      string `json:"name"`
	Keystones []Rune `json:"keystones"`
	Perks     []Rune `json:"perks"`
	IconUrl   string `json:"iconUrl"`
}
type RuneTrees = []RuneTree

type Rune struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	IconUrl string `json:"iconUrl"`
	Path    *Rune  `json:"path"`
}
type SummonerSpell struct {
	ID      int    `json:"id"`
	IconUrl string `json:"iconUrl"`
	Name    string `json:"name"`
}
type Item struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	IconUrl  string `json:"iconUrl"`
	IsMythic bool   `json:"isMythic"`
}
type ItemGroup struct {
	Items []Item `json:"items"`
	Name  string `json:"name"`
}

// Build is a build structure ready to be used in the app
type Build struct {
	Name           string          `json:"name"`
	Winrate        float64         `json:"winrate"`
	Matches        string          `json:"matches"`
	Primary        Rune            `json:"primary"`
	Secondary      Rune            `json:"secondary"`
	SelectedPerks  []Rune          `json:"selectedPerks"`
	SummonerSpells []SummonerSpell `json:"summonerSpells"`
	ItemGroups     []ItemGroup     `json:"itemGroups"`
	Mythic         Item            `json:"mythic"`
}

// BuildCollection is a collection of builds from a single source
type BuildCollection struct {
	Builds []Build `json:"builds"`
	Source string  `json:"source"`
}

// loadRuneTrees loads all rune trees from data dragon
func (l *Loader) loadRuneTrees() RuneTrees {
	if l.runeTrees != nil {
		return l.runeTrees
	}

	resp, err := http.Get(fmt.Sprintf(dataDragonRunesReforgedUrl, l.getLatestVersion()))
	if err != nil {
		log.Fatal(err)
	}

	var rawRuneTrees []RawRuneTree

	err = json.NewDecoder(resp.Body).Decode(&rawRuneTrees)
	if err != nil {
		log.Fatal(err)
	}

	l.runeTrees = transformRawRuneTrees(rawRuneTrees)

	return l.runeTrees
}

// GetRuneTree returns a rune tree by its name
func (l *Loader) GetRuneTree(name string) RuneTree {
	defer backend.LogPanic()

	runeTrees := l.loadRuneTrees()

	for _, runeTree := range runeTrees {
		if runeTree.Name == name {
			return runeTree
		}
	}

	log.Printf("Unknown rune tree %s", name)

	return RuneTree{}
}

// LoadBuilds loads builds for a given champion and role from specified sources
func (l *Loader) LoadBuilds(championName string, sources []string, role string) []BuildCollection {
	defer backend.LogPanic()

	var wg sync.WaitGroup
	builds := make([]BuildCollection, 0)

	results := make(chan BuildCollection, len(sources))

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			defer backend.LogPanic()

			// Try to load build from the latest version of the game and 2 previous versions if it fails
			for i := buildDataVersion; i >= buildDataVersion-2 && i > 0; i-- {
				build := l.loadBuild(championName, source, role, i)
				if build != nil {
					results <- *build
					break
				}
			}
		}(source)
	}

	go func() {
		defer backend.LogPanic()

		for result := range results {
			builds = append(builds, result)
			wg.Done()
		}
	}()

	wg.Wait()

	return builds
}

// loadBuild loads builds for a given champion and role from a specified source and version
func (l *Loader) loadBuild(championName, source, role string, version int) *BuildCollection {
	buildUrl := fmt.Sprintf("%s/%s/%d/%s/%s.json", buildCollectionsHost, source, version, role, l.clearChampionName(championName))

	log.Printf("Loading build from %s", buildUrl)

	resp, err := http.Get(buildUrl)
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	var buildCollection BuildCollection
	err = json.NewDecoder(resp.Body).Decode(&buildCollection)

	if err != nil {
		log.Println("Error decoding build: ", err)
		return nil
	}

	buildCollection.Source = l.getSourceName(source)

	log.Println("Build loaded")

	return &buildCollection
}

// clearChampionName clears champion name from spaces, apostrophes, dots and converts it to lowercase
func (l *Loader) clearChampionName(championName string) string {
	clearedChampionName := strings.Replace(championName, "'", "", -1)
	clearedChampionName = strings.Replace(clearedChampionName, ".", "", -1)
	clearedChampionName = strings.Replace(clearedChampionName, " ", "", -1)
	clearedChampionName = strings.ToLower(clearedChampionName)

	return clearedChampionName
}

// getSourceName returns a human-readable name for a given source
func (l *Loader) getSourceName(source string) string {
	names := map[string]string{
		"ugg":        "U.GG",
		"mobalytics": "Mobalytics",
	}

	name, ok := names[source]
	if !ok {
		return "Unknown"
	}

	return name
}

// getLatestVersion returns the latest version of the game
func (l *Loader) getLatestVersion() string {
	if l.version != "" {
		return l.version
	}

	resp, err := http.Get(dataDragonVersionsUrl)
	if err != nil {
		log.Panic("Error fetching versions")
	}

	var versions []string
	err = json.NewDecoder(resp.Body).Decode(&versions)
	if err != nil {
		log.Panic(err)
	}

	l.version = versions[0]

	return l.version
}

// transformRawRuneTrees transforms raw rune trees from data dragon into rune trees that are ready to be used in the app
func transformRawRuneTrees(rawRuneTrees []RawRuneTree) RuneTrees {
	runeTrees := make(RuneTrees, 0, len(rawRuneTrees))

	for _, rawRuneTree := range rawRuneTrees {
		runeTree := RuneTree{
			Name:    rawRuneTree.Name,
			IconUrl: dataDragonAssetsUrl + rawRuneTree.Icon,
		}

		keystones := make([]Rune, 0)
		for _, rawRune := range rawRuneTree.Slots[0].Runes {
			keystones = append(keystones, Rune{
				ID:      rawRune.ID,
				Name:    rawRune.Name,
				Slug:    rawRune.Key,
				IconUrl: dataDragonAssetsUrl + rawRune.Icon,
			})
		}

		perks := make([]Rune, 0)
		for i := 1; i < 4; i++ {
			for _, rawRune := range rawRuneTree.Slots[i].Runes {
				perks = append(perks, Rune{
					ID:      rawRune.ID,
					Name:    rawRune.Name,
					Slug:    rawRune.Key,
					IconUrl: dataDragonAssetsUrl + rawRune.Icon,
				})
			}
		}

		runeTree.Keystones = keystones
		runeTree.Perks = perks

		runeTrees = append(runeTrees, runeTree)
	}

	return runeTrees
}
