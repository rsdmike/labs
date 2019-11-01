# Lab 2 - App Service Configurable - Ingest custom types via REST

Sometimes the need arises to ingest data that doesn't necessarily fit in the Device Services category, but still want the data to be shared w/ other app services that may be listening to core-data. App Service Configurable provides an easy way to send data into core-data without requiring the complexity of a device-service.

In this lab you will learn how to do the following:

- Leverage Consul for updating App Service Configurable Configuration
- Leverage byte[] for custom data type ingestion
- Sharing data with other app services via Core Data

## Context

For this lab, we will assume a retail use case and we will pretend that we have some enterprise data comprised of Stock Keeping Units (SKU) to Universal Product Code (UPC) mappings. A SKU is an identification that represents one or more UPCs and are unique to each retailer. A UPC is universal. That is, the same barcode is printed on all retail packages and represent a particular item. Typically the UPC is what is scanned at a checkout kiosk and if we assume the checkout kiosk sends data to EdgeX, then the value we will be receiving is a UPC -- thus we would need to understand which SKU is mapped to the UPC so we can inform an Enterprise Resource Planning (ERP) System which SKU and UPC has been scanned. Here is an example of a mapping:

    SKU XYZ: - Cool Shoes
        - UPC: 1234 - Black Shoes
        - UPC: 5678 - Red Shoes

We need a way to load this data into an app service and share it among other app services in our edge based micro-service ecosystem.  With this information available in app service, we can take advantage of local compute instead of requiring constant requests to an ERP System.


## Steps

> We will be leveraging the previous app-service-configurable from lab 1.

1. Currently, the binding, or, trigger, for our pipeline is set to "MessageBus" and configured to listen to core data. We are going to change this to be HTTP. This will stop listening to core data and provided an endpoint at `/api/v1/trigger` to `POST` data and initiate the pipeline. 
    1.  Set the `Binding` to `http`. Located here:
    http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Binding/Type/edit
    2. Note that this setting is NOT in `Writable` and thus a restart of the service is required.
    ```
    docker restart edgex-app-service-configurable-lab
    ```      
2. Set `UseTargetTypeOfByteArray` to "true":
    http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/UseTargetTypeOfByteArray/edit
    This will no longer unmarshal the data received from either HTTP POST or messages bus as an EdgeX Event, but instead now as a `[]byte`.

4. Update the parameters for `PushToCore`. Before we can push to core, we need to specify a `DeviceName` and `ReadingName` (aka ValueDescriptorName). 
You can find both of these under the parameters section for `PushToCore` here:
```
http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/Functions/PushToCore/Parameters/
```
You can set these as you wish, example DeviceName can be "mydevicename" and ReadingName can be "myreadingname". Make note what you choose as we will use these to query to the core-data API to ensure it has been sent. 

3. Replace the entire pipeline `ExecutionOrder` with function `PushToCore`:
```
http://ipofmachine:8500/ui/dc1/kv/edgex/appservices/1.0/AppService-mqtt-export/Writable/Pipeline/ExecutionOrder/edit
```

4. Now, let's use the following SKU to UPC mapping to post data into core-data.
```bash
curl -X POST \
  http://127.0.0.1:48095/api/v1/trigger \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: b23cc01a-03e2-48e2-b5f3-093bcf181944' \
  -H 'cache-control: no-cache' \
  -d '[{"sku":"ABC1","upc":["1234","2345"]},{"sku":"DEF2","upc":["3456","4567"]}]'
```

5. If you received no errors, you should now be able to see the data you sent in by querying core-data.
```bash
curl -X GET \
  http://ipofmachine:48080/api/v1/event/device/mydevicename/10
```
You should have retrieved a record similar to:
```json
[
    {
        "id": "2ce01448-c48e-40b4-93c3-a2cf338b4d01",
        "device": "mydevicename",
        "created": 1572564758034,
        "modified": 1572564758034,
        "origin": 1572564758029381870,
        "readings": [
            {
                "id": "1b5f4ce7-f30b-4fea-bac1-4c92485820b3",
                "created": 1572564758032,
                "origin": 1572564758029381870,
                "modified": 1572564758032,
                "device": "mydevicename",
                "name": "myreadingname",
                "value": "[{\"sku\":\"ABC1\",\"upc\":[\"1234\",\"2345\"]},{\"sku\":\"DEF2\",\"upc\":[\"3456\",\"4567\"]}]"
            }
        ]
    }
]
```

You have successfully sent data to core-data via App Service Configurable.



