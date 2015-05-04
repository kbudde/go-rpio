package ds18b20

import (
	"testing"
)

func TestGetSensors(t *testing.T) {
	baseDirectory = "./testdata/test1/"
	s, err := GetSensors()
	if len(s) != 4 {
		t.Errorf("GetSensors should return 4 sensors, but it returned %v", len(s))
	}
	if s[0].Id != "28-000004acb8ee" {
		t.Error("First Sensorname not correct")
	}
	if s[1].Id != "28-000004ac6bf5" {
		t.Error("First Sensorname not correct")
	}
	if s[2].Id != "28-000004abf273" {
		t.Error("First Sensorname not correct")
	}
	if s[3].Id != "28-000004abebab" {
		t.Error("First Sensorname not correct")
	}
	if err != nil {
		t.Errorf("GetSensors returned an error but it should not. Error: %v", err)
	}
}

func TestGetSensorsWithInvalidDirectory(t *testing.T) {
	baseDirectory = "./testdata/invalidDir/"
	s, err := GetSensors()
	if s != nil {
		t.Errorf("s should be nil as there is no testdata: %v", s)
	}
	if err == nil {
		t.Errorf("GetSensors returned no error but it should. Error: %v", err)
	}
}

func TestSensor0(t *testing.T) {
	baseDirectory = "./testdata/test1/"
	s, _ := GetSensors()
	sensor := s[0]
	value, err := sensor.ReadValue()
	if err != nil {
		t.Errorf("ReadValue returned an error for the first sensor. Error was unexpected: %v", err)
	}
	if value != 38.312 {
		t.Errorf("Expected temperature to be 38.312°C but returned was: %v", value)
	}
}

func TestSensor1(t *testing.T) {
	baseDirectory = "./testdata/test1/"
	s, _ := GetSensors()
	sensor := s[1]
	value, err := sensor.ReadValue()
	if err == nil {
		t.Errorf("Error was expected. File does not exist")
	}
	if value != DefaultErrorValue {
		t.Errorf("Expected temperature to be default (%v°C) but returned was: %v", DefaultErrorValue, value)
	}
}

func TestSensor2(t *testing.T) {
	baseDirectory = "./testdata/test1/"
	s, _ := GetSensors()
	sensor := s[2]
	value, err := sensor.ReadValue()
	if err == nil {
		t.Errorf("Error was expected. File is invalid")
	}
	if value != DefaultErrorValue {
		t.Errorf("Expected temperature to be default (%v°C) but returned was: %v", DefaultErrorValue, value)
	}
}
