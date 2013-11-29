package lock

import (
	"path"
	"net/http"

	"github.com/gorilla/mux"
)

// releaseLockHandler deletes the lock.
func (h *handler) releaseLockHandler(w http.ResponseWriter, req *http.Request) {
	h.client.SyncCluster()

	vars := mux.Vars(req)
	keypath := path.Join(prefix, vars["key_with_index"])

	// Delete the lock.
	_, err := h.client.Delete(keypath)
	if err != nil {
		http.Error(w, "delete lock index error: " + err.Error(), http.StatusInternalServerError)
		return
	}
}

