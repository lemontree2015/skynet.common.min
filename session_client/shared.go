package session_client

import (
	"fmt"
	"github.com/lemontree2015/skynet/config"
	"github.com/lemontree2015/skynet/misc"
	"strconv"
	"strings"
)

// skynet.conf.default
//
// # Shared session service name:
// # GIM Session[0] Service
// # GIM Session[1] Service
// # GIM Session[2] Service
// # GIM Session[3] Service
// # GIM Session[4] Service
// # GIM Session[5] Service
// # GIM Session[6] Service
// # GIM Session[7] Service
// #
// # 0 - 7
// # e.g.
// # 0:1:2:3:4:5:6:7:
// # 0:1:2:
// # 0:1:
// # 0:
func SessionServiceName(account string) string {
	return SessionServiceNameByShared(SessionShared(account))
}

func SessionServiceNameByShared(shared int) string {
	return fmt.Sprintf("GIM Session[%v] Service", shared)
}

func SessionShared(account string) int {
	md5Account := misc.MD5Str(account)
	switch md5Account[0] {
	case '0':
		return 0
	case '1':
		return 0
	case '2':
		return 1
	case '3':
		return 1
	case '4':
		return 2
	case '5':
		return 2
	case '6':
		return 3
	case '7':
		return 3
	case '8':
		return 4
	case '9':
		return 4
	case 'a':
		return 5
	case 'b':
		return 5
	case 'c':
		return 6
	case 'd':
		return 6
	case 'e':
		return 7
	case 'f':
		return 7
	default:
		panic(fmt.Errorf("Logic Error"))
	}
}

func LocalSessionShareds() []int {
	if str, err := config.DefaultString("gim_session.shards"); err == nil {
		shareds := strings.Split(str, ":")
		rets := make([]int, 0, 0)
		for _, s := range shareds {
			if s == "" {
				continue
			}

			if v, err := strconv.ParseInt(s, 10, 64); err == nil {
				if v >= 0 && v <= 7 {
					rets = append(rets, int(v))
				} else {
					panic(fmt.Errorf("gim_session.shards Format Error: %v", str))
				}
			} else {
				panic(fmt.Errorf("gim_session.shards Format Error: %v", str))
			}
		}
		return rets
	} else {
		panic(fmt.Errorf("Can't Find gim_session.shards"))
	}
}
