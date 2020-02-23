// Package getactivespacestable gets Table data from ActiveSpaces
package getactivespacestable

import (

	"tibco.com/tibdg"
	"fmt"
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"tibco.com/tibdg"
)

// Constants used by the code to represent the input and outputs of the JSON structure
const (
	ivconnectionURL = "connectionURL"
	ivtableName     = "tableName"
	ivkey           = "key"
	ovResult        = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-getactivespacestable")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// ExpressionAttribute is a structure representing the JSON payload for the expression syntax
type ExpressionAttribute struct {
	Name  string
	Value string
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// Get the inputs
	connectionURL := context.GetInput(ivconnectionURL).(string)
	tableName := context.GetInput(ivtableName).(string)
	key := context.GetInput(ivkey).(string)

	connection, err := tibdg.NewConnection(connectionURL, "", nil)
	if err != nil {
		log.Fatal(err)
	}

	session, err := connection.NewSession(nil)
	if err != nil {
		log.Fatal(err)
	}

	table, err := session.OpenTable(tableName, nil)
	if err != nil {
		log.Fatal(err)
	}

	keyRow, err := table.NewRow()
	if err != nil {
		log.Fatal(err)
	}

	content := tibdg.RowContent{"key": key}
	err = keyRow.Set(content)
	getRow, err := table.Get(keyRow)
	if err != nil {
		log.Fatal(err)
	}

	keyRow.Destroy()
	if getRow != nil {
		getRow.Destroy()
	}

	context.SetOutput(ovResult, string(getRow))
	return true, nil

}
