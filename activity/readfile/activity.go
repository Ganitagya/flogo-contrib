package readfile

import (
	"fmt"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"

	"io/ioutil"
)

// log is the default package logger
var log = logger.GetLogger("Activity Akash-File Reader")

const (
	filename   = "filename"
	lineNumber = "lineNumber"

	ovresult = "result"
)

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error) {

	// do eval
	ivfilename, ok := context.GetInput(filename).(string)
	if !ok {
		context.SetOutput("result", "FILENAME_NOT_SET")
		return true, fmt.Errorf("Filename not set")
	}

	b, err := ioutil.ReadFile(ivfilename) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	

	context.SetOutput("result", str)

	return true, nil
}
