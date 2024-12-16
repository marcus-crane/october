package backend

import (
	"log/slog"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	elipsaDeviceId   = "NXXX,4.9.56,4.30.18838,4.9.56,4.9.56,00000000-0000-0000-0000-000000000387"
	libraTwoDeviceId = "NXXXXXXXXXX,4.1.15,4.30.18838,4.1.15,4.1.15,00000000-0000-0000-0000-000000000388"
	miniDeviceId     = "NXXX,4.9.56,4.30.18838,4.9.56,4.9.56,00000000-0000-0000-0000-000000000340"
	clara2EDeviceId  = "NXXX,4.9.56,4.30.18838,4.9.56,4.9.56,00000000-0000-0000-0000-000000000386"
	unknownDeviceId  = "NXXX,4.9.56,4.30.18838,4.9.56,4.9.56,00000000-0000-0000-0000-000000009999"
)

// We receive a tempdir from the test runner itself, which will handle cleanup for us.
func setupTmpKobo(dir string, deviceId string) string {
	content := []byte(deviceId)
	err := os.Mkdir(filepath.Join(dir, ".kobo"), 0777)
	if err != nil {
		panic(err)
	}
	tmpfn := filepath.Join(dir, ".kobo", "version")
	if err := os.WriteFile(tmpfn, content, 0666); err != nil {
		panic(err)
	}
	return dir
}

func TestGetKoboMetadata_HandleNoDevices(t *testing.T) {
	var expected []Kobo
	actual := GetKoboMetadata([]string{}, slog.New(&discardHandler{}))
	assert.Equal(t, expected, actual)
}

func TestGetKoboMetadata_HandleConnectedDevices(t *testing.T) {
	libraTempDir := t.TempDir()
	fakeLibraVolume := setupTmpKobo(libraTempDir, libraTwoDeviceId)
	miniTempDir := t.TempDir()
	fakeMiniVolume := setupTmpKobo(miniTempDir, miniDeviceId)
	elipsaTempDir := t.TempDir()
	fakeElipsaVolume := setupTmpKobo(elipsaTempDir, elipsaDeviceId)
	clara2ETempDir := t.TempDir()
	fakeClara2EVolume := setupTmpKobo(clara2ETempDir, clara2EDeviceId)
	unknownTempDir := t.TempDir()
	fakeUnknownVolume := setupTmpKobo(unknownTempDir, unknownDeviceId)
	expected := []Kobo{
		{
			Name:       "Kobo Libra 2",
			Storage:    32,
			DisplayPPI: 300,
			MntPath:    libraTempDir,
			DbPath:     filepath.Join(libraTempDir, "/.kobo/KoboReader.sqlite"),
		},
		{
			Name:       "Kobo Mini",
			Storage:    2,
			DisplayPPI: 200,
			MntPath:    miniTempDir,
			DbPath:     filepath.Join(miniTempDir, "/.kobo/KoboReader.sqlite"),
		},
		{
			Name:       "Kobo Elipsa",
			Storage:    32,
			DisplayPPI: 227,
			MntPath:    elipsaTempDir,
			DbPath:     filepath.Join(elipsaTempDir, "/.kobo/KoboReader.sqlite"),
		},
		{
			Name:       "Kobo Clara 2E",
			Storage:    16,
			DisplayPPI: 300,
			MntPath:    clara2ETempDir,
			DbPath:     filepath.Join(clara2ETempDir, "/.kobo/KoboReader.sqlite"),
		},
		{
			MntPath: unknownTempDir,
			DbPath:  filepath.Join(unknownTempDir, "/.kobo/KoboReader.sqlite"),
		},
	}
	detectedPaths := []string{fakeLibraVolume, fakeMiniVolume, fakeElipsaVolume, fakeClara2EVolume, fakeUnknownVolume}
	actual := GetKoboMetadata(detectedPaths, slog.New(&discardHandler{}))
	assert.Equal(t, expected, actual)
}
