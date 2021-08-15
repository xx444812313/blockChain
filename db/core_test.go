package db

import "testing"

func TestCreateDB(t *testing.T) {
	_, e := CreateDB("xs.db")
	if e != nil {
		t.Fatal(e)
	}
	t.Log("创建bolt数据库成功")
}

func TestCreateBucket(t *testing.T) {
	e := CreateBucket()
	if e != nil {
		t.Fatal(e)
	}
}
