# October CLI

As of [v1.9.0](http://localhost:5173/changelog#v1-9-0), October ships with a [CLI interface](https://en.wikipedia.org/wiki/Command-line_interface).

If you are unfamiliar with that means, it's probably not the right tool for you and you can continue to use the existing desktop app as normal.

For those who consider themselves power users, you can use the CLI to quickly sync bookmarks without having to boot up the desktop app and click around.

>[!WARNING]
> At the time of writing, the CLI syncs highlights and that's about it.
>
> The output has been left in a messy state in order to ship something out. It will be brought up to feature parity eventually.
>
> Settings can only be configured via the desktop app in the meantime.

In order to keep distribution easy, and to keep the filesize of October small, the desktop app and CLI app are bundled together rather than being two separate binaries.

The CLI is invoked explicitly with a subcommand, in order to maintain standard OS behaviours when executing binaries, even if it admittedly feels a little weird.

This also ensures that automation tools such as [Keyboard Maestro](https://www.keyboardmaestro.com/) should work with the October CLI where their execution environments don't look like a traditional terminal, so it's better to be explicit than try and guess user intent.

Here is an exhaustive list of supported commands at the time of writing:

```console
$ october cli help
$ october cli sync
```

With all that out of the way, here are some examples of how to use it for the various platforms that October supports.

## Platform usage

### Windows

For Windows, the executable for October lives at `C:\Program Files\utf9k\October\October.exe`.

Here's an example of it in use:

```powershell
PS C:\Users\marcus\Desktop> & 'C:\Program Files\utf9k\October\October.exe' cli sync
time="2024-06-10T18:46:58+12:00" level=info msg="Found an attached device" device_id=00000000-0000-0000-0000-000000000384
time="2024-06-10T18:46:58+12:00" level=info msg="Successfully parsed highlights" batch_count=1 highlight_count=248
time="2024-06-10T18:47:02+12:00" level=info msg="Successfully sent bookmarks to Readwise" batch_count=1
time="2024-06-10T18:47:02+12:00" level=info msg="Successfully synced 248 highlights to Readwise"
```

You may prefer to add the CLI folder to your PATH for easier access eg;

```powershell
PS C:\Users\marcus\Desktop> $env:PATH += ";C:\Program Files\utf9k\October"
PS C:\Users\marcus\Desktop> October.exe cli
NAME:
   october cli - sync your kobo highlights to readwise from your terminal

USAGE:
   october cli [global options] command [command options]

VERSION:
   v1.9.0

AUTHOR:
   Marcus Crane <october@utf9k.net>

COMMANDS:
   sync, s  sync kobo highlights to readwise
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### macOS

For macOS, the executable for October lives at `/Applications/October.app/Contents/MacOS/October`.

Here's an example of it in use:

```console
/Applications/October.app/Contents/MacOS/October cli sync
INFO[0000] Found an attached device                      device_id=00000000-0000-0000-0000-000000000384
INFO[0000] Successfully parsed highlights                batch_count=1 highlight_count=248
INFO[0001] Successfully sent bookmarks to Readwise       batch_count=1
INFO[0001] Successfully synced 248 highlights to Readwise
```

You may prefer to add the CLI folder to your PATH for easier access eg;

```console
$ export PATH=$PATH:/Applications/October.app/Contents/MacOS
$ october cli help
NAME:
   october cli - sync your kobo highlights to readwise from your terminal

USAGE:
   october cli [global options] command [command options]

VERSION:
   v1.9.0

AUTHOR:
   Marcus Crane <october@utf9k.net>

COMMANDS:
   sync, s  sync kobo highlights to readwise
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Linux

For Linux, the executable for October lives at `/usr/local/bin`.

Here's an example of it in use:

```console
$ october cli sync
INFO[0000] Found an attached device                      device_id=00000000-0000-0000-0000-000000000384
INFO[0000] Successfully parsed highlights                batch_count=1 highlight_count=248
INFO[0001] Successfully sent bookmarks to Readwise       batch_count=1
INFO[0001] Successfully synced 248 highlights to Readwise
```

The CLI tool should already be in your path by default on any modern Linux system:

```console
$ october cli help
NAME:
   october cli - sync your kobo highlights to readwise from your terminal

USAGE:
   october cli [global options] command [command options]

VERSION:
   v1.9.0

AUTHOR:
   Marcus Crane <october@utf9k.net>

COMMANDS:
   sync, s  sync kobo highlights to readwise
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```