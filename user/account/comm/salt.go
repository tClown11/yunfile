package comm

import (
	"math/rand"
	"time"
)

//积累，如果不在调用前对rand.seek设置种子，会导致重启服务后的前n次随机数与一开始启动的前n次随机的值一样
//所以，倘若每次都想获取的是不同的值(没有规律)，即使服务重启后也一样，就可以在每次重启给其设置一个唯一值(时间戳)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	rand.Seed(time.Now().Unix())
}

//RandStringBytesMaskImpr 参考：https://www.flysnow.org/2019/09/30/how-to-generate-a-random-string-of-a-fixed-length-in-go.html
//GetSalt: 获取一个加密的盐，暂且设定长度为六
func GetSalt() string {
	b := make([]byte, 6)
	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := 6-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
