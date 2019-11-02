package main

import (
	"fmt"
	"os"

	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
	"github.com/edgexfoundry/app-functions-sdk-go/appsdk"
)

func main() {

	// 1) First thing to do is to create an instance of the EdgeX SDK, giving it a service key
	edgexSdk := &appsdk.AppFunctionsSDK{
		ServiceKey: "Lab3Exercise", // Key used by Registry (Aka Consul)
	}

	// 2) Next, we need to initialize the SDK
	if err := edgexSdk.Initialize(); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK initialization failed: %v\n", err))
		os.Exit(-1)
	}

	// 3) Since our FilterByDeviceName Function requires the list of Device Names we would
	// like to search for, we'll go ahead and define that now.
	// TODO: define list of filters

	// 4) This is our pipeline configuration, the collection of functions to
	// execute every time an event is triggered.
	if err := edgexSdk.SetFunctionsPipeline(
	// TODO: Define Function Pipeline
	); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK SetPipeline failed: %v\n", err))
		os.Exit(-1)
	}

	// 5) Lastly, we'll go ahead and tell the SDK to "start" and begin listening for events to trigger the pipeline.
	edgexSdk.MakeItRun()
}

//TODO: Define custom function for transformation
func Fahrenheit(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	// TODO: Check that there is a result to work with

	// TODO: Convert the value

	// TODO: Print out the value for debugging purposes using the logging on the context

	// TODO: Publish the result to the ZMQ Topic (Step 4)

	// TODO: Return appropriate values back to the pipeline

	return true, nil
}
