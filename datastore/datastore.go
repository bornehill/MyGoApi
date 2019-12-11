package datastore

import "github.com/bornehill/mygoapi/loader"

type MenuData interface {
	Initialize()
	GetMenuView(parentId int) *[]*loader.MenuOption
}