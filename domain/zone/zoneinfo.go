package zone

import locationapi "github.com/izacarias/location-api"

type ZoneInfo struct {
	// zone is the root entity
	zone *locationapi.Zone
}

func NewZoneInfo(zoneid string) (ZoneInfo, error) {
	// zone is the root entity
	zone := locationapi.NewZone(zoneid)
	return ZoneInfo{zone: &zone}, nil
}
