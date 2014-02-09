package mongo

import (
	"fmt"
	"labix.org/v2/mgo"
	"testing"
)

var (
	taro   = &User{Name: "Taro", Age: 20}
	hanako = &User{Name: "Hanako", Age: 22}
)

// テストデータを投入
func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("droneTest").C("users")
	c.DropCollection()
	err = c.Insert(taro)
	if err != nil {
		panic(err)
	}
	err = c.Insert(hanako)
	if err != nil {
		panic(err)
	}
}

// Find関数をテスト
func TestFind(t *testing.T) {
	// User(Nmae="Taro")を抽出
	user, err := Find(taro.Name)
	if err != nil {
		t.Error(err)
	}
	if user.Name != taro.Name || user.Age != taro.Age {
		t.Error(fmt.Sprintf("検索結果が不正です。[期待値: %+v][実際: %+v]", taro, user))
	}

	// 存在しないUserを抽出
	user, err = Find("X")
	if err == nil || err.Error() != "not found" {
		t.Error("検索結果が不正です。検索結果が存在します。")
	}
}
