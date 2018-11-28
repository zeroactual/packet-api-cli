package util

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

type facilitiesObj struct {
	Facilities []facility
}

type facility struct {
	Id   string
	Name string
	Code string
	Plans *[]plan
}

var facilitiesResp facilitiesObj

func LoadFacilities() error {
	return handle(fmt.Sprintf("https://api.packet.net/projects/%s/facilities", viper.Get("project.id")), GET, nil, &facilitiesResp)
}

func GetFacilities() *[]facility {
	// Grab Project ID for later use
	var allFacilities = &facilitiesResp.Facilities
	var facilitiesMapped []facility

	var mux sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(*allFacilities))

	for _, f := range *allFacilities {
		go func (f facility) {
			defer wg.Done()
			plans := GetPlansForFacility(f.Id)
			if len(*plans) > 0 {
				f.Plans = plans
				mux.Lock()
				facilitiesMapped = append(facilitiesMapped, f)
				mux.Unlock()
			}
		} (f)
	}
	wg.Wait()
	return &facilitiesMapped
}