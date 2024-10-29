package service

import (
	"beef/internal/models"
	"beef/internal/repository"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type ServiceBeef struct {
	Repostiory repository.BeefRepository
}

func NewServiceBeef(repo repository.BeefRepository) *ServiceBeef {
	return &ServiceBeef{Repostiory: repo}
}

func (s *ServiceBeef) FindBeef() (*models.ModelBeef, error) {
	body, err := fetchBeefData()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch beef data: %v", err)
	}

	beefItems := processBeefData(body)
	beefCount := countValidBeefs(beefItems)

	res, err := s.Repostiory.FindBeef(beefCount)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func fetchBeefData() (string, error) {
	resp, err := http.Get("https://baconipsum.com/api/?type=beef-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func processBeefData(data string) []string {
	cleanedData := regexp.MustCompile(`[,.!\n]`).ReplaceAllString(data, " ")
	return strings.Fields(strings.ToLower(cleanedData))
}

func countValidBeefs(beefItems []string) map[string]uint32 {
	validBeefs := map[string]bool{
		"t-bone": true, "fatback": true, "pastrami": true, "pork": true,
		"meatloaf": true, "jowl": true, "enim": true, "bresaola": true,
	}

	beefCount := make(map[string]uint32)
	for _, beef := range beefItems {
		if validBeefs[beef] {
			beefCount[beef]++
		}
	}
	return beefCount
}
