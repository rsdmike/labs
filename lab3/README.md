# Lab 3 - Using the App Functions SDK

This lab will walk you through leveraging the App Functions SDK to build your own custom functions. 

In this lab you will learn how to do the following:
    
- Leverage Built-In Functions (ex, Filtering)
- Create your Function
- Publish data to a new ZMQ Topic

## Context

For this lab we will pretend that the values coming in as Int32 from the virtual device is temperature data and will assume it is being published in celsius. Our goal is to filter out other values from the device service, convert celsius to fahrenheit, and publish to a new ZMQ topic. 

## Steps
