---
template: overrides/main.html
title: Prerequisites
---

# Prerequisites

Before we get started with the desktop app itself, there are a few things that you'll benefit from knowing before you start uploading books only to have some frustrations later on.

**Cover uploading is disabled by default** as your Kobo will contain highly compressed images by default without some initial setup.

If you'd like to enable it, carry on reading.

## Background

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

## Setting up Calibre (Enabling cover support)

For the unfamiliar, [Calibre](https://calibre-ebook.com/) is basically the go-to tool for eReader management.

If you aren't already using it, I highly recommend it. That said, October should still work fine with regular ePubs uploaded to your Kobo via USB drag-and-drop.

Once you've plugged in your device to your computer via USB, and accepted the connection request on your Kobo, you should open Calibre. It should detect your device but this can often take a bit of time. Anywhere up to 30 seconds is usually reasonable.

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
    If you're reading this documentation in order, we haven't actually installed October just yet. If you're skimming, I highly recommend reading the next section as well.

You can find the covers setting in October under the "Settings" button in the top left once you've selected your Kobo.

![](./assets/october_covers.png)

If you're reading this documentation in order, this step will be mentioned again so you don't forget to turn it.

## Picking the best ebook format for your Kobo

As mentioned before, October should work fine with ePubs but Kobo has its own variation called [Kepub](https://wiki.mobileread.com/wiki/Kepub).

!!! INFO
    Regardless of what format you pick, it's recommended to upload books via Calibre (and fetching metadata) to ensure that your Readwise shelf entries have as much detail as possible.

There is a lot of detail involved in comparing the two but in short, I can personally recommend using Kepubs if you want to enable extra features such as:

* Page turn requests being quicker
* Highlighting text is faster
* Ability to see how many minutes of reading are left for each chapter

For our purposes, highlighting is core to our use case when it comes to Readwise so it's worth adopting but there is a bit of overhead required to get started.

## Setting up Calibre to upload KePubs

TBA