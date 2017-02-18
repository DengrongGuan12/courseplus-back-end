package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Carousel struct {
	Id          int
	Img_url     string `orm:"size(255)"`
	Link_url    string `orm:"size(255)"`
	Priority    int
	State       string
	Create_time time.Time `orm:"auto_now_add;type(datetime)"`
	Update_time time.Time `orm:"auto_now_add;type(datetime)"`
	Delete_time time.Time `orm:"type(datetime)"`
}

func AddCarousel(carousel Carousel) (int64, error) {
	o := orm.NewOrm()
	carousel.State = "UNACTIVE"
	carousel.Create_time = time.Now()
	id, err := o.Insert(&carousel)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}
}

func DeleteCarousel(id int) (int64, error) {
	o := orm.NewOrm()
	carousel := *GetCarouselById(id)
	carousel.State = "DELETE"
	carousel.Delete_time = time.Now()
	result, err := o.Update(&carousel)
	if err == nil {
		return result, nil
	} else {
		return 0, err
	}
}

func UpdateCarousel(carousel Carousel) (int64, error) {
	o := orm.NewOrm()
	carousel.State = "UNACTIVE"
	carousel.Update_time = time.Now()
	id, err := o.Update(&carousel)
	if err == nil {
		return id, nil
	} else {
		return 0, err
	}
}

func GetAllCarousels() []*Carousel {
	o := orm.NewOrm()

	var carousels []*Carousel
	o.QueryTable("carousel").OrderBy("-priority").All(&carousels)
	//fmt.Printf("Returned Rows Num: %s, %s", num, err)

	// if err == nil {
	// 	fmt.Printf("Result Nums: %d\n", num)
	// 	for _, carousel := range carousels {
	// 		fmt.Println(carousel)
	// 	}
	// }
	return carousels
}

func GetCarouselById(id int) *Carousel {
	//变量判断 加错误处理
	if id <= 0 {
		return nil
	}
	c := Carousel{Id: id}
	err := orm.NewOrm().Read(&c)
	if err != nil {
		return nil
	}
	return &c
}
