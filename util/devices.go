package util

import (
	"fmt"

	"github.com/spf13/viper"
)

type deviceReq struct {
	FacilityId string `json:"facility"`
	PlanId string `json:"plan"`
	OsId string `json:"operating_system"`
	Hostname string `json:"hostname"`
}

type deviceResp struct {
	Devices []device
}

type device struct {
	Id string
	Hostname string
	Facility facility
	Os os `json:"operating_system"`
	Plan plan
}

func CreateDevice(facilityId string, planId string, osId string, hostname string) error {
	var d = &deviceReq{
		facilityId,
		planId,
		osId,
		hostname,
	}

	return handle(fmt.Sprintf("https://api.packet.net/projects/%s/devices", viper.Get("project.id")), POST, d, nil)
}

func DeleteDevice(deviceId string) error {
	return handle(fmt.Sprintf("https://api.packet.net/devices/%s", deviceId), DELETE, nil, nil)
}

func GetDevices() (*[]device, error) {
	var devices deviceResp
	err := handle(fmt.Sprintf("https://api.packet.net/projects/%s/devices", viper.Get("project.id")), GET, nil, &devices)
	for i, d := range devices.Devices {
		fmt.Printf("[%d] ID: %s, Type: %s, Hostname: %s\n", i, d.Id, d.Plan.Name, d.Hostname)
	}
	return &devices.Devices, err
}