package device

import (
	"fmt"

	"github.com/marcus-crane/october/pkg/logger"
	"github.com/pgaskin/koboutils/v2/kobo"
)

func GetKoboMetadata(detectedPaths []string) []Kobo {
	var kobos []Kobo
	logger.Log.Debugf("Found the location of %d possible Kobos", len(detectedPaths))
	for _, path := range detectedPaths {
		_, _, deviceId, err := kobo.ParseKoboVersion(path)
		if err != nil {
			logger.Log.Errorw("Failed to parse Kobo version", "error", err)
		}
		logger.Log.Debugf("Found Kobo with Device ID of %s", deviceId)
		device, found := kobo.DeviceByID(deviceId)
		if !found {
			// We can handle unsupported Kobos in future but at present, there are none
			logger.Log.Debugf("Unrecognised Kobo with device ID of %s", deviceId)
			continue
		}
		logger.Log.Infof(fmt.Sprintf("Detected a %s", device.Name()))
		kobos = append(kobos, Kobo{
			Name:       device.Name(),
			Storage:    device.StorageGB(),
			DisplayPPI: device.DisplayPPI(),
			MntPath:    path,
			DbPath:     fmt.Sprintf("%s/.kobo/KoboReader.sqlite", path),
		})
	}
	return kobos
}
