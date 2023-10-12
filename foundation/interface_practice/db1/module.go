package main

import "context"

type XormSession XormDB

type XormDB struct {
	db *XormSession
}

func (x *XormDB) Insert(ctx context.Context, instance interface{}) {

}

type Trade struct {
	db *DBCommon
}

func (t *Trade) InsertTrade() {
	//t.db.Xxx()
}

func (t *Trade) AddDB() {
	//t.db = new(XormDB)
}
