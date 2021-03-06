package tide

import (
	"net/http"

	"github.com/jenkins-x/lighthouse/pkg/tide/history"
)

// Controller the interface for all tide controllers
// whether regular or the GitHub App flavour which has to handle tokens differently
type Controller interface {
	Sync() error
	Shutdown()
	GetPools() []Pool
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	GetHistory() *history.History
}
