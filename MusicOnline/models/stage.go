package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type Stage struct {
	Id      int64     `orm:"auto"`
	Title   string    `orm:"size(128)"`
	Content string    `orm:"size(1024)"`
	Time    time.Time `orm:"type(datetime)"`
	Views   int64     `orm:"index"`
}

func init() {
	orm.RegisterModel(new(Stage))
}

func AddStage(title string, content string) error {
	o := orm.NewOrm()

	stage := &Stage{
		Title:   title,
		Content: content,
		Time:    time.Now(),
	}
	_, err := o.Insert(stage)
	if err != nil {
		return err
	}
	return nil
}

func ModifyStage(tid string, title string, content string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	stage := &Stage{Id: tidNum}
	if o.Read(stage) == nil {
		stage.Title = title
		stage.Content = content
		stage.Time = time.Now()
		stage.Views--
		o.Update(stage)
		return nil
	}
	return nil
}

func DeleteStage(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	_, err = o.Delete(&Stage{Id: tidNum})
	if err != nil {
		return err
	}
	return nil
}

func GetStage(tid string) (*Stage, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	stage := new(Stage)

	qs := o.QueryTable("stage")
	err = qs.Filter("id", tidNum).One(stage)
	if err != nil {
		return nil, err
	}

	stage.Views++
	_, err = o.Update(stage)

	return stage, err
}

func GetStage2(tid string) (*Stage, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	stage := new(Stage)

	qs := o.QueryTable("stage")
	err = qs.Filter("id", tidNum).One(stage)
	if err != nil {
		return nil, err
	}

	return stage, err
}

func GetAllStage(isDesc bool) ([]*Stage, error) {
	o := orm.NewOrm()

	stage := make([]*Stage, 0)

	qs := o.QueryTable("stage")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-time").All(&stage)
	} else {
		_, err = qs.All(&stage)
	}

	return stage, err
}
