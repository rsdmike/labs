# Lab 1 - Configuring App Service Configurable

This lab will walk you through configuring an app-service-configurable instance. It is the quickest way to get started with exporting data from EdgeX.

In this lab you will learn how to do the following:

- Use the existing deployed App Service Configurable instance
- Update the configuration in real-time via Consul
- Export data w/ the App Service configurable
    
## Context

For this lab we will consume data from EdgeX, perform some basic filtering, and then leverage the built in MQTT function to send data to a local MQTT broker.  

## Steps

1. run `docker ps --format "table {{.ID}}\t{{.Names}}\t{{.Ports}}"` and you should see the following services running

    |CONTAINER ID    |NAMES                            |PORTS                      |
    |----------------|---------------------------------|---------------------------|
    | 000000000000   | mqtt                            |  0.0.0.0:1883->1883/tcp
    | 000000000000   |     edgex-device-virtual        |  0.0.0.0:49990->49990/tcp |
    | 000000000000   |     edgex-ui-go                 |  0.0.0.0:4000->4000/tcp |
    | 000000000000   |     edgex-export-distro         |  0.0.0.0:5566->5566/tcp,  0.0.0.0:48070->48070/tcp |
    | 000000000000   |     edgex-export-client         |  0.0.0.0:48071->48071/tcp |
    | 000000000000   |     edgex-support-scheduler     |  0.0.0.0:48085->48085/tcp |
    | 000000000000   | edgex-app-service-configurable-lab | 48095/tcp, 0.0.0.0:48100->48100/tcp |
    | 000000000000   |     edgex-core-command          |  0.0.0.0:48082->48082/tcp |
    | 000000000000   |     edgex-core-data             |  0.0.0.0:5563->5563/tcp,  0.0.0.0:48080->48080/tcp |
    | 000000000000   |     edgex-core-metadata         |  0.0.0.0:48081->48081/tcp,  |48082/tcp
    | 000000000000   |     edgex-sys-mgmt-agent        |  0.0.0.0:48090->48090/tcp |
    | 000000000000   |     edgex-support-notifications |  0.0.0.0:48060->48060/tcp |
    | 000000000000   |     edgex-support-logging       |  0.0.0.0:48061->48061/tcp |
    | 000000000000   |     edgex-core-consul           |  0.0.0.0:8400->8400/tcp,  8300-8302/tcp, 8301-8302/udp, 8600/tcp, 8600/udp, 0.0.0.0:8500->8500/tcp |
    | 000000000000   |     edgex-mongo                 |  0.0.0.0:27017->27017/tcp |

    > Note that `edgex-core-consul` is accessible via a web-browser on port 8500 and that the there is an app-service-configurable instance running.

2. Let's take a look at the logs of app-service-configurable with the following command:
    ```shell
    docker logs docker logs edgex-app-service-configurable-lab
    ```
    We should see something akin to:
    ```
    level=INFO ts=2019-10-31T21:07:36.637614117Z app=AppService-mqtt-export source=sdk.go:492 msg="MessageBus trigger selected"
    level=INFO ts=2019-10-31T21:07:36.637649211Z app=AppService-mqtt-export source=messaging.go:46 msg="Initializing Message Bus Trigger. Subscribing to topic: events on port 5563 , Publish Topic:  on port 0"
    level=INFO ts=2019-10-31T21:07:36.63805047Z app=AppService-mqtt-export source=sdk.go:646 msg="Listening for changes from registry"
    level=INFO ts=2019-10-31T21:07:36.638539997Z app=AppService-mqtt-export source=sdk.go:153 msg="StoreAndForward disabled. Not running retry loop."
    level=INFO ts=2019-10-31T21:07:36.638969667Z app=AppService-mqtt-export source=telemetry.go:79 msg="Starting CPU Usage Average loop"
    level=INFO ts=2019-10-31T21:07:36.638803628Z app=AppService-mqtt-export source=sdk.go:156 msg="AppServiceConfigurable-mqtt-export has Started"
    level=INFO ts=2019-10-31T21:07:36.647000493Z app=AppService-mqtt-export source=server.go:232 msg="Starting HTTP Server on port :48100"
    level=INFO ts=2019-10-31T21:07:36.650305497Z app=AppService-mqtt-export source=sdk.go:674 msg="Writable configuration has been updated from Registry"
    level=INFO ts=2019-10-31T21:07:36.650569007Z app=AppService-mqtt-export source=sdk.go:743 msg="Reloaded Configurable Pipeline from Registry"
    ```

3. Now that we know App Service Configurable is up and running. Visit Consul at
http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/

    This represents the entire current configuration for the app service configurable instance. The `Writable` section (should be last in the list) is where we will make our changes. Any changes in the section are real-time and do not require a restart of the instance. For the sake of this lab, all settings are already set to send data to a local MQTT Broker. You can inspect these settings by checking out the values at the following location:

    http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/Functions/MQTTSend/Addressable/



4. You can subscribe to the local mqtt broker using this command to view messages currently being sent to the mqtt broker from the app-service-configurable instance:
    ```shell
    docker run --init -it --rm --net=host efrecon/mqtt-client sub -h localhost -t "#" -v
    ```
    After a moment, you should start seeing lots of data flow through to the mqtt broker that has been generated for by the device-virtual service. 

5. Our next step is to add a filter to the function pipeline to narrow down what is being sent to the MQTT Broker. Visit the following URL to configure the parameters for our `FilterByDeviceName` function. 
    ```
    http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/Functions/FilterByDeviceName/Parameters/DeviceNames/edit
    ```
    Choose a device name that you've observed from the MQTT output. Example: `Random-Integer-Device`.
    Be sure and click "Save" when you've added a device name. You can add multiple device names as this is a comma-delimited list.

6. Check the MQTT broker again (command from Step 4). Notice that the filter isn't working quite yet. This is because we must add the `FilterByDeviceName` to our functions pipeline. This is specified by the `ExecutionOrder` key in consul.
    ```
    http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/ExecutionOrder/edit
    ```
    Currently it is set to: `TransformToJSON, MQTTSend, MarkAsPushed`. Lets add our filter:
    `FilterByDeviceName,TransformToJSON, MQTTSend, MarkAsPushed`. Now, if we check the MQTT broker again, we should now only see data from the devices specified in the in step 5.

7. Go ahead and play around with different device names in step 5 again while you monitor the MQTT broker. You should see the filter instantly applied as you save your changes. Or if you're feeling adventurous, try using a different SDK Function such as `FilterByValueDescriptor`. 





