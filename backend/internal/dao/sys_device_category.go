// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"backend/internal/dao/internal"
)

// internalSysDeviceCategoryDao is internal type for wrapping internal DAO implements.
type internalSysDeviceCategoryDao = *internal.SysDeviceCategoryDao

// sysDeviceCategoryDao is the data access object for table sys_device_category.
// You can define custom methods on it to extend its functionality as you wish.
type sysDeviceCategoryDao struct {
	internalSysDeviceCategoryDao
}

var (
	// SysDeviceCategory is globally public accessible object for table sys_device_category operations.
	SysDeviceCategory = sysDeviceCategoryDao{
		internal.NewSysDeviceCategoryDao(),
	}
)

// Fill with you ideas below.