//---------------------------------
//File Name    : interal/scan/scan.go
//Author       : aico
//Mail         : 2237616014@qq.com
//Github       : https://github.com/TBBtianbaoboy
//Site         : https://www.lengyangyu520.cn
//Create Time  : 2021-12-21 13:32:58
//Description  :
//----------------------------------
package scan

import "nas-web/support"

func GetScanType(scanType int32) string {
	switch scanType {
	case 0:
		return support.ScanType_FastMode
	case 1:
		return support.ScanType_MostCommon
	case 2:
		return support.ScanType_PortRange
	case 3:
		return support.ScanType_PortSingle
	case 4:
		return support.ScanType_PortService
	default:
		return support.ScanType_Default
	}
}
