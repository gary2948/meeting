package utils

import (
	"strconv"
)

//空间容量显示数据转换
func SizeConvert(kbsize int64) (ret string) {
	if kbsize < 1024 {
		ret = strconv.FormatInt(kbsize, 10) + "K"
	} else {
		mbsize := float64(kbsize) / 1024
		if mbsize < 1024 {
			ret = strconv.FormatFloat(mbsize, 'f', 2, 64) + "M"
		} else {
			gbsize := mbsize / 1024
			ret = strconv.FormatFloat(gbsize, 'f', 2, 64) + "G"
		}
	}

	return ret
}

//空间容量百分比
func SizePercent(usize, msize int64) (ret string) {
	p := float64(usize) / float64(msize)
	ret = strconv.FormatFloat(p, 'f', 0, 64) + "%"
	return ret
}
