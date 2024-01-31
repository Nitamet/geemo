// Package lolbuild provides functions for loading builds from different sources
package lolbuild

import (
	"encoding/json"
	"fmt"
	"github.com/Nitamet/geemo/backend"
	"golang.org/x/text/cases"
	golangLanguage "golang.org/x/text/language"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

const buildDataVersion = 3

// previousVersionsToTry is a number of previous build versions to try if the current one is unavailable
const previousVersionsToTry = 2

// URLs-related constants
const (
	buildCollectionsHost       = "https://geemo.app"
	dataDragonVersionsUrl      = "https://ddragon.leagueoflegends.com/api/versions.json"
	dataDragonRunesReforgedUrl = "https://ddragon.leagueoflegends.com/cdn/%s/data/%s/runesReforged.json"
	dataDragonSpellsUrl        = "https://ddragon.leagueoflegends.com/cdn/%s/data/%s/summoner.json"
	dataDragonAssetsUrl        = "https://ddragon.leagueoflegends.com/cdn/img/"
	dataDragonSpellIconUrl     = "https://ddragon.leagueoflegends.com/cdn/%s/img/spell/%s.png"
	dataDragonPassiveIconUrl   = "https://ddragon.leagueoflegends.com/cdn/%s/img/passive/%s.png"
	dataDragonItemUrl          = "https://ddragon.leagueoflegends.com/cdn/%s/data/%s/item.json"
	dataDragonItemIcon         = "https://ddragon.leagueoflegends.com/cdn/%s/img/item/%s.png"
	dataDragonChampionsUrl     = "https://ddragon.leagueoflegends.com/cdn/%s/data/%s/champion.json"
	dataDragonChampionUrl      = "https://ddragon.leagueoflegends.com/cdn/%s/data/%s/champion/%s.json"
)

type AssetData struct {
	Name    string
	Slug    string
	IconUrl string
}

type Loader struct {
	runeTrees         RuneTrees // All game rune trees
	runeData          map[int]AssetData
	summonerSpellData map[int]AssetData
	championData      map[int]AssetData
	itemData          map[int]AssetData
	language          string
	version           string // Latest version of the game
	allSources        map[string]string
	mutex             sync.Mutex
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
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	IconUrl    string `json:"iconUrl"`
	IsCoreItem bool   `json:"IsCoreItem"`
}
type ItemGroup struct {
	Items []Item `json:"items"`
	Name  string `json:"name"`
}

// Build is a build structure ready to be used in the app
type Build struct {
	Name           string          `json:"name"`
	Winrate        float64         `json:"winrate"`
	Matches        int             `json:"matches"`
	Primary        Rune            `json:"primary"`
	Secondary      Rune            `json:"secondary"`
	SelectedPerks  []Rune          `json:"selectedPerks"`
	SummonerSpells []SummonerSpell `json:"summonerSpells"`
	ItemGroups     []ItemGroup     `json:"itemGroups"`
	CoreItem       Item            `json:"coreItem"`
	SkillMaxOrder  []int           `json:"skillMaxOrder"`
	SkillOrder     []int           `json:"skillOrder"`
}

// BuildCollection is a collection of builds from a single source
type BuildCollection struct {
	Builds []Build `json:"builds"`
	Source string  `json:"source"`
}

type SpellImage struct {
	Full string `json:"full"`
}
type Spell struct {
	Name  string     `json:"name"`
	Image SpellImage `json:"image"`
}
type Passive struct {
	Name  string     `json:"name"`
	Image SpellImage `json:"image"`
}
type ChampionData struct {
	Spells  []Spell `json:"spells"`
	Passive Passive `json:"passive"`
}
type championDataResponse struct {
	Data map[string]ChampionData `json:"data"`
}

// loadRuneTrees loads all rune trees from data dragon
func (l *Loader) loadRuneTrees() RuneTrees {
	if l.runeTrees != nil && len(l.runeTrees) > 0 {
		return l.runeTrees
	}

	l.runeData = make(map[int]AssetData)

	resp, err := http.Get(fmt.Sprintf(dataDragonRunesReforgedUrl, l.getLatestVersion(), l.language))
	if err != nil {
		log.Fatal(err)
	}

	var rawRuneTrees []RawRuneTree

	err = json.NewDecoder(resp.Body).Decode(&rawRuneTrees)
	if err != nil {
		log.Fatal(err)
	}

	l.runeTrees = transformRawRuneTrees(rawRuneTrees)

	l.mutex.Lock()
	for _, style := range rawRuneTrees {
		l.runeData[style.ID] = AssetData{
			Name:    style.Name,
			IconUrl: style.Icon,
		}
		for _, slot := range style.Slots {
			for _, lolRune := range slot.Runes {
				l.runeData[lolRune.ID] = AssetData{
					Name:    lolRune.Name,
					IconUrl: fmt.Sprintf(dataDragonAssetsUrl + lolRune.Icon),
				}
			}
		}
	}
	l.mutex.Unlock()

	return l.runeTrees
}

type item struct {
	Name string          `json:"name"`
	Maps map[string]bool `json:"maps"`
}
type items struct {
	Data map[string]item `json:"data"`
}

func (l *Loader) loadItems() {
	if l.itemData != nil && len(l.itemData) > 0 {
		return
	}

	resp, err := http.Get(fmt.Sprintf(dataDragonItemUrl, l.getLatestVersion(), l.language))
	if err != nil {
		log.Fatalln("Error fetching items data")
	}

	var lolItems items
	err = json.NewDecoder(resp.Body).Decode(&lolItems)
	if err != nil {
		log.Fatalln("Error decoding items data")
		return
	}

	l.mutex.Lock()
	for id, lolItem := range lolItems.Data {
		itemId, _ := strconv.Atoi(id)

		l.itemData[itemId] = AssetData{
			IconUrl: fmt.Sprintf(dataDragonItemIcon, l.getLatestVersion(), id),
			Name:    lolItem.Name,
		}
	}
	l.mutex.Unlock()
}

func (l *Loader) loadSummonerSpells() {
	if l.summonerSpellData != nil && len(l.summonerSpellData) > 0 {
		return
	}

	resp, err := http.Get(fmt.Sprintf(dataDragonSpellsUrl, l.getLatestVersion(), l.language))
	if err != nil {
		log.Fatalln("Error fetching spells data")
	}

	var spells struct {
		Data map[string]struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"data"`
	}
	err = json.NewDecoder(resp.Body).Decode(&spells)
	if err != nil {
		log.Fatalln("Error decoding spells data")
		return
	}

	l.mutex.Lock()
	for _, spell := range spells.Data {
		key, _ := strconv.Atoi(spell.Key)

		l.summonerSpellData[key] = AssetData{
			IconUrl: fmt.Sprintf(dataDragonSpellIconUrl, l.getLatestVersion(), spell.ID),
			Name:    spell.Name,
		}
	}
	l.mutex.Unlock()
}

func (l *Loader) loadChampions() {
	if l.championData != nil && len(l.championData) > 0 {
		return
	}

	resp, err := http.Get(fmt.Sprintf(dataDragonChampionsUrl, l.getLatestVersion(), l.language))
	if err != nil {
		log.Fatalln("Error fetching champions data")
	}

	var champions struct {
		Data map[string]struct {
			Key  string `json:"key"`
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&champions)
	if err != nil {
		log.Fatalln("Error decoding champions data " + err.Error())
		return
	}

	l.mutex.Lock()
	for _, champion := range champions.Data {
		key, _ := strconv.Atoi(champion.Key)

		l.championData[key] = AssetData{
			IconUrl: "",
			Name:    champion.Name,
			Slug:    strings.ToLower(champion.ID),
		}
	}
	l.mutex.Unlock()
}

type ChampionName struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (l *Loader) GetChampionName(id int, language string) ChampionName {
	defer backend.LogPanic()

	if l.language != language {
		l.setLanguage(language)
		l.championData = nil
	}

	if l.championData == nil {
		l.championData = make(map[int]AssetData)
		l.loadChampions()
	}

	championData, ok := l.championData[id]
	if !ok {
		log.Printf("Unknown champion %d", id)

		return ChampionName{}
	}

	return ChampionName{
		Name: championData.Name,
		Slug: championData.Slug,
	}
}

// GetRuneTree returns a rune tree by its name
func (l *Loader) GetRuneTree(name string, language string) RuneTree {
	defer backend.LogPanic()

	if l.language != language {
		l.setLanguage(language)
		l.runeTrees = nil
	}

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
func (l *Loader) LoadBuilds(championName string, sources []string, role, language string) []BuildCollection {
	defer backend.LogPanic()

	if l.language != language {
		l.setLanguage(language)
		l.runeTrees = make(RuneTrees, 0)
		l.runeData = make(map[int]AssetData)
		l.summonerSpellData = make(map[int]AssetData)
		l.itemData = make(map[int]AssetData)
	}

	l.loadRuneTrees()
	l.loadSummonerSpells()
	l.loadItems()

	var wg sync.WaitGroup
	builds := make([]BuildCollection, 0)

	results := make(chan *BuildCollection, len(sources))

	wg.Add(len(sources))

	for _, source := range sources {
		go func(source string) {
			defer backend.LogPanic()

			// Try to load build from the latest version of the game and 2 previous versions if it fails
			for i := buildDataVersion; i >= buildDataVersion-previousVersionsToTry && i > 0; i-- {
				buildCollection := l.loadBuildCollection(championName, source, role, i)

				// If this is the last iteration, send any result to the channel
				isLastIteration := i == 1 || i == (buildDataVersion-previousVersionsToTry)
				if isLastIteration || buildCollection != nil {
					results <- buildCollection
					break
				}
			}
		}(source)
	}

	go func() {
		defer backend.LogPanic()

		for result := range results {
			if result != nil {
				builds = append(builds, *result)
			}

			wg.Done()
		}
	}()

	wg.Wait()

	return builds
}

// loadBuildCollection loads builds for a given champion and role from a specified source and version
func (l *Loader) loadBuildCollection(championName, source, role string, version int) *BuildCollection {
	buildUrl := fmt.Sprintf("%s/%s/%d/%s/%s.json", buildCollectionsHost, source, version, role, l.clearChampionName(championName))

	log.Printf("Loading build collection from %s", buildUrl)

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

	l.mutex.Lock()
	for buildIdx := range buildCollection.Builds {
		build := &buildCollection.Builds[buildIdx]

		build.Primary.Name = l.runeData[build.Primary.ID].Name
		build.Primary.IconUrl = l.runeData[build.Primary.ID].IconUrl

		build.Secondary.Name = l.runeData[build.Secondary.ID].Name
		build.Secondary.IconUrl = l.runeData[build.Secondary.ID].IconUrl

		for perkIdx := range build.SelectedPerks {
			perk := &build.SelectedPerks[perkIdx]

			runeData, ok := l.runeData[perk.ID]
			if !ok {
				log.Printf("Unknown rune %d", perk.ID)
				continue
			}

			perk.Name = runeData.Name
			perk.IconUrl = runeData.IconUrl
		}

		for spellIdx := range build.SummonerSpells {
			spell := &build.SummonerSpells[spellIdx]

			spellData, ok := l.summonerSpellData[spell.ID]
			if !ok {
				log.Printf("Unknown summoner spell %d", spell.ID)
				continue
			}

			spell.Name = spellData.Name
			spell.IconUrl = spellData.IconUrl
		}

		for itemGroupIdx := range build.ItemGroups {
			for itemIdx := range build.ItemGroups[itemGroupIdx].Items {
				lolItem := &build.ItemGroups[itemGroupIdx].Items[itemIdx]

				itemData, ok := l.itemData[lolItem.ID]
				if !ok {
					log.Printf("Unknown item %d", lolItem.ID)
					continue
				}

				lolItem.Name = itemData.Name
				lolItem.IconUrl = itemData.IconUrl

				if lolItem.IsCoreItem {
					build.CoreItem = *lolItem
				}
			}
		}
	}
	l.mutex.Unlock()

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
	if l.allSources == nil {
		l.loadSources()
	}

	name, ok := l.allSources[source]
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

func (l *Loader) setLanguage(language string) {
	// Replace - with _ in language name
	language = strings.Replace(language, "-", "_", -1)
	l.language = language
}

type BuildSource struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

func (l *Loader) loadSources() {
	resp, err := http.Get(buildCollectionsHost + "/sources.json")
	if err != nil {
		log.Panic("Error fetching sources")
	}

	var sources struct {
		Sources map[string]string `json:"sources"`
	}
	err = json.NewDecoder(resp.Body).Decode(&sources)
	if err != nil {
		log.Panic(err)
	}

	l.allSources = sources.Sources
}

func (l *Loader) GetSources() []BuildSource {
	if l.allSources == nil {
		l.loadSources()
	}

	var result []BuildSource
	for slug, name := range l.allSources {
		result = append(result, BuildSource{
			Slug: slug,
			Name: name,
		})
	}

	return result
}

// GetChampionData loads champion data from data dragon
func (l *Loader) GetChampionData(championId int, language string) *ChampionData {
	// TODO: Add cache
	if l.language != language {
		l.setLanguage(language)
	}

	championSlug := l.GetChampionName(championId, language).Slug

	championUrl := fmt.Sprintf(
		dataDragonChampionUrl,
		l.getLatestVersion(),
		l.language,
		cases.Title(golangLanguage.English, cases.NoLower).String(championSlug),
	)

	log.Printf("Loading championg data from %s", championUrl)

	resp, err := http.Get(championUrl)
	if err != nil {
		log.Panic(err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil
	}

	var championDataResponse championDataResponse
	err = json.NewDecoder(resp.Body).Decode(&championDataResponse)

	if err != nil {
		log.Println("Error decoding champion data: ", err)
		return nil
	}

	var championData *ChampionData
	for _, data := range championDataResponse.Data {
		championData = &data
		break
	}

	if championData == nil {
		log.Println("Champion data response is empty")
		return nil
	}

	for i := range championData.Spells {
		spell := &championData.Spells[i]
		spell.Image.Full = strings.Replace(spell.Image.Full, ".png", "", -1)
		spell.Image.Full = fmt.Sprintf(dataDragonSpellIconUrl, l.getLatestVersion(), spell.Image.Full)
	}

	championData.Passive.Image.Full = strings.Replace(championData.Passive.Image.Full, ".png", "", -1)
	championData.Passive.Image.Full = fmt.Sprintf(dataDragonPassiveIconUrl, l.getLatestVersion(), championData.Passive.Image.Full)

	return championData
}
