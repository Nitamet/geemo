package lolbuild

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
)

const host = "https://pub-a6897392811e428684cf50a774ddc3fc.r2.dev"

type Loader struct{}

type Rune struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Slug    string `json:"slug"`
	IconUrl string `json:"iconUrl"`
	Path    *Rune  `json:"path"`
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
	Name          string `json:"name"`
	Winrate       string `json:"winrate"`
	Matches       string `json:"matches"`
	Primary       Rune   `json:"primary"`
	Secondary     Rune   `json:"secondary"`
	SelectedPerks []Rune `json:"selectedPerks"`
	Items         Items  `json:"items"`
}

type BuildCollection struct {
	Builds []Build `json:"runes"`
	Source string  `json:"source"`
}

func (l *Loader) LoadBuilds(championName string, sources []string) []BuildCollection {
	var wg sync.WaitGroup
	builds := make([]BuildCollection, 0)

	results := make(chan BuildCollection, len(sources))

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			results <- *l.loadBuild(championName, source)
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

func (l *Loader) loadBuild(championName string, source string) *BuildCollection {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s.json", host, source, l.clearChampionName(championName)))
	if err != nil {
		println(err.Error())
		return nil
	}

	var build BuildCollection
	err = json.NewDecoder(resp.Body).Decode(&build)

	if err != nil {
		println(err.Error())
		return nil
	}

	build.Source = l.getSourceName(source)

	return &build
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
