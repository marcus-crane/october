package updater

import (
	"testing"

	"github.com/blang/semver"
	"github.com/stretchr/testify/assert"
)

func Test_CompareVersions_Older(t *testing.T) {
	expected := false
	currentVersion := semver.MustParse("2.0.0")
	newerVersion := semver.MustParse("1.0.0")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}

func Test_CompareVersions_Newer(t *testing.T) {
	expected := true
	currentVersion := semver.MustParse("1.0.0")
	newerVersion := semver.MustParse("2.0.0")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}

func Test_CompareVersions_Newer_Beta(t *testing.T) {
	expected := true
	currentVersion := semver.MustParse("1.3.0-beta3")
	newerVersion := semver.MustParse("1.3.0-beta5")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}

func Test_CompareVersions_Newer_AlphaToBeta(t *testing.T) {
	expected := true
	currentVersion := semver.MustParse("1.3.0-alpha2")
	newerVersion := semver.MustParse("1.3.0-beta1")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}

func Test_CompareVersions_Newer_Patch(t *testing.T) {
	expected := true
	currentVersion := semver.MustParse("1.3.0-beta3")
	newerVersion := semver.MustParse("1.3.0")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}

func Test_CompareVersions_Dev(t *testing.T) {
	expected := true
	currentVersion := semver.MustParse("0.0.0-DEV")
	newerVersion := semver.MustParse("1.3.0")
	actual, _ := compareVersions(currentVersion, newerVersion)
	assert.Equal(t, expected, actual)
}
