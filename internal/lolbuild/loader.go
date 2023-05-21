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
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type RuneData struct {
	Name          string `json:"name"`
	WinRate       string `json:"winrate"`
	Primary       Rune   `json:"primary"`
	Secondary     Rune   `json:"secondary"`
	SelectedPerks []Rune `json:"selectedPerks"`
}
type Build struct {
	Runes []RuneData `json:"runes"`
}

func (l *Loader) LoadBuilds(championName string, sources []string) []Build {
	var wg sync.WaitGroup
	builds := make([]Build, len(sources))

	results := make(chan Build, len(sources))
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

func (l *Loader) loadBuild(championName string, source string) *Build {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s.json", host, source, l.clearChampionName(championName)))
	if err != nil {
		println(err.Error())
		return nil
	}

	var build Build
	err = json.NewDecoder(resp.Body).Decode(&build)
	println(fmt.Sprintf("%v", build))
	if err != nil {
		println(err.Error())
		return nil
	}

	return &build
}

func (l *Loader) clearChampionName(championName string) string {
	clearedChampionName := strings.Replace(championName, "'", "", -1)
	clearedChampionName = strings.Replace(championName, " ", "", -1)
	clearedChampionName = strings.ToLower(clearedChampionName)

	return clearedChampionName
}
