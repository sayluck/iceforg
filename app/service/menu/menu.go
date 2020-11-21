package menu

import (
	"context"
	"fmt"
	"iceforg/app/common"
	. "iceforg/app/log"
	"iceforg/app/model"
	"iceforg/pkg/utils"
	"sort"
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
	menuM.Creator = menu.Context.Value(common.UserID).(string)
	return menuM.Save()
}

func List(pageNum string) (interface{}, error) {
	var (
		menuM model.Menu
		err   error
		menus interface{}
	)
	menuM.PageNum, err = strconv.Atoi(pageNum)
	if err != nil {
		return nil, err
	}
	menus, err = menuM.List()
	if err != nil {
		return nil, err
	}
	ms := menus.([]*model.Menu)

	sort.Slice(ms, func(i, j int) bool {
		return ms[i].Level >= ms[j].Level && ms[i].Sort <= ms[j].Sort
	})
	mTrees := []*MenuTree{}
	for _, v := range ms {
		mt := &MenuTree{}
		if err := utils.TramsStruct(v, mt); err != nil {
			return nil, err
		}
		mt.Code = v.Code
		mTrees = append(mTrees, mt)
	}
	return constructMenuTree(mTrees), nil
}

func constructMenuTree(mTrees []*MenuTree) []*MenuTree {
	menu := mTrees[0]
	if menu.Level <= 1 {
		return mTrees
	}
	mts := []*MenuTree{}
	for _, v := range mTrees[1:] {
		if menu.SupCode == v.Code {
			v.SubMenu = append(v.SubMenu, menu)
			mts = append(mTrees[:0], mTrees[1:]...)
		}
	}
	return constructMenuTree(mts)
}
