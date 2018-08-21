package readfile

import (

	"bufio"
	"fmt"
	"os"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

// activityLog is the default logger for the Log Activity
var activityLog = logger.GetLogger("activity-flogo-filereader")

const (
	ivMessage   = "filename"
	ivlineNumber = "linenumber"

	ovMessage = "result"
)

func init() {
	activityLog.SetLogLevel(logger.InfoLevel)
}

// LogActivity is an Activity that is used to log a message to the console
// inputs : {message, flowInfo}
// outputs: none
type LogActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &LogActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *LogActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *LogActivity) Eval(context activity.Context) (done bool, err error) {

	//mv := context.GetInput(ivMessage)
	message, ok := context.GetInput(ivMessage).(string)

	if !ok {
    		context.SetOutput("result", "FILENAME_NOT_SET")
    		return true, fmt.Errorf("Filename not set")
  	}
	
	msg := message
	activityLog.Info(msg)

	lnumber, ok := context.GetInput(ivlineNumber).(int)
	if !ok {
  		context.SetOutput("result", "LINE NUMBER NOT SET")
  		return true, fmt.Errorf("line number not set")
   	}


	activityLog.Info(lnumber)

/*
        b, err := ioutil.ReadFile(msg) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'
*/
	fileHandle, _ := os.Open(msg)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	lastLine := 0
	line := ""

	for fileScanner.Scan() {
   		lastLine++

   		if lastLine == lnumber {
     			line = fileScanner.Text()
     			break
   		}
	}

	context.SetOutput("result", line)

	

	return true, nil
}


