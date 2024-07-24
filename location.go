package locationapi

type LocationInfo struct {
	// Latitude Location latitude, expressed in the range -90° to +90°. Cardinality greater than one only if "shape" equals 7.
	Latitude []float32 `json:"latitude"`

	// Longitude Location longitude, expressed in the range -180° to +180°. Cardinality greater than one only if "shape" equals 7.
	Longitude []float32 `json:"longitude"`

	// Altitude Location altitude relative to the WGS84 ellipsoid surface.
	Altitude *float32 `json:"altitude,omitempty"`

	// Accuracy Horizontal accuracy/(semi-major) uncertainty of location provided in meters, as defined in [14]. Present only if "shape" equals 4, 5 or 6.
	Accuracy *int `json:"accuracy,omitempty"`
}
