package device

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupTmpKobo() string {
	content := []byte("NXXXXXXXXXX,4.1.15,4.30.18838,4.1.15,4.1.15,00000000-0000-0000-0000-000000000388")
	dir, err := ioutil.TempDir("", "KOBOeReader")
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer os.RemoveAll(dir)
	os.Mkdir(filepath.Join(dir, ".kobo"), 0777)
	tmpfn := filepath.Join(dir, ".kobo", "version")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
		return ""
	}
	return dir
}

func TestGetKoboMetadata_HandleNoDevices(t *testing.T) {
	var expected []Kobo
	actual := GetKoboMetadata([]string{})
	assert.Equal(t, expected, actual)
}

func TestGetKoboMetadata_HandleConnectedDevice(t *testing.T) {
	fakeKoboVolume := setupTmpKobo()
	expected := []Kobo{
		{
			Name: "Kobo eReader",
		},
	}
	detectedPaths := []string{fakeKoboVolume}
	actual := GetKoboMetadata(detectedPaths)
	assert.Equal(t, expected, actual)
}
