package global

import (
	"fmt"
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet/config"
	"github.com/lemontree2015/skynet/logger"
	"strconv"
	"strings"
)

var (
	GIMServerServiceKeys []*skynet.ServiceKey // GIM Server[%v] Service
)

func init() {
	// 初始化GIMServerServiceKeys
	if str, err := config.DefaultString("gim_server.indexs"); err == nil {
		indexs := strings.Split(str, ":")
		serviceKeys := make([]*skynet.ServiceKey, 0, 0)
		for _, s := range indexs {
			if s == "" {
				continue
			}

			if v, err := strconv.ParseInt(s, 10, 64); err == nil {
				if v >= 0 {
					name := fmt.Sprintf("GIM Server[%v] Service", v)
					serviceKey := skynet.NewServiceKey(name, "1.0.0")
					serviceKeys = append(serviceKeys, serviceKey)
				} else {
					panic(fmt.Errorf("gim_server.indexs Format Error: %v", str))
				}
			} else {
				panic(fmt.Errorf("gim_server.indexs Format Error: %v", str))
			}

			GIMServerServiceKeys = serviceKeys
		}
	} else {
		panic(fmt.Errorf("Can't Find gim_server.indexs"))
	}
	logger.Logger.Infof("GIMServerServiceKeys=%v", GIMServerServiceKeys)

}
