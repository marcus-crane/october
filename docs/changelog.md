# Changelog

## v1.8.0

This release adds support for Linux machines running arm64 such as newer Raspberry Pi devices apparently.

It has been tested using an Ubuntu arm64 VM running on a macOS Silicon device via [Parallels](https://www.parallels.com/) but any feedback from users with real Linux arm64 devices would be appreciated.

Some general improvements for the Linux `.deb` packages has been done too, mainly defining required dependencies that can cause October to not boot within certain distros, such as [pop_OS!](https://pop.system76.com/) due to dev libraries not being installed out of the box.

Behind the scenes, a bunch of work has been done to slim down the release pipeline and generate build artifacts against pull requests, which should greatly speed up the feedback loop for development.

## v1.7.0

This release mostly contains maintenance fixes that have cropped up since the last release as well as a new distribution method.

I'm looking to get back to some feature work during the quieter Christmas period but first I had to fix some breakages introduced by Apple changing how code signing works.

- Removed dependency on `CGO` for database access which will make future maintenance much easier
- Patched a race condition that can lead to an extra closing `}` being added to settings files
- Added a small bit of logic that tries to repair settings files that may have been corrupted by the previously mentioned bug
- Fixed up release pipelines which were affected by Apple code signing changes

Additionally, October is now available via Homebrew going forward for macOS users.

It can be installed by running `brew update` to get the latest Cask definition followed by `brew install --cask october`.

![CleanShot 2023-12-05 at 10 17 50@2x](https://github.com/marcus-crane/october/assets/14816406/50ca1407-42b8-4ec0-85b0-5555f39f06a3)

## v1.6.0

This version introduces a portable version of October for Windows users who may want to sync highlights from their Kobo on a managed machine where they can't otherwise install software.

You might also like to install it on a USB so you can use October across multiple computers without having to set it up each time.

Come to think of it, I suppose you could technically copy this version to your Kobo itself (which is basically a glorified USB) and carry it around that way. That said, I haven't tested this out so your mileage may vary! I think it would work fine though.

---

Upon startup, it will create a folder next to the exe containing settings files so just make sure to keep the exe and folder together.

While I haven't tested it extensively, one existing user seemed to have no problems using it.

## v1.5.0

This release brings support for Linux systems which means that October is available for all major platforms.

![](./public/linux/linux_overview_light.png)

To celebrate this milestone, I've given the docs site a bit of a content refresh although there are still more changes I intend to make.

There were also a few other small changes:

- Fixed pesky sizing issue that means you could scroll down slightly on most screens. The Windows settings screen specifically may still have this issue depending on your screen resolution
- Updated artifact file format naming to follow a consistent standard
- Slightly increased the default window width for October to give it a little more breathing room

## v1.4.0

_Note: This release was originally accidentally published as `v.1.4.0` instead of `v1.4.0`. The contents has not changed but the release was recreated as `v.1.4.0` is not a valid Windows version identifier._

This release adds back the ability to sync store-purchased titles, although it needs to be explicitly enabled in Settings just to ensure users aren't accidentally double syncing if they are also using the official Readwise integration.

![CleanShot 2023-06-03 at 15 08 55@2x](https://github.com/marcus-crane/october/assets/14816406/e0897f79-0347-4ea8-b955-3f94f546f16b)

In order to bring awareness, the first initial sync will also show this popup if highlights are detected that are not from sideloaded titles. This popup will only show one time. Either clicking "Understood" or enabling store-purchased highlight syncing via Settings will cause it to disappear.

<img width="1136" alt="CleanShot 2023-06-03 at 15 33 07@2x" src="https://github.com/marcus-crane/october/assets/14816406/87cbeff0-c4f2-414a-a26d-c03228028c74">

Lastly, the upload button now lists how many side loaded and official highlights are detected

This release also fixes some buginess:

- [v1.3.2](https://github.com/marcus-crane/october/releases/tag/v1.3.2) added logic to only upload sideloaded books but in doing so, some queries for discovering Kobo content were incorrect. These should be fixed now so the correct number of sideloaded titles are uploaded.

## v1.3.4

This release adds detection support for `Kobo Clara 2E` and `Kobo Elipsa 2E`.

As always, unreleased devices should work just fine with October. This change just adds their device strings so that October can display their names instead of "Unknown Kobo".

I don't own these devices so I don't know for sure that this will work but I pulled the device IDs from the latest Kobo firmware so I assume they should work just fine.

If they don't and you own one of these devices, please feel free to file an issue.

## v1.3.3

This release doesn't add any new features but brings a bunch of maintenance chores that were overdue.

- Refreshes the dark mode colour scheme which had been partly implemented. It should now be much less jarring but it can probably do with more tweaks in future
- Fixes #85 where starting October on Windows 11 with newer versions of the WebView2 runtime would briefly cause an "Unexpected Error" screen to appear

There are some other changes that won't be visible to users but for those interested:

- Removes unneeded dependency on `github.com/go-resty/resty` in favour of plain `net/http`. This has no visible change for users. It may or may not reduce the size of October very slightly but I haven't checked. It probably doesn't.
- Updated most dependencies to their latest versions which brings October up to date and provides a fresh base to start developing again.
- Begin signing releases with a new developer certificate as the current version expires on May 30th. This has no impact to end users. It just means that I can continue signing new installers (so they pass Windows Smartscreen) as my current certificate is about to expire. The new signature lasts until 2025 when it will need to be renewed again.
- Updated year references from `2022` to `2023` that were missed. It doesn't mean anything in a legal sense but it's always nice when software appears fresh, such as referring to the current year.

<details>
<summary>Toggle to see screenshots of the new dark mode scheme</summary>

<img width="1136" alt="CleanShot 2023-05-28 at 16 21 16@2x" src="https://github.com/marcus-crane/october/assets/14816406/77068de9-f6b7-44b8-8219-ed34840741f7">

<img width="1136" alt="CleanShot 2023-05-28 at 16 21 22@2x" src="https://github.com/marcus-crane/october/assets/14816406/332ffc41-c391-4afe-ae44-3497d5a2bf7c">

<img width="1136" alt="CleanShot 2023-05-28 at 16 21 36@2x" src="https://github.com/marcus-crane/october/assets/14816406/88f77ca5-be2d-4544-b6bd-31697c932a86">

</details>

## v1.3.2

This release updates the various bookmark/book queries so that store-purchased Kobo titles are ignored.

For most users, they'll want to use the officially supported Readwise integration to make use of remote syncing and what not.

October is intended to just cover the sideloading part of the equation but until now, there was never any special handling of store-purchased titles so they happened to be picked up if they existed.

For any users who have synced store-purchased titles previously, and have duplicate highlights, they'll need to manually remove them from their Readwise accounts but this is a one time thing.

## v1.3.1

This release batches highlights to Readwise into multiple requests if you are trying to send more than 2000 requests at a time.

If you're trying to upload a large batch of highlights in one go, this should now be possible without the risk of the Readwise API rejecting the request due to it being too large.

## v1.3.0

This release refactors all of the internal logging, dropping [zerolog](https://github.com/rs/zerolog) in favour of trusty [logrus](https://github.com/Sirupsen/logrus) which is not user facing at all.

As a result, logging now works properly again on Windows closing #25

While going through all the various log entries, I've also switched the logging format to JSON which makes it easier to parse through each file. Some log entry sizes have also been greatly reduced where before they would log out an entire struct even if the whole thing wasn't useful. This should greatly reduce the size of the average log file.

In order to making it easier for users to submit logs for investigations, there is now a button in the Settings view that will open your computer's file explorer to the location that October stores logs in (#12)

**Windows**
![windows-logs](https://user-images.githubusercontent.com/14816406/206837799-237f9dbd-74eb-4530-9e79-4e45b87059e8.png)

**macOS**
![macos-log](https://user-images.githubusercontent.com/14816406/206837886-755d068d-a505-4ae6-a5dc-8d948e15f953.png)

You might also notice a couple of other small additions in the screenshots above.

There is now a little build identifier in the Settings view where you can see build information about the version of October you're running (#65) as well as a button for reporting bugs which will open a new Github issue pre-populated with said build information to make it easier to provide support to end users.

I've still got some other refactoring tasks to do which should hopefully then make it easier to add in some of the larger features that have been outstanding such as better handling of upload failures as well as cutting down the time to upload

## v1.2.0

This release does not contain any new features but tidies up a lot of underlying metadata associated with each executable, thanks in part to a bunch of upgrades that were introduced with Wails 2.0, the Go framework that powers Wails.

**NOTE:** For Windows users, I recommend uninstalling October before installing this latest version just to avoid any confusion. Nothing will break if you decide not to but as some metadata has changed, using the latest installer will change the install location resulting in two October entries. You can install the older version at any point but it may be confusing if you aren't aware of this change. I figure best to do it now than wait until it might impact more users in future.

At a high level, this update corrects the version information that is available in various places as well as adding some publisher information and other things that you would generally not need to be aware of but can be handy.

I don't expect this to be that interesting to most users but if you'd like to read the nitty gritty details, feel free to keep reading:

### Windows

#### Install path

Publisher information has been set to `utf9k` which is just the sort of "catch all" that I use for my software side projects and happens to be the domain of my website as well.

This in itself isn't particularly interesting but it does inform the installation path for October so for Windows users, October is now installed at `C:\Program Files\utf9k\October\october.exe`. Before this change, it was installed at `C:\Program Files\October\October\october.exe`.

While I don't distribute any other Windows software, this just helps with grouping it all together in a standard place in line with how Windows expects applications to be laid out.

![installer-name](https://user-images.githubusercontent.com/14816406/204116218-b0e23130-651f-4ed3-9390-f3dab5f49e48.png)

In order to remove the previous version, you can either run the uninstaller located at `C:\Program Files\October\October\uninstall.exe` or you can use the `Add or remove programs` section in Control Panel. You'll be able to spot the proper version as all versions before this one will show with the version set as `1.0.0`

![uninstall](https://user-images.githubusercontent.com/14816406/204116312-ff627bba-3479-4fa9-97f8-90c2dd3c8a33.png)

This only needs to be done one time as future installations of October will always be installed at `C:\Program Files\utf9k\October` going forward.

#### Version information

Something I hadn't realised until it was mentioned in #65 is that there isn't actually a clear way at all to identify version information. This still has some work to be done, such as surfacing it within October itself but now there are a few places that can be checked to see what version you are using.

**NOTE:** Due to limitations in how versions are laid out, they are always represented as `x.x.x` where x is a number regardless of any extra versioning. For example, both `1.2.2-alpha1` and `1.2.2-beta2` would be represented as `1.2.2` due to what Windows (and installer files) expect. That said, I don't expect any users to actually run any pre-release versions as they're most for my own testing to simulate what the final release may look like.

The first place that is updated is `Add or remove programs` which now reflects the installed version, in this case `1.2.0`. It also reflects an updated publisher name as well.

![installed-compare](https://user-images.githubusercontent.com/14816406/204116360-d0728a5a-76d0-4705-af81-30e87040d06f.png)

Right clicking on the October executable and selecting `Properties -> Details` will also show relevant versioning information as well.

![windows](https://user-images.githubusercontent.com/14816406/204116626-31e51889-9d6c-42f0-b3c0-4f0b17d9e39e.png)

If there are any other Windows metadata areas that I'm not aware of and have missed, please let me know.

### macOS

Historically the macOS app has had a little "About October" popup that is accessible from the application menu, as is standard with all macOS applications.

<img width="372" alt="CleanShot 2022-11-27 at 15 30 32@2x" src="https://user-images.githubusercontent.com/14816406/204116557-900f59ce-72e8-4502-a885-62c62d05e41c.png">

We can do better than that though and there is now extra metadata visible when viewing `October.app` within your `Applications` folder:

Right clicking on the application and selecting "Get Info" will now show a bunch of updated information such as the installed version.

![CleanShot 2022-11-27 at 15 06 56@2x](https://user-images.githubusercontent.com/14816406/204116460-7acf1c2a-a44b-42cd-bae0-fc30e45359dc.png)

Similarly, pressing Spacebar with `October.app` highlighted will also show a bunch of information using Finder's preview functionality

![CleanShot 2022-11-27 at 15 06 38@2x](https://user-images.githubusercontent.com/14816406/204116458-4d549606-8084-4a79-99e8-96eb44d2f261.png)

There aren't any other relevant areas of macOS that I'm aware of to check as applications are very self contained (ie; there isn't a messy install process) unlike with Windows and as such, there is no need to worry about install paths changing or anything like that.

## v1.1.2

This is a minor update to fix an edge case that some users have run into in the past.

While Readwise does not impose any limits on how many highlights you can have, there is a system limit that highlights can't be any longer than 8191 characters long.

Generally speaking, if you're making highlights that long, you're missing the point of highlights but I've seen occasional reports of Kobo highlighting accidentally capturing an entire chapter.

Without this being obvious to the end user, they'll be confused when October appears to fail (as any error will cause the upload process to fail currently) so it's better to work around this issue than cause 99% of valid highlights to fail to upload I think.

The real fix for this is to delete those highlights on your device (find the highlight, tap it and hit delete) but for users who might want to make such long highlights for whatever reason, October will now split your highlight text into appropriate chunks.

## v1.1.1

This version adds a fallback for highlights that are missing the `DateCreated` field, which causes October to fail to continue processing.

In the event that `DateCreated` is missing, October will use the `DateModified` field. If both are somehow missing, it'll do a further fallback and use the current date.

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
