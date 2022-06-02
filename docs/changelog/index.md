---
template: overrides/main.html
---

# Changelog

## v1.1.0

This release brings quite a few changes although a lot of them are under the hood so you won't notice them but they'll make it easier to add new features to October going forward.

**Codesigning on Windows**

First of all, October for Windows is now codesigned meaning there should no longer be any scary warnings going forward. As this version is the first to be codesigned, and my developer certificate is now, Windows SmartScreen is expected to appear for a brief period of time while trust is established against October but once that's done, new releases should no longer trigger a warning. You can read more about this addition in [this announcement](https://github.com/marcus-crane/october/discussions/54)

**Internal refactoring**

As mentioned, a lot of changes have happened under the hood. October is powered by [Wails](https://wails.io/) and while understanding how to use the latest version, I rearranged the internals of the project and the latest iteration ended up with everything being a bit too far apart resulting in duplicate code and other things. Without boring you with the details, everything is now "closer together" making it easier to do changes going forward.

The frontend now also uses [Vite](https://vitejs.dev/) and with that comes live reloading which means any developer changes will recompile the application. This means it becomes much faster to test out changes (and ultimately to release them as well!)

**Support for unreleased Kobo devices**

Under the hood, October uses [pgaskin/koboutils](https://github.com/pgaskin/koboutils) which currently has support for all released Kobo devices. At one point however, it did not and it may lag behind when new devices are refreshly released. October effectively just uses `koboutils` to get metadata about devices so I've updated the device selector to allow you to select unrecognised Kobos instead of just ignoring them.

![unknownkobo](https://user-images.githubusercontent.com/14816406/171544836-41ad52b2-6222-410f-95d8-1a85c43c663d.png)

**Condensed settings**

The Settings page has been condensed a little bit and the Readwise token link will now actually open that window in your browser whereas before it was just a piece of text.

![settings](https://user-images.githubusercontent.com/14816406/171545072-a29ca661-3321-4a39-b549-0c620b359d30.png)

In future, I'd like to add the ability to both export and view system logs from within October itself for any advanced users who might like to try and diagnose their own issues.

**Updated toasts**

The library I was using for toasts wasn't the nicest so I've swapped out `react-toastify` for `react-hot-toast` which has cut down on a bunch of code. As a result, toasts now appear at the top of the screen and take up much less space visually.

![toast](https://user-images.githubusercontent.com/14816406/171544879-704be58d-3d74-4f48-aecd-26eeeb0ce2f4.png)

**What's next**

As you'll notice, there wasn't much new in this release as a lot of it was spent on things behind the scenes. I think October is in a good spot to start extending out the UI so users have much more control over highlight uploading instead of just a big sync button.

As a bit of a teaser, here's a screenshot of something I threw together in a short period of time. I was starting straight on a v2 but I decided it's better (and safer) to do things piece by piece rather than doing a big bang release that may ultimately introduce more bugs. Admittedly, doing the best thing isn't as fun though.

![v2](https://user-images.githubusercontent.com/14816406/171545791-76510be9-f640-46fe-a3d9-f88cfa740fed.png)

This isn't necessarily what an updated October might look like but while this is using an email template, it is rendering real data from my Kobo that was connected to my computer. Ideally if one highlight fails, it shouldn't cause your entire upload process to be blocked so that's something I'd like to move towards next.

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
