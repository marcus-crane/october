package backend

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/pkg/errors"
)

type Settings struct {
	path                  string `json:"-"`
	ReadwiseToken         string `json:"readwise_token"`
	UploadCovers          bool   `json:"upload_covers"`
	UploadStoreHighlights bool   `json:"upload_store_highlights"`
}

func LoadSettings() (*Settings, error) {
	settingsPath, err := xdg.ConfigFile(configFilename)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create settings directory. Do you have proper permissions?")
	}
	s := &Settings{
		path:         settingsPath,
		UploadCovers: false,
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
		return nil, errors.Wrap(err, "Failed to parse settings file. Is it valid?")
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

func (s *Settings) CheckReadwiseConfig() bool {
	return s.ReadwiseToken != ""
}
