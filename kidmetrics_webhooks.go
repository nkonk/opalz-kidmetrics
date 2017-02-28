package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/influxdata/telegraf"
)

type KidMetricsWebhook struct {
	Path string
	acc  telegraf.Accumulator
}

func (gh *KidMetricsWebhook) Register(router *mux.Router, acc telegraf.Accumulator) {
	router.HandleFunc(gh.Path, gh.eventHandler).Methods("POST")
	log.Printf("I! Started the webhooks_kidmetrics on %s\n", gh.Path)
	gh.acc = acc
}

func (gh *KidMetricsWebhook) eventHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	eventType := r.Header["X-Opalz-Event"][0]
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	e, err := NewEvent(data, eventType)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p := e.NewMetric()
	//TODO:Capture Node Number or more info, if needed, for sharding and other purposes
	gh.acc.AddFields("kidmetrics_node", p.Fields(), p.Tags(), p.Time())
	//TODO: Respond with messages along with status codes 
	w.WriteHeader(http.StatusOK)
}

func generateEvent(data []byte, event Event) (Event, error) {
	err := json.Unmarshal(data, event)
	if err != nil {
		return nil, err
	}
	return event, nil
}

type newEventError struct {
	s string
}

func (e *newEventError) Error() string {
	return e.s
}

func NewEvent(data []byte, name string) (Event, error) {
	log.Printf("D! New %v event received", name)
	switch name {
//	case "kid_zone_transition":
//		return generateEvent(data, &KidZoneTransitionEvent{})
	case "outdoor_telemetry_general":
		return generateEvent(data, &KidGeneralTelemetryEvent{})
/*	case "outdoor_telemetry_sos":
		return generateEvent(data, &KidGeneralTelemetryEvent{})
	case "entered_school":
		return generateEvent(data, &KidEnteredSchoolEvent{})
	case "exited_school":
		return generateEvent(data, &KidExitedSchoolEvent{})
*/
	case "indoor_telemetry_general":
		return generateEvent(data, &KidIndoorTelemetryGeneralEvent{})
	case "indoor_telemetry_sos":
		return generateEvent(data, &KidIndoorTelemetryEmergencyEvent{})

		//TODO: Add more events here as necessary after adding to the Models file.
	}
	return nil, &newEventError{"Not a recognized event type"}
}
