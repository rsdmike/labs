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

We need a way to load this data into an app service and share it among other app services in our edge based micro-service ecosystem. 

With this context, a device-service doesn't really align with our needs, and simply wish to 

## Steps

1) Update trigger to http
2) Set byte array true
3) update pipeline to push to core
4) Now view data in database -- or rest call to core




