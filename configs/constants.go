package configs

const (
	UserRoleAdmin      = "admin"
	UserStatusActive   = "active"
	UserApprovedStatus = "approved"

	CompanyStatusActive             = "active"
	CompanyArchivedStatusUnarchived = "unarchived"

	VehicleStatusSuccess    = "success"
	VehicleRegistered       = "registered"
	VehicleUnRegistered     = "unregistered"
	VehicleStatusQuarantine = "quarantine"
	VehicleStatusNew        = "new"
	UploadStatusFail        = "fail"
	UploadBatchesDone       = "done"
	UploadBatchesProcess    = "in-process"
	RecordExist             = "existed"

	DefaultDateCsvLayoutFormat = "02/01/2006"
	MilesKmRatio               = 1.60934

	TypeVehicle   = "vehicle"
	TypeVehicleCD = "vehicle_cd"

	ErrorRegistrationNumber = "invalid registration number"
	ErrorVehicleFile        = "vehicle must have 5 columns"
	ErrorCDFile             = "c&d must have 19 columns"
	LenOfVehicleCsv         = 5
	LenOfCDCsv              = 19

	MaximumFileSizeUpload = 5 * 1024 * 1024

	CompanyOrderDefault = "updated_at DESC"
	VehicleOrderDefault = "updated_at DESC"
)

const (
	RespCodeSuccess = 200
	RespCodeCommon  = 502
)

func VehicleTypes() (rs []string) {
	rs = []string{
		"Cargo",
		"Bike",
		"Car",
		"Lorry",
		"Motorbike",
		"Van",
		"Walker",
		"Pedal Cycle",
		"Sit on Scooter",
		"Sit on eScooter",
		"Pedal Assisted Cargo Bike",
		"Solar Powered Vehicle",
		"Car Derived Van",
		"Micro Van",
		"Small Van",
		"Medium Van",
		"Large Van",
		"3.5 ft Luton Can",
		"7.5ft Truck",
		"12t Truck",
		"18t Truck",
		"26t Truck",
	}
	return
}
