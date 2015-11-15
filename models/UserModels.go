package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

///region User
func AddUser(uname, pwd string) error {
	o := orm.NewOrm()
	user := &User{
		Name:     uname,
		Password: pwd,
		Created:  time.Now(),
		Updated:  time.Now(),
	}

	_, err := o.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(uname string) error {
	o := orm.NewOrm()

	user := &User{Name: uname}
	if o.Read(user) == nil {
		_, err := o.Delete(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func ModifyUser(uname, pwd string) error {
	o := orm.NewOrm()
	user := &User{Name: uname}
	if o.Read(user) == nil {
		user.Password = pwd
		user.Updated = time.Now()
		_, err := o.Update(user)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetUser(uname string) (*User, error) {

	o := orm.NewOrm()

	user := new(User)

	qs := o.QueryTable("user")
	err := qs.Filter("name", uname).One(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetAllUsers(isDesc bool) (users []*User, err error) {

	o := orm.NewOrm()

	users = make([]*User, 0)
	qs := o.QueryTable("user")
	if isDesc {
		_, err = qs.OrderBy("-created").All(&users)

	} else {
		_, err = qs.All(&users)
	}
	return users, err
}

///endregion
