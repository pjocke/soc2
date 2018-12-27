package scraper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/cheggaaa/pb"
	"github.com/kracekumar/go-mwapi"
)

type titles struct {
	Query struct {
		Pages map[int64]struct {
			Categoryinfo struct {
				Pages int64
			}
		}
	}
}

type categories struct {
	Continue struct {
		Cmcontinue string
	}
	Query struct {
		Categorymembers []struct {
			Pageid int64
			Title  string
		}
	}
}

type pages struct {
	Query struct {
		Pages map[int64]struct {
			PageID    int64
			Revisions []struct {
				Text string `json:"*"`
			}
		}
	}
}

//type words []word

type word struct {
	Word   string
	Genus  string
	Wahlin string
	Rhyme  string
}

func getcategorymembercount(api *mwapi.MWApi, category string) int64 {
	params := url.Values{
		"action": {"query"},
		"prop":   {"categoryinfo"},
		"titles": {category},
	}

	res := api.Get(params)

	if res.StatusCode != 200 {
		log.Fatal()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal()
	}

	t := titles{}
	err = json.Unmarshal([]byte(body), &t)
	if err != nil {
		log.Fatal(err)
	}

	// This is pretty ugly but needs to be done due to how the returned JSON is formatted.
	for _, page := range t.Query.Pages {
		return page.Categoryinfo.Pages
	}

	return 0
}

func getcategories(api *mwapi.MWApi, category string, words *[]word, pb *pb.ProgressBar, continueToken string) {
	params := url.Values{
		"action":     {"query"},
		"list":       {"categorymembers"},
		"cmtitle":    {category},
		"format":     {"json"},
		"cmlimit":    {"500"},
		"cmcontinue": {continueToken},
	}

	res := api.Get(params)

	if res.StatusCode != 200 {
		log.Fatal()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal()
	}

	c := categories{}
	err = json.Unmarshal([]byte(body), &c)
	if err != nil {
		log.Fatal(err)
	}

	for _, member := range c.Query.Categorymembers {
		pb.Increment()
		if !strings.HasPrefix(member.Title, "Kategori:") {
			genus := getpage(api, member.Pageid)
			end, _ := getending(member.Title)
			*words = append(*words, word{
				Word:  member.Title,
				Genus: genus,
				Rhyme: end,
			})
		}
	}

	if c.Continue.Cmcontinue != "" {
		getcategories(api, category, words, pb, c.Continue.Cmcontinue)
	}
}

func getpage(api *mwapi.MWApi, pageID int64) string {
	params := url.Values{
		"action":  {"query"},
		"prop":    {"revisions"},
		"rvprop":  {"content"},
		"pageids": {strconv.FormatInt(pageID, 10)},
	}

	res := api.Get(params)

	if res.StatusCode != 200 {
		log.Fatal()
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal()
	}

	p := pages{}

	err = json.Unmarshal([]byte(body), &p)
	if err != nil {
		log.Fatal(err)
	}

	if strings.Contains(p.Query.Pages[pageID].Revisions[0].Text, "sv-subst-n") {
		return ("n")
	} else if strings.Contains(p.Query.Pages[pageID].Revisions[0].Text, "sv-subst-t") {
		return ("t")
	} else {
		return ("")
	}
}

func getending(s string) (string, error) {
	word := []rune(s)

	l := len(word)

	levels := 1
	if strings.ContainsAny(strings.ToLower(string(word[l-1])), "aeiouyåäö") {
		levels = 2
	}

	level := 0
	for i := l; i > 0; i-- {
		if strings.ContainsAny(strings.ToLower(string(word[i-1])), "aeiouyåäö") {
			level++
		}
		if level == levels {
			return string(word[i-1 : l]), nil
		}
	}

	return "", errors.New("fel fel fel")
}

// Scrape my bitch up
func Scrape() {
	category := "Kategori:Svenska/Substantiv"

	api := mwapi.NewMWApi(url.URL{
		Scheme: "https",
		Host:   "sv.wiktionary.org",
		Path:   "/w/api.php",
	})

	words := []word{}

	count := getcategorymembercount(api, category)
	if count > 0 {
		pb := pb.StartNew(int(count))
		getcategories(api, category, &words, pb, "")
		pb.Finish()
	}

	b, err := json.MarshalIndent(words, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("./soc2.json", b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
