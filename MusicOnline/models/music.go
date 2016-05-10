package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
)

type Music struct {
	Id       int64  `orm:"auto"`
	Name     string `orm:"size(128)"`
	Singer   string `orm:"size(128)"`
	Special  string `orm:"size(128)"`
	Category string `orm:"size(128)"`
	Image    string `orm:"size(128)"`
}

func init() {
	orm.RegisterModel(new(Music))
}

func AddMusic(name string, singer string, special string, category string) error {
	o := orm.NewOrm()

	music := &Music{
		Name:     name,
		Singer:   singer,
		Special:  special,
		Category: category,
	}
	_, err := o.Insert(music)
	if err != nil {
		return err
	}
	return nil
}

func GetMusic(tid string) (*Music, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	music := new(Music)

	qs := o.QueryTable("music")
	err = qs.Filter("id", tidNum).One(music)
	if err != nil {
		return nil, err
	}

	return music, err
}

func ModifyMusic(tid string, name string, singer string, special string, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	music := &Music{Id: tidNum}
	if o.Read(music) == nil {
		music.Name = name
		music.Singer = singer
		music.Special = special
		music.Category = category
		o.Update(music)
		return nil
	}
	return nil
}

func DeleteMusic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	_, err = o.Delete(&Music{Id: tidNum})
	if err != nil {
		return err
	}
	return nil
}


func GetAllMusic(isDesc bool) ([]*Music, error) {
	o := orm.NewOrm()

	music := make([]*Music, 0)

	qs := o.QueryTable("music")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-time").All(&music)
	} else {
		_, err = qs.All(&music)
	}

	return music, err
}
