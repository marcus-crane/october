package settings

import (
	"github.com/marcus-crane/october/pkg/logger"
)

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

func (s *Settings) SetCoverUploadStatus(enabled bool) error {
	s.Lock()
	defer s.Unlock()
	s.UploadCovers = enabled
	if err := s.Save(); err != nil {
		logger.Log.Errorw("Failed to set cover upload status", "error", err)
		return err
	}
	logger.Log.Infow("Saved cover upload status to storage")
	return nil
}

func (s *Settings) ReadwiseTokenExists() bool {
	return s.ReadwiseToken != ""
}
