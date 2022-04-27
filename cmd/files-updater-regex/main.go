package main

import (
	regexUpdater "github.com/a5r0n/files-updater-regex/pkg/updater"
	"github.com/go-semantic-release/semantic-release/v2/pkg/plugin"
	"github.com/go-semantic-release/semantic-release/v2/pkg/updater"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		FilesUpdater: func() updater.FilesUpdater {
			return &regexUpdater.Updater{}
		},
	})
}
