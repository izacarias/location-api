package locationapi

type UserInfo struct {
	// Address Address of user (e.g. 'sip' URI, 'tel' URI, 'acr' URI) currently on the access point
	// We are currently only using IP addresses
	Address string `json:"address"`

	// AccessPointId The identity of the access point the user is currently on
	AccessPointId *string `json:"AccessPointId,omitempty"`

	// ZoneId The identity of the zone the user is currently within
	ZoneId string `json:"zoneId"`

	// ResourceURL Self-referring URL
	ResourceURL string `json:"resourceURL"`

	// Date and time that location was collected.
	// Timestamp   TimeStamp `json:"timestamp"`

	// The geographical coordinates where the user is
	LocationInfo *LocationInfo `json:"locationInfo,omitempty"`

	// Contextual information of a user location (e.g. aisle, floor, room number, etc.).
	ContextLocationInfo *string `json:"contextLocationInfo,omitempty"`

	// AncillaryInfo Reserved for future use.
	// AncillaryInfo *string `json:"ancillaryInfo,omitempty"`

	// CivicInfo Indicates a Civic address
	// CivicInfo            *CivicAddress         `json:"civicInfo,omitempty"`
	// LocationInfo         *LocationInfo         `json:"locationInfo,omitempty"`
	// RelativeLocationInfo *RelativeLocationInfo `json:"relativeLocationInfo,omitempty"`

}
