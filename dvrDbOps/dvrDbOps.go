//Package dvrDbOps provides database interactions for dvr database
//Includes basic CRUD functions as well as initializing the database upon init
package dvrDbOps

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

var o orm.Ormer

const DATABASE = "../database/new.db"

type Dvr struct {
	Uid          int       `orm:"pk;auto"`
	Name         string    `orm:"size(160);unique"`
	IpAddress    string    `orm:"size(20);unique"`
	Version      string    `orm:"size(160)"`
	Status       string    `orm:"size(20)"`
	CreationTime time.Time `orm:"auto_now_add;type(datetime)"`
}

func NewDb() {
	o = orm.NewOrm()
	o.Using("default")
	// err := orm.RunSyncdb("default", true, true)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

//init function initialises the db
func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", DATABASE)
	orm.RegisterModel(new(Dvr))
}

func AddNewDvr(newDvr *Dvr) bool {
	dvr := new(Dvr)
	dvr = newDvr

	if created, id, err := o.ReadOrCreate(dvr, "Name"); err == nil {
		if created {
			fmt.Println("New dvr inserted. Id:", id)
			return true
		} else {
			fmt.Println("This dvr already exists. Name:", dvr.Name)
			return true
		}
	} else {
		fmt.Println("error creating object", err)
		return false
	}
}

func GetAliveDvr() ([]*Dvr, int64) {
	var dvrs []*Dvr
	count, err := o.QueryTable("dvr").Filter("status", "alive").All(&dvrs)
	if err != nil {
		fmt.Println(err)
		return nil, -1
	}
	return dvrs, count
}

func GetAllDvr() ([]*Dvr, int64) {
	var dvrs []*Dvr
	count, err := o.QueryTable("dvr").All(&dvrs)
	if err != nil {
		fmt.Println(err)
		return nil, -1
	}
	return dvrs, count
}

func GetDvrById(id int) *Dvr {
	dvr := new(Dvr)
	err := o.QueryTable("dvr").Filter("uid", id).One(dvr)
	if err == orm.ErrMultiRows {
		fmt.Println("Returned multirows not one")
		return nil
	} else if err == orm.ErrNoRows {
		fmt.Println("no row found")
		return nil
	} else {
		return dvr
	}
}

func GetDvrByName(name string) *Dvr {
	dvr := new(Dvr)
	err := o.QueryTable("dvr").Filter("name", name).One(dvr)
	if err == orm.ErrMultiRows {
		fmt.Println("Returned multirows not one")
		return nil
	} else if err == orm.ErrNoRows {
		fmt.Println("no row found")
		return nil
	} else {
		return dvr
	}
}

func GetDvrByIpAddr(ipAddr string) *Dvr {
	dvr := new(Dvr)
	err := o.QueryTable("dvr").Filter("ipAddress", ipAddr).One(dvr)
	if err == orm.ErrMultiRows {
		fmt.Println("Returned multirows not one")
		return nil
	} else if err == orm.ErrNoRows {
		fmt.Println("no row found")
		return nil
	} else {
		return dvr
	}
}

func UpdateDvrName(name string, newName string) bool {
	dvr := new(Dvr)
	dvr.Name = name
	if o.Read(dvr, "Name") == nil {
		dvr.Name = newName
		_, err := o.Update(dvr, "Name")
		if err != nil {
			return false
		} else {
			return true
		}
	} else {
		return false
	}
}

func UpdateDvrIpAddr(ipAddr string, newIpAddr string) bool {
	dvr := new(Dvr)
	dvr.IpAddress = ipAddr
	if o.Read(dvr, "IpAddress") == nil {
		dvr.IpAddress = newIpAddr
		_, err := o.Update(dvr, "IpAddress")
		if err != nil {
			fmt.Println("Update failed", err)
			return false
		} else {
			return true
		}
	} else {
		fmt.Println("no dvr found in specified name", ipAddr)
		return false
	}
}

func UpdateDvrStatusByName(name string, status string) bool {
	dvr := new(Dvr)
	dvr.Name = name
	if o.Read(dvr, "Name") == nil {
		dvr.Status = status
		_, err := o.Update(dvr, "Status")
		if err != nil {
			fmt.Println("Update failed", err)
			return false
		} else {
			return true
		}
	} else {
		fmt.Println("no dvr found in specified name", name)
		return false
	}
}

func UpdateDvrStatusById(id int, status string) bool {
	dvr := new(Dvr)
	dvr.Uid = id
	if o.Read(dvr) == nil {
		dvr.Status = status
		_, err := o.Update(dvr, "Status")
		if err != nil {
			fmt.Println("Update failed", err)
			return false
		} else {
			return true
		}
	} else {
		fmt.Println("no dvr found in specified id", id)
		return false
	}
}

func UpdateDvrStatusByIpAddr(ipAddr string, status string) bool {
	dvr := new(Dvr)
	dvr.IpAddress = ipAddr
	if o.Read(dvr, "IpAddress") == nil {
		dvr.Status = status
		_, err := o.Update(dvr, "Status")
		if err != nil {
			fmt.Println("Update failed", err)
			return false
		} else {
			return true
		}
	} else {
		fmt.Println("no dvr found in specified ip address", ipAddr)
		return false
	}
}

func DeleteDvrById(id int) bool {
	dvr := new(Dvr)
	dvr.Uid = id
	if num, err := o.Delete(dvr); err == nil {
		fmt.Println("Deleting dvr by primary key", num)
		return true
	} else {
		fmt.Println("Error deleting Dvr #%d:", id, err)
		return false
	}
}

func DeleteDvrByName(name string) bool {
	dvr := new(Dvr)
	dvr.Name = name
	if o.Read(dvr, "Name") == nil {
		if num, err := o.Delete(dvr); err == nil {
			fmt.Println("Deleting dvr by primary key", num)
			return true
		} else {
			fmt.Println("Error deleting Dvr by name %s:", name, err)
			return false
		}
	} else {
		return false
	}
}

func DeleteDvrByIpAddr(ipAddr string) bool {
	dvr := new(Dvr)
	dvr.IpAddress = ipAddr
	if o.Read(dvr, "Name") == nil {
		if num, err := o.Delete(dvr); err == nil {
			fmt.Println("Deleting dvr by primary key", num)
			return true
		} else {
			fmt.Println("Error deleting Dvr by ipAddress %s:", ipAddr, err)
			return false
		}
	} else {
		return false
	}
}
