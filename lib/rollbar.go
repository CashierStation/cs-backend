package lib

import (
	"os"

	rollbar "github.com/rollbar/rollbar-go"
)

func InitRollbar() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment("production")                 // defaults to "development"
	rollbar.SetCodeVersion("v2")                         // optional Git hash/branch/tag (required for GitHub integration)
	rollbar.SetServerHost("web.1")                       // optional override; defaults to hostname
	rollbar.SetServerRoot("github.com/heroku/myproject") // path of project (required for GitHub integration and non-project stacktrace collapsing)
}

func CloseRollbar() {
	rollbar.Close()
}
