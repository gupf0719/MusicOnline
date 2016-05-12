package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type News struct {
	Id       int64     `orm:"auto"`
	Title    string    `orm:"size(128)"`
	Content  string    `orm:"size(1024)"`
	Category string    `orm:"size(128)"`
	Time     time.Time `orm:"type(datetime)"`
	Views    int64     `orm:"index"`
}

func init() {
	orm.RegisterModel(new(News))
}

func AddNews(title string, content string, category string) error {
	o := orm.NewOrm()

	news := &News{
		Title:    title,
		Content:  content,
		Category: category,
		Time:     time.Now(),
	}
	_, err := o.Insert(news)
	if err != nil {
		return err
	}
	return nil
}

func ModifyNews(tid string, title string, content string, category string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	news := &News{Id: tidNum}
	if o.Read(news) == nil {
		news.Title = title
		news.Content = content
		news.Category = category
		news.Time = time.Now()
		o.Update(news)
		return nil
	}
	return nil
}

func DeleteNews(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	_, err = o.Delete(&News{Id: tidNum})
	if err != nil {
		return err
	}
	return nil
}

func GetNews(tid string) (*News, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	news := new(News)

	qs := o.QueryTable("news")
	err = qs.Filter("id", tidNum).One(news)
	if err != nil {
		return nil, err
	}

	news.Views++
	_, err = o.Update(news)

	return news, err
}

func GetNews2(tid string) (*News, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	o := orm.NewOrm()
	news := new(News)

	qs := o.QueryTable("news")
	err = qs.Filter("id", tidNum).One(news)
	if err != nil {
		return nil, err
	}

	return news, err
}

func GetAllNews2(cate string,isDesc bool) ([]*News, error) {
	o := orm.NewOrm()

	news := make([]*News, 0)

	qs := o.QueryTable("news")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-time").Limit(3).All(&news)
	} else {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.All(&news)
	}

	return news, err
}


func GetAllNews(cate string,isDesc bool) ([]*News, error) {
	o := orm.NewOrm()

	news := make([]*News, 0)

	qs := o.QueryTable("news")

	var err error
	if isDesc {
		_, err = qs.OrderBy("-time").All(&news)
	} else {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.All(&news)
	}

	return news, err
}
