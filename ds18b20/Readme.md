# RPI Start
sudo modprobe w1-gpio pullup=1
sudo modprobe w1-therm

# Example
pi@pheater ~ $ cat /sys/bus/w1/devices/28-000004abebab/w1_slave 
73 02 4b 46 7f ff 0d 10 06 : crc=06 YES
73 02 4b 46 7f ff 0d 10 06 t=39187
pi@pheater ~ $ cat /sys/bus/w1/devices/w1_bus_master1/w1_master_slaves
28-000004acb8ee
28-000004ac6bf5
28-000004abf273
28-000004abebab


# How to Use

    package main

    import (
        "github.com/kbudde/go-rpio/ds18b20"
        "log"
    )

    func main() {
        s, err := ds18b20.GetSensors()
        if err != nil {
            panic(err)
        }
        for _, sensor := range s {
            value, err := sensor.ReadValue()
            log.Printf("[%v] Value: %v, Error: %v", sensor.Id, value, err)
        }
    }


# Compile

    GOARCH=arm go build