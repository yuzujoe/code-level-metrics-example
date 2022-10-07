package main

import (
	"github.com/gorilla/mux"
	"github.com/newrelic/go-agent/v3/integrations/nrgorilla"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"os"
	"time"
)

func makeNewRelicApplication() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("code-level-metrics-example"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigAppLogEnabled(false),
		newrelic.ConfigCodeLevelMetricsEnabled(true),
	)
	if err != nil {
		return nil, err
	}

	if err = app.WaitForConnection(5 * time.Second); err != nil {
		return nil, err
	}

	return app, nil
}

func exampleHandler(w http.ResponseWriter, r *http.Request) {
	txn := newrelic.FromContext(r.Context())
	defer txn.StartSegment("ExampleHandler").End()

	w.WriteHeader(http.StatusOK)
}

func main() {
	r := mux.NewRouter()

	app, err := makeNewRelicApplication()
	if err != nil {
		os.Exit(1)
	}

	r.Use(nrgorilla.Middleware(app))

	r.HandleFunc(newrelic.WrapHandleFunc(app, "/example", exampleHandler))

	http.ListenAndServe(":8000", r)
}
