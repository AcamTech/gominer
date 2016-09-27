package adl

/*
// XXX all the C implementations use dlopen()
#cgo linux CFLAGS: -DLINUX
#cgo linux LDFLAGS: -latiadlxx -ldl
#include <stddef.h>
#include <stdbool.h>
#include <adl_sdk.h>
int getADLFanPercent(int deviceid);
int getADLTemp(int deviceid);
int setADLFanAutoManage(int deviceid);
int setADLFanPercent(int deviceid, int fanPercent);
*/
import "C"

// DeviceFanGetPercent fetches and returns fan utilization for a device index
func DeviceFanGetPercent(index int) uint32 {
	fanPercent := uint32(0)

	fanPercent = uint32(C.getADLFanPercent(C.int(index)))

	return fanPercent
}

// DeviceFanSetPercent sets the fan to a percent value for a device index
// and returns the ADL return value
func DeviceFanSetPercent(index int, fanPercent uint32) int {
	return int(C.setADLFanPercent(C.int(index), C.int(fanPercent)))
}

// DeviceTemperature fetches and returns temperature for a device index
func DeviceTemperature(index int) uint32 {
	temperature := uint32(0)

	temperature = uint32(C.getADLTemp(C.int(index)))

	return temperature
}

// DeviceFanAutoManage sets auto-management of fanspeed for a device index
// and returns the ADL return value
func DeviceFanAutoManage(index int) int {
	return int(C.setADLFanAutoManage(C.int(index)))
}
