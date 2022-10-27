#!/bin/bash

# WIP standalone script to manage signing and notarization

# eventually call this with goreleaser as a hook, so that this asset is uploaded?

# TODO: this needs also options=runtime in order to be notarized,
# and we may also want to include an entitlements file
codesign -s 'Developer ID Application: Timothy Sutton (43Y295X5WU)' dist/speedwagon_darwin_all/speedwagon

# zip it (consider if using DMG would somehow make this easier)
# notarytool upload it

# staple the ticket
xcrun stapler staple <artifact>

# not sure if 'install' type makes sense for a zip
spctl --assess -vv --type install <artifact>
