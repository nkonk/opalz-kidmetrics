package github

import (
//	"fmt"
	"log"
	"time"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/metric"
)

const meas = "kidmetrics_webhooks"

type Event interface {
	NewMetric() telegraf.Metric
}
/* TODO: CLub Together all the commoon fields for neater embedding.
type KidCommon struct {
	KID int64 `json:kid_id`
}
*/
type KidGeneralTelemetryEvent struct {
	MetricTStamp time.Time `json:"evt_tstamp"`
	KID					string		`json:"kid_uid"`
	RecLat		 float32	`json:"rec_lat"`
	RecLong		 float32	`json:"rec_long"`
	RecDelta	 float32	`json:"rec_delta"`
	DeviceId	 string		`json:"device_id"`
	//PolyId		 float32	`json:"poly_id"`
	AuxPayload	 string		`json:"payload_data"`
}


func (s KidGeneralTelemetryEvent) NewMetric() telegraf.Metric {
	event := ""
	t := map[string]string{
		"event" 	:		event,
		"uid"		:		s.KID,
		"deviceid"	:		s.DeviceId,
	}
	f := map[string]interface{}{
		"ev_time"		 :   s.MetricTStamp,
		"rec_lat"		 :   s.RecLat,
		"rec_long"		 :   s.RecLong,
		"rec_delta"		 :   s.RecDelta,
	//  "loc_poly_id"  :   s.Deployment.Task, // Polygon Id of the location that has been captured 
		"payload_aux"    :	 s.AuxPayload, // Time Since event for 
	}
	m, err := metric.New(meas, t, f, time.Now())
	if err != nil {
		log.Fatalf("Failed to create %v event", event)
	}
	return m
}



/*type KidEnteredSchoolEvent struct {
	
}


func (s KidEnteredSchoolEvent) NewMetric() telegraf.Metric {
	event := ""
	t := map[string]string{
		"event" 	:		event,
		"uid"		:		s.Repository.Repository,
		"deviceid"	:		fmt.Sprintf("%v", s.Repository.Private),
	}
	f := map[string]interface{}{
		"ev_time"		 :   s.metricTStamp,
		"rec_lat"		 :   s.RecLat,
		"rec_long"		 :   s.RecLong,
		"rec_delta"		 :   s.RecDelta,
	//  "loc_poly_id"  :   s.Deployment.Task, // Polygon Id of the location that has been captured 
		"payload_aux"    :	 s.AuxPayload, // Time Since event for 
	}
	m, err := metric.New(meas, t, f, time.Now())
	if err != nil {
		log.Fatalf("Failed to create %v event", event)
	}
	return m
}

*/

type KidIndoorTelemetryGeneralEvent struct {
	MetricTStamp time.Time `json:"evt_tstamp"`
	KID			 string		`json:"kid_uid"`
	VCell		 int64		`json:"loc_VCell"`
	HCell		 int64		`json:"loc_HCell"`
	AuxPayload	 string		`json:"payload_data"`
	DeviceId	 string		`json:"device_id"`
}

func (s KidIndoorTelemetryGeneralEvent) NewMetric() telegraf.Metric  {
	event := "indoor_telemetry_heartbeat"
	t := map[string]string{
	 	"event" 	:		event,
		"uid"		:		s.KID,
		"deviceid"	:		s.DeviceId,
	}
	f := map[string]interface{}{
		"ev_time"		 :   s.MetricTStamp,
		"rec_Vcell"		 :   s.VCell,
		"rec_Hcell"		 :   s.HCell,
	//  "loc_poly_id"  :     s.PolyId, // Polygon Id of the location that has been captured 
		"payload_aux"    :	 s.AuxPayload, // Time Since event for 
	}
	m, err := metric.New(meas, t, f, time.Now())
	if err != nil {
		log.Fatalf("Failed to create %v event", event)
	}
	return m

}


type KidIndoorTelemetryEmergencyEvent struct {
    MetricTStamp	    time.Time `json:"evt_tstamp"`
	KID					string		`json:"kid_uid"`
	VCell				int64		`json:"loc_VCell"`
	HCell				int64		`json:"loc_HCell"`
	EmergencyPayload	string		`json:"sos_payload"`
	AuxPayload			string		`json:"payload_data"`
	DeviceId			string		`json:"device_id"`
}

func (s KidIndoorTelemetryEmergencyEvent) NewMetric() telegraf.Metric  {
	event := "indoor_telemetry_heartbeat"
	t := map[string]string{
		"event" 	:		event,
		"uid"		:		s.KID,
		"deviceid"	:		s.DeviceId,
	}
	f := map[string]interface{}{
		"ev_time"		 :   s.MetricTStamp,
		"rec_Vcell"		 :   s.VCell,
		"rec_Hcell"		 :   s.HCell,
	//  "loc_poly_id"  :     s.PolyId, // Polygon Id of the location that has been captured 
		"payload_aux"    :	 s.AuxPayload, // Time Since event for 
		"sos_payload"	 :	 s.EmergencyPayload,
	}
	m, err := metric.New(meas, t, f, time.Now())
	if err != nil {
		log.Fatalf("Failed to create %v event", event)
	}
	return m

}





