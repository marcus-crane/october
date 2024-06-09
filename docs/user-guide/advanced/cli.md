# October CLI

As of [v1.9.0](http://localhost:5173/changelog#v1-9-0), October ships with a [CLI interface](https://en.wikipedia.org/wiki/Command-line_interface).

If you are unfamiliar with that means, it's probably not the right tool for you and you can continue to use the existing desktop app as normal.

For those who consider themselves power users, you can use the CLI to quickly sync bookmarks without having to boot up the desktop app and click around.

>[!WARNING]
> At the time of writing, the CLI syncs highlights and that's about it.
>
> The output has been left in a messy state in order to ship something out. It will be brought up to feature parity eventually.
>
> Settings can only be configured via the desktop app, although you can edit the config files directly which will be touched on below.

In order to keep distribution easy, and keep the file size of October small, the desktop app and CLI app are functionally the same thing.

The CLI is invoked explicitly with a subcommand, in order to maintain standard OS behaviours when executing binaries, rather than relying on a magic check to see if the executable is within a terminal or not.

This also ensures that automation tools such as [Keyboard Maestro](https://www.keyboardmaestro.com/) should work with the October CLI where their execution environments don't look like a traditional terminal so it's better to be explicit than try and guess user intent.

## Windows

## macOS

For macOS, most apps live in `/Applicatons` by default and October is no different.

The executable for October lives at `/Applications/October.app/Contents/MacOS`

```console
$ /Applications/October.app/Contents/MacOS/October --help
NAME:
   october - sync your kobo highlights to readwise from your terminal

USAGE:
   october [global options] command [command options]

VERSION:
   v1.9.0-beta2

AUTHOR:
   Marcus Crane <october@utf9k.net>

COMMANDS:
   launch   skip cli tool and launch desktop ui
   sync, s  sync kobo highlights to readwise
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

## Linux