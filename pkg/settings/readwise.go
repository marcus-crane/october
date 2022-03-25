package settings

import (
	"github.com/marcus-crane/october/pkg/logger"
)

func (s *Settings) GetReadwiseToken() string {
	s.Lock()
	defer s.Unlock()
	return s.ReadwiseToken
}

func (s *Settings) SetReadwiseToken(token string) error {
	s.Lock()
	defer s.Unlock()
	s.ReadwiseToken = token
	if err := s.Save(); err != nil {
		logger.Log.Errorw("Failed to save Readwise token", "error", err)
		return err
	}
	logger.Log.Infow("Saved Readwise token to storage")
	return nil
}

func (s *Settings) ReadwiseTokenExists() bool {
	return s.ReadwiseToken != ""
}
