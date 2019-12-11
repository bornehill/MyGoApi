package datastore

import "my-go-api/loader"

type MenuData interface {
	Initialize()
	GetMenuView(parentId int) *[]*loader.MenuOption
}