package menu

import (
	"fmt"
	. "iceforg/app/log"
	"iceforg/app/model"
	"iceforg/pkg/utils"
)

func AddMenu(menu *MenuReq) (string, error) {
	var err error

	if menu == nil {
		errMsg := "menu is nil"
		Log.Error(errMsg)
		err = fmt.Errorf(errMsg)
		return "", err
	}

	menuM := model.Menu{}

	err = utils.TramsStruct(&menu, &menuM)
	if err != nil {
		Log.Error(err)
		return "", err
	}
	return menuM.Save()
}
