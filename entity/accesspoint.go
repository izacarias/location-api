package entity

type ConnectionType string
type OperationStatus string

const (
	LTE    ConnectionType = "LTE"
	WIFI   ConnectionType = "WiFi"
	WIMAX  ConnectionType = "WiMAX"
	N5GNR  ConnectionType = "5G NR"
	UNKNOW ConnectionType = "Unknow"
)

const (
	Serviceable   OperationStatus = "Serviceable"
	Unknown       OperationStatus = "Unknown"
	Unserviceable OperationStatus = "Unserviceable"
)

type LocationInfo struct {
	Altitute  float32
	Latitude  float32
	Longitude float32
	Accuracy  int
}

type AccessPoint struct {
	ID              int
	LocationInfo    LocationInfo
	ConnectionType  ConnectionType
	OperationStatus OperationStatus
	NumberOfUsers   int
}
