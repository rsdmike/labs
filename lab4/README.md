# Lab 4 - App Functions SDK - Commanding a Device Service

This lab will build on lab 3. Now that we have our Int64 value filtered and converted to Fahrenheit, let's add a new function that makes a command against the virtual device service. 

In this lab you will learn how to do the following:

- Command a device service using the App Functions SDK context
- Understand how the SDK passes values from one function to the next

    
## Context

For this lab we will inspect the converted temperature value, determine if it falls into a certain temperature range (example: 25-75), and if it falls outside of this range, we will send a command to turn of the device. 

## Steps

1. Open up `main.go`. This file picks up where lab3 should be when completed. You should be seeing values printed to the console. 
2. Next, let's update the `ThresholdCheck` function with your threshold check. Feel free to adjust the threshold values as you like. They don't have to match whats listed above in the Context section. If it falls outside of your tolerances, lets send a command to the device-service to shut it down and stop sending values. 

To do this, you can leverage the `CommandClient` provided on the `edgexcontext` (the first parameter in the function signature). Check out the following for an idea how to use the `CommandClient`: 
`https://github.com/edgexfoundry/go-mod-core-contracts/blob/master/clients/command/client.go`

> Note: It's important that the CommandService is defined in configuration.toml under the `[Services]` section. This is already done for you in Lab4.

3. That's it! You've now sent a command to the device service to perform an action based on some logic performed in the app service. While this example is trivial, hopefully you get a sense of how to use the SDK to perform various actions in the EdgeX framework.
> Note: In our testing, it doesn't appear to actually turn off the values -- but the command is successfully sent. 

