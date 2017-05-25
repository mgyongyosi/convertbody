package convertbody

import (
	"encoding/json"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

var log = logger.GetLogger("activity-convertbody")

// Activity is a stub for your Activity implementation
type Activity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Activity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error) {

	// Get the activity data from the context
	jsonStr := context.GetInput("input").(string)

	var unmarshalled map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &unmarshalled); err != nil {
		panic(err)
	}

	log.Debugf("Unmarshall done")

	// Set the result as part of the context
	context.SetOutput("output", unmarshalled)

	// Signal to the Flogo engine that the activity is completed
	return true, nil
}
