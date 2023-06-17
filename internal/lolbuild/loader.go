package lolbuild

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

const (
	buildCollectionsHost       = "https://pub-a6897392811e428684cf50a774ddc3fc.r2.dev"
	dataDragonVersionsUrl      = "https://ddragon.leagueoflegends.com/api/versions.json"
	dataDragonRunesReforgedUrl = "http://ddragon.leagueoflegends.com/cdn/%s/data/en_US/runesReforged.json"
	dataDragonAssetsUrl        = "https://ddragon.leagueoflegends.com/cdn/img/"
)

type Loader struct {
	runeTrees RuneTrees
	version   string
}

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
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	IconUrl string `json:"iconUrl"`
}
type Items struct {
	Starting []Item `json:"starting"`
	Core     []Item `json:"core"`
	Mythic   Item   `json:"mythic"`
	Fourth   []Item `json:"fourth"`
	Fifth    []Item `json:"fifth"`
	Sixth    []Item `json:"sixth"`
}

type Build struct {
	Name           string          `json:"name"`
	Winrate        float64         `json:"winrate"`
	Matches        string          `json:"matches"`
	Primary        Rune            `json:"primary"`
	Secondary      Rune            `json:"secondary"`
	SelectedPerks  []Rune          `json:"selectedPerks"`
	SummonerSpells []SummonerSpell `json:"summonerSpells"`
	Items          Items           `json:"items"`
}

type BuildCollection struct {
	Builds []Build `json:"builds"`
	Source string  `json:"source"`
}

func (l *Loader) getRuneTrees() RuneTrees {
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

func (l *Loader) LoadRuneTree(name string) RuneTree {
	runeTrees := l.getRuneTrees()

	for _, runeTree := range runeTrees {
		if runeTree.Name == name {
			return runeTree
		}
	}

	return RuneTree{}
}

func (l *Loader) LoadBuilds(championName string, sources []string, role string) []BuildCollection {
	var wg sync.WaitGroup
	builds := make([]BuildCollection, 0)

	results := make(chan BuildCollection, len(sources))

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			results <- *l.loadBuild(championName, source, role)
		}(source)
	}

	go func() {
		for result := range results {
			builds = append(builds, result)
			wg.Done()
		}
	}()

	wg.Wait()

	return builds
}

func (l *Loader) loadBuild(championName string, source string, role string) *BuildCollection {
	println(fmt.Sprintf("%s/%s/%s/%s.json", buildCollectionsHost, source, role, l.clearChampionName(championName)))
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/%s.json", buildCollectionsHost, source, role, l.clearChampionName(championName)))
	if err != nil {
		println(err.Error())
		return nil
	}

	var buildCollection BuildCollection
	err = json.NewDecoder(resp.Body).Decode(&buildCollection)

	if err != nil {
		println(err.Error())
		return nil
	}

	buildCollection.Source = l.getSourceName(source)

	return &buildCollection
}

func (l *Loader) clearChampionName(championName string) string {
	clearedChampionName := strings.Replace(championName, "'", "", -1)
	clearedChampionName = strings.Replace(championName, " ", "", -1)
	clearedChampionName = strings.ToLower(clearedChampionName)

	return clearedChampionName
}

func (l *Loader) getSourceName(source string) string {
	names := map[string]string{
		"ugg": "U.GG",
	}

	name, ok := names[source]
	if !ok {
		return "Unknown"
	}

	return name
}

func (l *Loader) getLatestVersion() string {
	if l.version != "" {
		return l.version
	}

	resp, err := http.Get(dataDragonVersionsUrl)
	if err != nil {
		log.Fatalln("Error fetching versions")
	}

	var versions []string
	json.NewDecoder(resp.Body).Decode(&versions)

	l.version = versions[0]

	return l.version
}

func transformRawRuneTrees(rawRuneTrees []RawRuneTree) RuneTrees {
	runeTrees := make(RuneTrees, 0)

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
