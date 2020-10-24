package menu

import (
	"context"
	"fmt"
	. "iceforg/app/log"
	"iceforg/app/model"
	"iceforg/pkg/utils"
	"strconv"
)

func AddMenu(menu *MenuAddReq) (string, error) {
	var err error

	if menu == nil {
		errMsg := "menu is nil"
		IceLog.Error(context.Background(), errMsg)
		return "", fmt.Errorf(errMsg)
	}

	menuM := model.Menu{}

	err = utils.TramsStruct(&menu, &menuM)
	if err != nil {
		IceLog.Error(menu.Context, err)
		return "", err
	}
	return menuM.Save()
}

func List(pageNum string) (interface{}, error) {
	var (
		menuM model.Menu
		err   error
	)
	menuM.PageNum, err = strconv.Atoi(pageNum)
	if err != nil {
		return nil, err
	}
	return menuM.List()
}
