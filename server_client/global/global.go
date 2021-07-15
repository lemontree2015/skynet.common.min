package global

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/lemontree2015/skynet"
	"github.com/lemontree2015/skynet/config"
	"strconv"
	"strings"
)

const (
	APPID_YISHENG = "10000211"
)

var (
	GIMServerServiceKey  *skynet.ServiceKey   // GIM Server[%v] Service
	GIMServerServiceKeys []*skynet.ServiceKey // GIM Server[%v] Service
)

func init() {
	// 初始化GIMServerServiceKey
	if index, err := config.DefaultInt("gim_server.index"); err == nil {
		name := fmt.Sprintf("GIM Server[%v] Service", index)
		serviceKey := skynet.NewServiceKey(name, "1.0.0")
		GIMServerServiceKey = serviceKey
	} else {
		panic(fmt.Errorf("Can't Find gim_server.index"))
	}

	glog.Infof("GIMServerServiceKey=%v", GIMServerServiceKey)

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

	glog.Infof("GIMServerServiceKeys=%v", GIMServerServiceKeys)

}
