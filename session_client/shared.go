package session_client

import (
	"fmt"
	"github.com/lemontree2015/skynet/config"
	"github.com/lemontree2015/skynet/misc"
	"strconv"
	"strings"
)

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
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
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
				if v >= 0 && v <= 15 {
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
