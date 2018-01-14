package main

import (
	"github.com/kayteh/restokit"
	"github.com/kayteh/skyfish/cmd/skyfish-server/api"
)

func main() {
	r := restokit.NewRestokit("127.0.0.1:5932")
	r.AppName = "skyfish-server"
	r.ShortName = "skyfish"
	api.FetchAPIRoutes(r.Router)
	r.Logger.Info("skyfish-server running on http://127.0.0.1:5932/")
	r.Start()
}
