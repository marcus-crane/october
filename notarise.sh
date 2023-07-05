#!/usr/bin/env bash

codesign -s "Developer ID Application: Marcus Crane" -f -v --timestamp --options runtime October.app
/usr/bin/ditto -c -k --keepParent October.app October.zip
xcrun notarytool submit --wait --apple-id marcus@utf9k.net --team-id <TEAMID> --password <PASSWORD> October.zip
xcrun stapler staple October.app
