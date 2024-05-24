# Linux

For Linux users, there is one build available:

- amd64 (.deb)

While I'm only personally using a Debian flavour machine myself, if there is a strong enough interest in either alternative distro formats (ie; `rpm`) or in alternative packaging formats (ie; `AppImage`, `flatpak`), I'd be happy to consider looking into these.

![](../public/linux/linux_overview_light.png)
![](../public/linux/linux_overview_dark.png)

## Installation

At the time of writing, the Linux build doesn't have any build requirements so it should be safe to install by way of `dpkg`.

Having said that, I've spent the least time with the Linux build to date so the installation process could need some improvement. To my knowledge, everything should be compiled in though.

The quickest way to install is by grabbing the latest `.deb` build from [Github](https://github.com/marcus-crane/october/releases) or [from this shortcut](https://october.utf9k.net/download/linux/latest) and installing it like so:

```console
$ sudo dpkg -i ~/Downloads/october_x.x.x_linux_amd64.deb
```

![](../public/linux/linux_install.png)

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