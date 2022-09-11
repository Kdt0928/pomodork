package common

import (
	"fmt"
	"strconv"
)

// CreateId 連番生成
func CreateId(previousId string) string {
	// IDのインクリメント
	serialNumber, _ := strconv.Atoi(previousId[2:12])
	serialNumber += 1
	userId := previousId[:2] + fmt.Sprintf("%010d", serialNumber)
	return userId
}
