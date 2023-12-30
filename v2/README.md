# v2

This folder contains some experiments with rewriting October. It may or may not amount to anything but I'd like to untangle various elements of the backend.

It isn't intended to be usable in its current state, I'm approximating the upload process by pointing at files on disc instead of actually using my mounted Kobo.

This version will also be attempting to incorporate better parsing by processing epubs and probably parallelised uploading using Goroutines

## Prerequisites

Quickest way to bootstrap a DB for experimenting:

```bash
git clone git@github.com:marcus-crane/kobodbgen
cd kobodbgen
./initdb.sh
```

This will create a mock Kobo DB for usage although I'm using a dump of my Kobo stored on my desktop at present