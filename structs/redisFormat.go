package structs

import "time"

var LoginRedis = make(map[string]UserLoginRedis)

type UserLoginRedis struct {
	UserId    int64
	Username  string
	Role      string
	LoginAt   time.Time
	ExpiredAt time.Time
}
