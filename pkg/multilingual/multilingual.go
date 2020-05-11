package multilingual

import (
	"fmt"
	"iceforg/pkg/common"
	"sync"

	"github.com/magiconair/properties"
)

var (
	defaultPath = ""
	mLang       = new(multilingual)
)

type multilingual struct {
	locker     sync.Mutex
	Properties *properties.Properties
}

func InitMultilingual(path ...string) *multilingual {
	mLang.locker.Lock()
	defer mLang.locker.Unlock()
	if mLang.Properties == nil {
		if len(path) >= 1 {
			defaultPath = path[0]
		}
		err := common.LoadFile(defaultPath, &mLang)
		if err != nil {
			panic(fmt.Sprintf("init multilingual info failed,error msg(%s)", err.Error()))
		}
	}
	return mLang
}

// GetStrMsg get a string msg
func GetStrMsg(err error) string {
	if err == nil {
		return ""
	}
	msg, ok := mLang.Properties.Get(err.Error())
	if !ok {
		return err.Error()
	}
	return msg
}
