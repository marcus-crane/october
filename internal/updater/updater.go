package updater

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

const (
	UPDATE_SOURCE = "marcus-crane/october"
)

func PerformUpdate(currentVersion string) (bool, error) {
	v, err := semver.Parse(currentVersion)
	if err != nil {
		return false, err
	}
	latest, err := selfupdate.UpdateSelf(v, UPDATE_SOURCE)
	if err != nil {
		return false, err
	}
	if latest.Version.Equals(v) {
		// Binary is up to date with latest version
		// LOG: Current version is up to date
		return true, nil
	}
	// Successfully updated to new version
	// LOG: Successfully updated to version blah
	return true, nil
}

func PerformUpdateDarwin(currentVersion string) (bool, error) {
	latest, found, err := selfupdate.DetectLatest(UPDATE_SOURCE)
	if err != nil {
		return false, err
	}
	if !found {
		return false, nil
	}

	// Happy path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return false, err
	}

	downloadPath := filepath.Join(homeDir, "Downloads", "October.zip")
	if err := exec.Command("curl", "-L", latest.AssetURL, "-o", downloadPath).Run(); err != nil {
		// LOG: CURL error
		return false, err
	}
	var installPath string
	cmdPath, err := os.Executable()
	installPath = strings.TrimSuffix(cmdPath, "October.app/Contents/MacOS/October")
	if err != nil {
		installPath = "/Applications/"
	}
	if err := exec.Command("ditto", "-xk", downloadPath, installPath).Run(); err != nil {
		// LOG: Ditto error
		return false, err
	}
	if err := exec.Command("rm", downloadPath).Run(); err != nil {
		// LOG: Failed to cleanup tmp folder
		return false, err
	}
	return true, nil
}

func CheckForNewerVersion(currentVersion string) (bool, string) {
	latest, found, err := selfupdate.DetectLatest(UPDATE_SOURCE)
	if err != nil {
		return false, ""
	}

	if !found {
		// LOG: Update manifest not found
		return false, ""
	}

	v, err := semver.Parse(currentVersion)
	if err != nil {
		return false, ""
	}

	return compareVersions(v, latest.Version)
}

func compareVersions(currentVersion semver.Version, latest semver.Version) (bool, string) {
	if latest.LTE(currentVersion) {
		return false, ""
	}
	return true, latest.String()
}
