package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"


	"github.com/edgexfoundry/app-functions-sdk-go/appcontext"
	"github.com/edgexfoundry/app-functions-sdk-go/appsdk"
	"github.com/edgexfoundry/app-functions-sdk-go/pkg/transforms"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
)

func main() {

	// 1) First thing to do is to create an instance of the EdgeX SDK, giving it a service key
	edgexSdk := &appsdk.AppFunctionsSDK{
		ServiceKey: "Lab4Exercise", // Key used by Registry (Aka Consul)
	}

	// 2) Next, we need to initialize the SDK
	if err := edgexSdk.Initialize(); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK initialization failed: %v\n", err))
		os.Exit(-1)
	}

	// 3) Since our FilterByDeviceName Function requires the list of Device Names we would
	// like to search for, we'll go ahead and define that now.
	deviceNames := []string{"Random-Integer-Device"}

	// 4) This is our pipeline configuration, the collection of functions to
	// execute every time an event is triggered.
	if err := edgexSdk.SetFunctionsPipeline(
		transforms.NewFilter(deviceNames).FilterByDeviceName,
		Fahrenheit,
		// TODO: add command function
	); err != nil {
		edgexSdk.LoggingClient.Error(fmt.Sprintf("SDK SetPipeline failed: %v\n", err))
		os.Exit(-1)
	}

	// 6) Lastly, we'll go ahead and tell the SDK to "start" and begin listening for events to trigger the pipeline.
	edgexSdk.MakeItRun()
}

func Fahrenheit(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	// TODO: Check that there is a result to work with
	if len(params) < 1 {
		// We didn't receive a result
		return false, errors.New("No Data Received")
	}
	event := params[0].(models.Event)
	temperatureInCelsius,_ := strconv.Atoi(event.Readings[0].Value)
	// TODO: Convert the value
	result := (temperatureInCelsius*9/5 + 32)
	// TODO: Print out the value for debugging purposes using the logging on the context
	edgexcontext.LoggingClient.Info(strconv.Itoa(result))
	// TODO: Publish the result to the ZMQ Topic (Step 4)
	edgexcontext.Complete([]byte(strconv.Itoa(result)))

	return true, nil
}

func ThresholdCheck(edgexcontext *appcontext.Context, params ...interface{}) (bool, interface{}) {
	// TODO: Check that there is a result to work with

	// TODO: Determine if value falls outside of threshold parameters

	// TODO: Send command using context to device-service to turn it off

	// TODO: Return appropriate values back to the pipeline
	return true, nil
}
