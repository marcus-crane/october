---
template: overrides/main.html
---

# Changelog

## v1.0.2

This release updated Windows binaries to be built with `CGO_ENABLED=1`. Without it, Windows builds failed to be usable end to end. They should work properly now.

## v1.0.1

This was a minor fix for an issue that was blocking October from compiling.

## v1.0.0

BREAKING CHANGE: A one time breaking change was made where additional metadata is submitted to Readwise. This resulted in any books uploaded pre-1.0.0 with October being duplicated. This change was used to properly match books on your device so covers aren't overwritten in future.

This release marked the first official release of October, as far as stability and error handling.

It added the ability to upload cover images to Readwise and a button for checking connectivity with Readwise (ie; is your token valid)

## v0.9.4-post2

This release disables `upx` as it can cause some unforseen issues. I haven't witnessed them but I'd like to avoid them all the same.

## v0.9.4-post1

Nothing new with this release but Windows builds are now installed via a NSIS installer.

As a result of this, users now have a desktop icon and start menu entry upon installation as well.

As always, please submit any feedback via Github Issues

## v0.9.4

This release fixes up an issue with Windows builds not working. As a result of adding logging, I didn't realise that zap doesn't initialize properly on Windows.

I've implemented a workaround and it seems to be working fine.

The README has also been tidied up with some Windows information and screenshots.

## v0.9.3

This release unlocks the previous hard requirement on books having a title. Sideloaded books with no title will now firstly fallback to extracting the epub filename.

For example, if you have a book like `cool author - interesting book.epub`, your Readwise "book" will be called `cool author - interesting book`. If that fails somehow, it will fall back again to just omitting the title/sending an empty string.

What Readwise will do then is creating a "book" called "Quotes".

NOTE: While you can change the title of your book in Readwise, this will cause future uploads to reduplicate highlights under the old title. For now, if you can leave the titles as is, I'll attempt to address this shortly.

As always, please let me know if you have any issues so I can add them to my v1.0.0 backlog.

## v0.9.2

This release adds basic logging in the following places. It's good enough for providing support but is intended to be cleaned up (and exporting via the UI) before v1.0.0 releases.

October should save logs in the following places:

Windows: `%APPDATA%\Local\october\logs`
macOS: `$HOME/Library/Application Support/october/logs`
Linux: `/usr/local/share/october/logs`

## v0.9.1

This release doesn't add anything new functionality wise but sets up a build pipeline that does the following:

* Builds Windows + macOS binaries
* Zips up Windows `.exes`
* Notarises macOS `.app` and packages it up as a mountable `dmg`
  * This means no more quarantining of the macOS app

## v0.9.0

This is the first public release of October. It's marked as `v0.9.0` as while it works end to end, the codebase is a bit messy. People are free to use it but shouldn't expect it to be "released" until `v1.0.0`.

I also need to set up proper CI/CD infrastructure so excuse the bare `.app` files zipped up for now.

Additionally, I intend to codesign October so you don't need to fiddle with the security settings.
