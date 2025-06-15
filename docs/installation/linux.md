# Linux

For Linux users, there are now multiple package formats available:

- **Debian/Ubuntu**: `.deb` packages (amd64, arm64)
- **Fedora/RHEL/CentOS**: `.rpm` packages (amd64, arm64)

## Installation

### Debian/Ubuntu (.deb packages)

> [!NOTE]
> The default Linux `.deb` is built against webkit2gtk-4.1 as that is the version that ships with most recent Distros. If you have no idea what that is and you're on a recent Linux distro, you should use the default package.
>
> If you're on an older distro and/or you know that you need webkit2gtk-4.0, a separate `.deb` package is available.

The quickest way to install is by grabbing the latest `.deb` build from [Github](https://github.com/marcus-crane/october/releases) or [from this shortcut](https://october.utf9k.net/download/linux/latest) and installing it like so:

```console
$ sudo dpkg -i ~/Downloads/october_x.x.x_linux_amd64.deb
```

# Fedora / RPM Installation

October now provides native RPM packages for Fedora and other RPM-based distributions like RHEL, CentOS, and openSUSE.

## Installation

### From GitHub Releases

1. Download the latest RPM package from [GitHub Releases](https://github.com/marcus-crane/october/releases)
2. Install using dnf (Fedora) or your distribution's package manager:

```bash
# Fedora
sudo dnf install october-*.rpm

# RHEL/CentOS
sudo yum install october-*.rpm

# openSUSE
sudo zypper install october-*.rpm
```

### Dependencies

The RPM package automatically handles dependencies. On modern distributions (Fedora 35+, RHEL 9+), it requires:
- `gtk3`
- `webkit2gtk4.1`

On older distributions, it will use:
- `gtk3` 
- `webkit2gtk3`

## Architecture Support

RPM packages are available for:
- x86_64 (AMD64)
- aarch64 (ARM64)

## Desktop Integration

The RPM package includes:
- Desktop entry file for application launcher integration
- Application icons in multiple sizes
- AppData metadata for software centers
- Automatic desktop database and icon cache updates

## Uninstallation

To remove October:

```bash
# Fedora
sudo dnf remove october

# RHEL/CentOS  
sudo yum remove october

# openSUSE
sudo zypper remove october
```

## Technical Details

### File Locations

- **Binary**: `/usr/local/bin/october`
- **Desktop file**: `/usr/share/applications/october.desktop`
- **Icons**: `/usr/share/icons/hicolor/*/apps/october.png`
- **Metadata**: `/usr/share/metainfo/net.utf9k.october.appdata.xml`

### User Data

User settings and logs are stored in standard XDG directories:
- **Settings**: `$HOME/.config/october/config.json`
- **Logs**: `$HOME/.local/share/october/logs/`

These are not removed when uninstalling the package. 

![](../public/linux/linux_overview_light.png)

## Architecture Support

Both package formats are available for:
- **x86_64** (AMD64) - Intel/AMD 64-bit processors
- **aarch64** (ARM64) - ARM 64-bit processors

## Desktop Integration

Both DEB and RPM packages include:
- Desktop entry file for application launcher integration
- Application icons in multiple sizes
- AppData metadata for software centers
- Automatic desktop database and icon cache updates

It should install a `.desktop` file in the proper place as well, allowing for quick access via the system launcher.

Please [let me know](mailto:october@utf9k.net) if this doesn't appear to work with your desktop environment of choice and I can have a look.

## Technical Details

If you're curious about any files that are generated or need to manually wipe October from existence, here are the following places that files are created:

### Application

The main October app lives at `/usr/local/bin/october` and is self contained as far as the shipped binary is concerned.

Some supporting material for the desktop entry lives in the following places:

- `/usr/share/applications/october.desktop`
- `/usr/share/icons/hicolor/<num>x<num>/apps/october.png` -> [View all icon size variations](https://github.com/marcus-crane/october/tree/main/build/linux/october_0.0.0_amd64/usr/share/icons/hicolor)

You can uninstall the base binary by running the following command:

```console
$ sudo dpkg -r october
```

### Logs

These are stored at `$HOME/.local/share/october/logs` with one log file created each time the application is launched.

Logs are in [ndjson](http://ndjson.org/) format and use the naming convention of `<current unix timestamp>.json`.

### Settings

October's internal settings are stored in a JSON file that lives at `$HOME/.config/october/config.json`.