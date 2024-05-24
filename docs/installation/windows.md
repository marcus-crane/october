# Windows

For Windows users, there is one build and one variant available:

- amd64 (NSIS installer + portable exe for USB drives)

![](../public/windows/windows_overview_light.png)

> [!TIP]
> While October should work on older versions of Windows, it has only officially been tested on Windows 11.
>
> You may be prompted upon first start to install [WebView2](https://developer.microsoft.com/en-us/microsoft-edge/webview2/) which is a component of Windows used by October.
>
> It is included by default on Windows 11 but will need to be manually installed on older Windows versions.

## Installation

For experienced Windows users, there's nothing fancy about the installation process. It's a stock standard install wizard.

To get started, you'll want to get a copy of the latest installer from [Github Releases](https://github.com/marcus-crane/october/releases) which has the file format of `october_<version>_windows_amd64.zip`.

Once downloaded, you'll find an application installer within the zip file.

Running it will present you with a fairly standard install wizard.

<center>

![](../public/windows/windows_installer_location.png)

</center>

As mentioned, there aren't any custom options so you'll just want to click Next until the process is complete.

You should find an entry in your start menu and a shortcut on your desktop.

## Technical Details

If you're curious about any files that are generated or need to manually wipe October from existence, here are the following places that files are created:

### Application

The main October app lives at `C:\Program Files\utf9k\October` and is self contained as far as the shipped binary is concerned.

You can uninstall it by searching for October in the start menu, and selecting `Uninstall` or finding it via `Add or Remove Programs` in your system settings.

### Logs

These are stored at `C:\Users\<username>\AppData\Local\october\logs` with one log file created each time the application is launched.

Logs are in [ndjson](http://ndjson.org/) format and use the naming convention of `<current unix timestamp>.json`.

### Settings

October's internal settings are stored in a JSON file that lives at `C:\Users\<username>\AppData\Local\october\config.json`.