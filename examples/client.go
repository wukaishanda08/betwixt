package main

import (
    . "github.com/zubairhamed/lwm2m"
)

func main() {
    client := NewLWM2MClient(":0", "localhost:5683")

    registry := NewDefaultObjectRegistry()
    client.UseRegistry(registry)

    setupResources(client, registry)

    client.OnStartup(func(){
        client.Register("GOCLIENT")
    })

    // client.OnRead(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {
    client.OnRead(func() {
        // log.Println(evt.Data["objectModel"].(*ObjectModel))
    })

    /*
    client.OnRegistered(func(evt *Event, path string){
        log.Println("Client is registered")
    })

    client.OnUnregistered(func(evt *Event){
        log.Println("Client is Unregistered")
    })

    client.OnExecute(func(evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {

    })

    client.OnWrite(func(evt *Event, m *ObjectInstance, i *ResourceInstance, value interface{}) (*LWM2MResponse) {

    })

    client.OnCreate(func (evt *Event, m *ObjectInstance, i *ResourceInstance) (*LWM2MResponse) {

    })
    */

    client.Start()
}


func setupResources (client *LWM2MClient, reg *ObjectRegistry) {

    client.EnableObject(OBJECT_LWM2M_SECURITY)
    client.EnableObject(OBJECT_LWM2M_SERVER)
    client.EnableObject(OBJECT_LWM2M_ACCESS_CONTROL)
    client.EnableObject(OBJECT_LWM2M_DEVICE)
    client.EnableObject(OBJECT_LWM2M_CONNECTIVITY_MONITORING)
    client.EnableObject(OBJECT_LWM2M_FIRMWARE_UPDATE)
    client.EnableObject(OBJECT_LWM2M_LOCATION)
    client.EnableObject(OBJECT_LWM2M_CONNECTIVITY_STATISTICS)

    instanceSec1 := reg.CreateObjectInstance(OBJECT_LWM2M_SECURITY, 0)
    instanceSec2 := reg.CreateObjectInstance(OBJECT_LWM2M_SECURITY, 1)
    instanceSec3 := reg.CreateObjectInstance(OBJECT_LWM2M_SECURITY, 2)
    instanceServer := reg.CreateObjectInstance(OBJECT_LWM2M_SERVER, 1)

    instanceAccessCtrl1 := reg.CreateObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 0)
    instanceAccessCtrl2 := reg.CreateObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 1)
    instanceAccessCtrl3 := reg.CreateObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 2)
    instanceAccessCtrl4 := reg.CreateObjectInstance(OBJECT_LWM2M_ACCESS_CONTROL, 3)
    instanceDevice := reg.CreateObjectInstance(OBJECT_LWM2M_DEVICE, 0)
    instanceConnMonitoring := reg.CreateObjectInstance(OBJECT_LWM2M_CONNECTIVITY_MONITORING, 0)
    instanceFwUpdate :=  reg.CreateObjectInstance(OBJECT_LWM2M_FIRMWARE_UPDATE, 0)

    client.AddObjectInstances(
        instanceSec1, instanceSec2, instanceSec3,
        instanceServer,
        instanceAccessCtrl1, instanceAccessCtrl2, instanceAccessCtrl3, instanceAccessCtrl4,
        instanceDevice,
        instanceConnMonitoring,
        instanceFwUpdate,
    )

/*
    client
        ObjectModel
            ObjectInstances
                Resources

    client
        LWM2MObject


    client
        [int] --> Object

    ---
    device := NewDeviceObject()
    device ...

    client.enableObject(OBJECT_LWM2M_DEVICE)


    client enables objects
    client adds instances
    client sets resource values

    ---

    LWM2MObject
    LWM2MResource
    ObjectInstance
    ResourceValue

    map[string] LWM2MObject

    map[LWM2MObjectType] LWM2MObject
                            LWM2MInstance
                                LWM2MResource
*/
}