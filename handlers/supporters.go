package handlers

import (
	"log"
	"net/http"
)

// SupportersHandler handles the supporters page
func (gh *GRCHandler) SupportersHandler(w http.ResponseWriter, r *http.Request) {
	t, err := gh.loadTemplate("/templates/supporters.html")
	if err != nil {
		log.Println("ERROR: could not get supporters template: ", err)
		http.Error(w, err.Error(), 500)
		return
	}

	if err := t.ExecuteTemplate(w, "base", map[string]interface{}{
		"google_analytics_key": googleAnalyticsKey,
	}); err != nil {
		log.Println("ERROR:", err)
	}
}
