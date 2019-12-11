package datastore

import (
	"log"
	"os"
	"strings"

	"github.com/bornehill/mygoapi/loader"
)

type Menus struct {
	Options *[]*loader.MenuOption `json:"options"`
}

func (b *Menus) Initialize() {
	filename := "./assets/menu.csv"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	b.Options = loader.LoadData(file)
}

func (b *Menus) GetMenuView(parentId int) *[]*loader.MenuOption {
	ret := Filter(b.Options, func(v *loader.MenuOption) bool {
		return v.ParentId == parentId
	})
	return &ret
}

func Filter(vs *[]*loader.MenuOption, f func(*loader.MenuOption) bool) *[]*loader.MenuOption {
	vsf := make([]*loader.MenuOption, 0)
	for _, v := range *vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return &vsf
}