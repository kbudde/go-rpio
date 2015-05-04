package ds18b20

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

const (
	DefaultErrorValue = -999
)

var (
	baseDirectory = "/sys/bus/w1/devices/"
)

type Sensor struct {
	Id string
}

//returns an array of Sensor. Nil if an error occured.
func GetSensors() ([]Sensor, error) {
	b, err := ioutil.ReadFile(baseDirectory + "w1_bus_master1/w1_master_slaves")
	if err != nil {
		return nil, err
	}
	ids := strings.Split(string(b), "\n")
	sensors := make([]Sensor, len(ids))
	for i, s := range ids {
		sensors[i] = Sensor{s}
	}
	return sensors, nil
}

func (s *Sensor) ReadValue() (float64, error) {
	regex_temperature := regexp.MustCompile(`t=(\d+)`)

	b, err := ioutil.ReadFile(baseDirectory + s.Id + "/w1_slave")
	if err != nil {
		return DefaultErrorValue, err
	}
	lines := strings.Split(string(b), "\n")
	if len(lines) != 2 {
		return DefaultErrorValue, errors.New("Sensor file: Unexpected size.")
	}
	valid, _ := regexp.MatchString(` YES$`, lines[0])
	if !valid {
		return DefaultErrorValue, errors.New("Unsuccessfull read. Try again")
	}
	match := regex_temperature.FindStringSubmatch(lines[1])
	if match == nil || len(match) != 2 {
		return DefaultErrorValue, errors.New("Sensor file: Unexpected format")
	}
	temperatureFloat, _ := strconv.ParseFloat(match[1], 32)
	return temperatureFloat / 1000, nil
}
