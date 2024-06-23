package backend

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

type Settings struct {
	path                  string `json:"-"`
	ReadwiseToken         string `json:"readwise_token"`
	UploadCovers          bool   `json:"upload_covers"`
	UploadStoreHighlights bool   `json:"upload_store_highlights"`
}

func LoadSettings(portable bool) (*Settings, error) {
	settingsPath, err := LocateConfigFile(configFilename, portable)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create settings directory. Do you have proper permissions?")
	}
	s := &Settings{
		path:                  settingsPath,
		UploadStoreHighlights: true, // default on as users with only store purchased books are blocked from usage otherwise but give ample warning during setup
		UploadCovers:          false,
	}
	b, err := os.ReadFile(settingsPath)
	if err != nil {
		if os.IsNotExist(err) {
			// We should always have settings but if they have been deleted, just use the defaults
			return s, nil
		}
		return nil, errors.Wrap(err, "Failed to read settings file. Is it corrupted?")
	}
	err = json.Unmarshal(b, s)
	if err != nil {
		// v1.6.0 and prior may have caused settings files to have an extra `}` so we check for common issues
		// We're not gonna go overboard here though, just basic text checking
		plainSettings := strings.TrimSpace(string(b))
		if strings.HasPrefix(plainSettings, "{{") {
			plainSettings = strings.Replace(plainSettings, "{{", "{", 1)
		}
		if strings.HasSuffix(plainSettings, "}}") {
			plainSettings = strings.Replace(plainSettings, "}}", "}", 1)
		}
		err := json.Unmarshal([]byte(plainSettings), s)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to parse settings file. Is it valid?")
		}
		// We managed to fix the settings file so we'll persist it to disc
		err = s.Save()
		if err != nil {
			return nil, errors.Wrap(err, "Failed to persist fixed up settings file!")
		}
		return s, nil
	}
	return s, nil
}

func (s *Settings) Save() error {
	b, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return errors.Wrap(err, "Failed to save settings to disc")
	}
	err = os.MkdirAll(filepath.Dir(s.path), 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to create settings directory. Do you have proper permissions?")
	}
	err = os.WriteFile(s.path, b, 0777)
	if err != nil {
		return errors.Wrap(err, "Failed to create settings file. Do you have proper permissions?")
	}
	return nil
}

func (s *Settings) SaveToken(token string) error {
	s.ReadwiseToken = token
	return s.Save()
}

func (s *Settings) SaveCoverUploading(uploadCovers bool) error {
	s.UploadCovers = uploadCovers
	return s.Save()
}

func (s *Settings) SaveStoreHighlights(uploadStoreHighlights bool) error {
	s.UploadStoreHighlights = uploadStoreHighlights
	return s.Save()
}
