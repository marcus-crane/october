package device

import (
	"fmt"

	"github.com/pgaskin/koboutils/v2/kobo"
	"github.com/rs/zerolog/log"
)

func GetKoboMetadata(detectedPaths []string) []Kobo {
	var kobos []Kobo
	log.Debug().Int("path_count", len(detectedPaths)).Msg("Found the location of possible Kobo(s)")
	for _, path := range detectedPaths {
		_, _, deviceId, err := kobo.ParseKoboVersion(path)
		if err != nil {
			log.Error().Err(err).Msg("Failed to parse Kobo version")
		}
		log.Debug().Str("device_id", deviceId).Msg("Found Kobo")
		device, found := kobo.DeviceByID(deviceId)
		if !found {
			// We can handle unsupported Kobos in future but at present, there are none
			log.Debug().Str("device_id", deviceId).Msg("Found an unrecognised Kobo")
			continue
		}
		log.Info().Str("device_name", device.Name()).Msg("Identified Kobo")
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
