package loader

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type MenuOption struct {
	Id            int     `json:"id"`
	GroupId       int     `json:"groupId"`
	Nivel         int     `json:"nivel"`
	GroupName     string  `json:"groupName"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	ImageName     string  `json:"imageName"`
	GoForm        string  `json:"goForm"`
	ParentId      int     `json:"parentId"`
}

func LoadData(r io.Reader) *[]*MenuOption {
	reader := csv.NewReader(r)

	ret := make([]*MenuOption, 0, 0)

	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Println("End of File")
			break
		} else if err != nil {
			log.Println(err)
			break
		}
		id, _ := strconv.Atoi(row[0])
		groupId, _ := strconv.Atoi(row[1])
		nivel, _ := strconv.Atoi(row[2])
		parentId, _ := strconv.Atoi(row[8])

		if err != nil {
			log.Println(err)
		}
		menu := &MenuOption{
			Id:           id,
			GroupId:      groupId,
			Nivel:        nivel,
			GroupName:    row[3],
			Name:         row[4],
			Description:  row[5],
			ImageName:    row[6],
			GoForm:       row[7],
			ParentId:     parentId,
		}

		if err != nil {
			log.Fatalln(err)
		}

		ret = append(ret, menu)
	}
	return &ret
}