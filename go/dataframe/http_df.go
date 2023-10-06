package main

import (
	"net/http"

	"github.com/go-gota/gota/dataframe"
)

func dataHandler(w http.ResponseWriter,
	r *http.Request) {
	df := dataframe.LoadMaps(
		[]map[string]any{
			map[string]any{
				"a": 1, "b": 2,
			}})
	df.WriteJSON(w)
}
func main() {
	http.HandleFunc("/data", dataHandler)
	http.ListenAndServe(":8888", nil)
}
