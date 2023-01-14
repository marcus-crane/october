package kobo

import (
	"github.com/pgaskin/koboutils/v2/kobo"
)

// FindMountDevices runs a variety of detection methods for different operating systems (Linux, macOS, Windows)
// returning any results it can find. Those results are then checked for the existence of a .kobo
// folder which is used to determine whether a device is valid or not.
func FindMountedDevices() ([]string, error) {
	confirmedLocations := []string{}
	locations, err := kobo.Find()
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		if kobo.IsKobo(location) {
			confirmedLocations = append(confirmedLocations, location)
		}
	}
	return confirmedLocations, nil
}

// GetDeviceMetadata looks up the connected Kobo by checking the version string located at <kobo>/.kobo/version
// An uninitialised KoboConnection will be returned populated with identifiers like device name, storage etc
// If a device is unknown (ie; released before we support it), a minimally populated KoboConnection will be
// returned which should still contain enough information for October to still connect regardless.
func GetDeviceMetadata(path string) (KoboConnection, error) {
	serial, version, deviceId, err := kobo.ParseKoboVersion(path)
	if err != nil {
		return KoboConnection{}, err
	}
	device, known := kobo.DeviceByID(deviceId)
	if !known {
		// We can handle unsupported Kobos that release in future, before support is added to koboutils
		// but we should tell the user that no support is 100% guaranteed if there are DB changes etc
		// We only read data anyway so there is very little risk of just trying out best
		return KoboConnection{
			Name:      "Unknown Device",
			MountPath: path,
			DbPath:    formatUsualDbPath(path),
		}, nil
	}
	return KoboConnection{
		Name:       device.Name(),
		Storage:    device.StorageGB(),
		DisplayPPI: device.DisplayPPI(),
		MountPath:  path,
		DbPath:     formatUsualDbPath(path),
		Serial:     serial,
		Version:    version,
		DeviceId:   deviceId,
	}, nil
}
