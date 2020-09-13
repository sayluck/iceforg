package model

import (
	"iceforg/app/test"
	"iceforg/pkg/utils"
	"testing"
)

func init() {
	test.TestInit()
}

func TestUser_Save(t *testing.T) {
	u := User{
		UserName: "testCase4",
		Password: "testPw",
		NickName: "xiaoyiyi",
	}
	n, err := u.Save()
	if err != nil || n == "" {
		t.Fatalf("user save failed:%v-%v", n, err)
	}
	t.Logf("save num:%v\n", n)
}

func TestUser_DetailByKeyProperty(t *testing.T) {
	u := User{
		UserName: "testCase4",
	}
	data, err := u.DetailByKeyProperty()
	if err != nil {
		t.Fatalf("detail user faield,err:%v", err)
	}
	utils.PrettyJsonPrint(data)
}

func TestUser_IsExistedByKeyProperty(t *testing.T) {
	u := User{
		UserName: "testCase1",
	}
	existed, err := u.IsExistedByKeyProperty()
	if err != nil {
		t.Fatalf("check is existed failed,err:%v", err)
	}
	t.Log(existed)
}
