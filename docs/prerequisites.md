# Prerequisites

Before we get started with the desktop app itself, there are a few things that you'll benefit from knowing before you start uploading books only to have some frustrations later on.

This guide is focused around Calibre as it's the most popular tool for ebook management.

There are some defaults we'll want to change in order to have the best possible experience with October, such as being able to support cover uploads, which are disabled by default.

## Background

For those who like to curate their Readwise library, it can be quite nice to have book covers uploaded to fill out your shelf.

<center>
![](assets/covers.png)
</center>

The good news is that October does indeed support cover uploading but by default, your Kobo compresses them so much that they're... blue.

<center>
![](assets/blue.png)
</center>


!!! INFO
    The covers aren't really intended to be blue. Normally, you would view this image in greyscale on your Kobo but what you're seeing here is how it looks after being extremely compressed to save space. I wouldn't be surprised if this heavily distorted image would actually look much better than a plain black and white image too, when uploaded to a Kobo.

In order to make sure that we retain images properly, we need to tell the Kobo to not try to handle images on our behalf.

## Configuring Calibre to sync high quality covers

For the unfamiliar, [Calibre](https://calibre-ebook.com/) is basically the go-to tool for eReader management.

If you aren't already using it, I highly recommend it.

October should still work fine with regular ePubs uploaded to your Kobo via USB drag-and-drop but there are no guarantees when it comes to cover art.

Once you've plugged in your device to your computer via USB, and accepted the connection request on your Kobo, you should open Calibre.

It should detect your device but this can often take a bit of time. Anywhere up to 30 seconds is usually reasonable.

You'll know your device is ready when Calibre visually reloads and the button labelled "Device" appears in the middle of header toolbar.

From here, you'll want to click on the little arrow to the right of the "Device" button, not on the button itself.

![](./assets/calibre_device.png)

You should see a dropdown appear with a few options.

Click on the one labelled "Configure this device".

![](./assets/calibre_configure.png)

From here, a dialogue box with a bunch of options relating to your Kobo should appear.

You'll want to click on the middle tab labelled "Collections, covers & uploads" and then make sure the "Upload covers" checkbox in the middle of the window is ticked.

It's probably worth ticking "Keep cover aspect ratio" while you're here as well.

![](./assets/calibre_covers.png)

Once that's done, you'll see a dialogue box pop up saying that you'll want to restart Calibre in order to apply your changes.

![](./assets/calibre_restart.png)

The very last thing is to enable covers in October itself.

!!! NOTE
    If you're reading this documentation in order, we haven't actually installed October just yet.
    
    Chicken and egg problems are hard.

You can find the cover upload toggle in October under the "Settings" button in the top left once you've selected your Kobo.

![](../assets/settings/settings_coveruploads_light.png#only-light)
![](../assets/settings/settings_coveruploads_dark.png#only-dark)

If you're reading this documentation in order, this step will be mentioned again so you don't forget to turn it.

## Uploading highlights from store-purchased books

While October is primarily aimed at syncing highlights from [sideloaded](https://en.wikipedia.org/wiki/Sideloading) titles, it does work with titles purchased from the Kobo store as well.

This behaviour is disabled by default as Readwise has [an official integration](https://help.readwise.io/article/135-how-do-i-import-highlights-from-kobo) to support syncing with your Kobo account, which is recommended over October.

It should allow the ability to sync highlights wirelessly which is something that October can't do.

A few users have reported having issues with the official integration from time to time so if you'd like to enable uploading store-purchased highlights via October, you can do so via Settings.

!!! warning

    Do note that using the two integrations together may cause more trouble than it's worth.

    October may not send the exact same metadata as the Kobo integration so you might expect duplication of your highlights.

    The ability to sync store-purchased highlights was more of a coincidental bug than an intentional feature but any side effects should be harmless.

![](../assets/settings/settings_storehighlights_light.png#only-light)
![](../assets/settings/settings_storehighlights_dark.png#only-dark)