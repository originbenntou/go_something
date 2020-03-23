package main

import (
	"testing"
)

// ダメな例
//func TestInitRepository(t *testing.T) {
//	var us UserService
//	// serviceに依存してしまっている
//	// All関数をモックしたいとき、service自体をモックするしかない
//	// しかしそれではserviceのテストとはならない
//	us.initRepository()
//	r, err := us.All()
//	if err != nil {
//		t.Fatalf("fail")
//	}
//	if reflect.TypeOf(r) != reflect.TypeOf([]User{}) {
//		t.Fatalf("fail")
//	}
//}


// OKなやつ
func TestUserService_All(t *testing.T) {
	repo := &UserRepositoryStub{}
	s := NewUserService(repo)
	/* 以下任意にテスト */
	if _, err := s.All(); err != nil {
		t.Error(err)
	}
}

// このあたりがモック
type UserRepositoryStub struct {}
func (*UserRepositoryStub) All() ([]User, error) {
	return []User{}, nil
}
