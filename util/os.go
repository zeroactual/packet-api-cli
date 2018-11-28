package util

import (
	"sync"
)

type osResp struct {
	Os []os `json:"operating_systems"`
}

type os struct {
	Id   string
	Name string
	Slug string
	Plans []string `json:"provisionable_on"`
}

var oss *osResp

func LoadOs() error {
	return handle("https://api.packet.net/operating-systems", GET, nil, &oss)
}

func GetOssForPlan(planName string) *[]os{
	var matchingPlans []os

	var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(oss.Os))

	for _, o := range oss.Os {
		go func (o os) {
			defer wg.Done()
			for _, plan := range o.Plans {
				if planName == plan {
					mux.Lock()
					matchingPlans = append(matchingPlans, o)
					mux.Unlock()
					break
				}
			}
		} (o)
	}
	wg.Wait()
	return &matchingPlans
}