package handler

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/viper"
	"github.com/titolins/jira_scrapper/internal/cache"
)

type dataFetcher func() (interface{}, error)

type Handler struct {
	c *cache.Cache
	f dataFetcher
}

func New(c *cache.Cache, f dataFetcher) *Handler {
	return &Handler{c, f}
}

func (h *Handler) Handle(v interface{}) (err error) {
	if fi, exists := h.c.Exists(); exists {
		if err = h.handleExists(fi, v); err != nil {
			return err
		}
	} else {
		if err = h.fetch(v); err != nil {
			return err
		}
	}

	if viper.GetBool("print") {
		log.Println("Printing boards")
		fmt.Println(v)
	}

	return
}

func (h *Handler) handleExists(fi os.FileInfo, v interface{}) error {
	var in string
	if viper.GetBool("skip-cached") {
		in = "no"
	} else {
		in = h.prompt(fi)
	}

	switch in {
	case "no":
		log.Println("Loading sprints from cache")
		if err := h.loadCache(v); err != nil {
			return fmt.Errorf("failed to load cache: %v\n", err)
		}
	case "yes":
		log.Println("Fetching sprints from jira")
		if err := h.fetch(v); err != nil {
			return fmt.Errorf("failed to fetch data: %v\n", err)
		}
	default:
		return errors.New("invalid choice")
	}

	return nil
}

func (h *Handler) loadCache(v interface{}) error {
	return h.c.Load(v)
}

func (h *Handler) fetch(v interface{}) error {
	d, err := h.f()
	if err != nil {
		return fmt.Errorf("failed to fetch data: %v\n", err)
	}

	if err = h.c.Save(d); err != nil {
		return fmt.Errorf("failed to save cache: %v\n", err)
	}
	return h.loadCache(v)
}

func (h *Handler) info(fi os.FileInfo) {
	fmt.Printf(
		"Cache file existent!\nFile: %s\nModified at: %s\n",
		h.c.Path(),
		fi.ModTime().Format(time.RFC3339Nano),
	)
}

func (h *Handler) prompt(fi os.FileInfo) string {
	h.info(fi)
	return prompt.Input(
		"Would you like to update it? (yes / no)",
		completer,
		prompt.OptionPrefixBackgroundColor(prompt.Red),
		prompt.OptionPrefixTextColor(prompt.White),
	)
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "yes", Description: "Fetches data from JIRA and updates the cache"},
		{Text: "no", Description: "Prints current existing cache"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
