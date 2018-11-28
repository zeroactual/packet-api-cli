package util

import (
	"fmt"
	"strings"
	"sync"

	"github.com/spf13/viper"
)

type plans struct {
	Plans []plan
}

type plan struct {
	Id   string
	Name string
	Facilities []struct{Href string} `json:"available_in"`
	Os *[]os
}

var plansResp *plans

func LoadPlans() error {
	return handle(fmt.Sprintf("https://api.packet.net/projects/%s/plans", viper.Get("project.id")), GET, nil, &plansResp)
}

func GetPlansForFacility(id string) *[]plan {
	var matchingPlans []plan
	var replacer = strings.NewReplacer("/facilities/", "")

	var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(plansResp.Plans))

	for _, p := range plansResp.Plans {
		go func(p plan) {
			defer wg.Done()
			for _, facility := range p.Facilities {
				var facilityId = replacer.Replace(facility.Href)
				if id == facilityId {
					p.Os = GetOssForPlan(p.Name)
					mux.Lock()
					matchingPlans = append(matchingPlans, p)
					mux.Unlock()
					break
				}
			}
		}(p)
	}

	wg.Wait()

	return &matchingPlans
}