package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Roy19/distributed-transaction-2pc/store-svc/controllers"
	"github.com/Roy19/distributed-transaction-2pc/store-svc/db"
	"github.com/Roy19/distributed-transaction-2pc/utils"
	"github.com/go-chi/chi/v5"
)

func initRoutes(mux *chi.Mux, controller *controllers.StoreController) {
	mux.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.Get("/store/item/{itemID}", func(w http.ResponseWriter, r *http.Request) {
		itemID := chi.URLParam(r, "itemID")
		itemIDAsInt, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			errorMessage := map[string]any{
				"error": "itemID is required",
			}
			utils.Respond(w, http.StatusBadRequest, errorMessage)
			return
		}
		controller.GetItem(itemIDAsInt)
	})
	mux.Post("/store/item/{itemID}/reserve", func(w http.ResponseWriter, r *http.Request) {
		itemID := chi.URLParam(r, "itemID")
		itemIDAsInt, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			errorMessage := map[string]any{
				"error": "itemID is required",
			}
			utils.Respond(w, http.StatusBadRequest, errorMessage)
			return
		}
		controller.ReserveItem(itemIDAsInt)
	})
	mux.Post("/store/item/{itemID}/book", func(w http.ResponseWriter, r *http.Request) {
		itemID := chi.URLParam(r, "itemID")
		itemIDAsInt, err := strconv.ParseInt(itemID, 10, 64)
		if err != nil {
			errorMessage := map[string]any{
				"error": "itemID is required",
			}
			utils.Respond(w, http.StatusBadRequest, errorMessage)
			return
		}
		controller.BookItem(itemIDAsInt)
	})
}

func initDependencies() *controllers.StoreController {
	db.InitDB()
	db.PutDummyData()
	return &controllers.StoreController{}
}

func main() {
	mux := chi.NewRouter()
	controller := initDependencies()
	initRoutes(mux, controller)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("failed to start server")
	}
}
