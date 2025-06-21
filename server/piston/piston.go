package piston

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/sriramr98/dsa_server/utils"
)

var PISTON_BASE_API_URL = utils.GetEnv("PISTON_API_URL", "http://localhost:2000")

type LanguageConfig struct {
	Language  string `json:"language"`
	Version   string `json:"language_version"`
	Installed bool   `json:"installed"`
}

// SetupLanguages initializes the languages used in the application.
// It installs the language runtimes with piston and makes sure it's ready for use.
func SetupLanguages() error {
	supportedLanguages, err := getSupportedLanguages()
	if err != nil {
		utils.LogError(err)
		return err
	}

	for _, lang := range SUPPORTED_LANGUAGES {
		langConfig, found := utils.FindInSlice(supportedLanguages, func(l LanguageConfig) bool {
			return l.Language == lang.PistonAlias && l.Version == lang.Version
		})

		if !found {
			err := fmt.Errorf("language %s with version %s not found in piston supported languages", lang.PistonAlias, lang.Version)
			utils.LogError(err)
			return err
		}

		if !langConfig.Installed {
			if err := installLanguage(langConfig); err != nil {
				utils.LogError(err)
				return fmt.Errorf("failed to install language %s: %w", langConfig.Language, err)
			}
		}
	}

	return nil
}

func installLanguage(langConfig LanguageConfig) error {
	log.Printf("Installing language %s with version %s", langConfig.Language, langConfig.Version)
	apiUrl := fmt.Sprintf("%s/api/v2/packages", PISTON_BASE_API_URL)
	payload := map[string]string{
		"language": langConfig.Language,
		"version":  langConfig.Version,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		utils.LogError(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, apiUrl, bytes.NewBuffer(jsonPayload))
	if err != nil {
		utils.LogError(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogError(err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to install language %s, status code: %d", langConfig.Language, resp.StatusCode)
		utils.LogError(err)
		return err
	}
	return nil
}

func getSupportedLanguages() ([]LanguageConfig, error) {
	apiUrl := fmt.Sprintf("%s/api/v2/packages", PISTON_BASE_API_URL)
	log.Println("Fetching supported languages from:", apiUrl)
	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		utils.LogError(err)
		return []LanguageConfig{}, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		utils.LogError(err)
		return []LanguageConfig{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("failed to fetch supported languages, status code: %d", resp.StatusCode)
		utils.LogError(err)
		return []LanguageConfig{}, err
	}

	var languages []LanguageConfig
	if err := json.NewDecoder(resp.Body).Decode(&languages); err != nil {
		utils.LogError(err)
		return []LanguageConfig{}, err
	}

	return languages, nil
}
