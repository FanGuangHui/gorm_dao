package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateAdmin = errors.New("create Admin failed")
	ErrDeleteAdmin = errors.New("delete Admin failed")
	ErrGetAdmin    = errors.New("get Admin failed")
	ErrUpdateAdmin = errors.New("update Admin failed")
)

// AdminDao
type AdminDao struct {
	Db *gorm.DB
	*Admin
}

// Add add one record
func (t *AdminDao) Add() (err error) {
	if err = t.Db.Create(t).Error; err != nil {

		err = ErrCreateAdmin
		return
	}
	return
}

// Delete delete record
func (t *AdminDao) Delete() (err error) {
	if err = t.Db.Delete(t).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}

// Updates update record
func (t *AdminDao) Updates(m map[string]interface{}) (err error) {
	if err = t.Db.Model(&Admin{}).Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateAdmin
		return
	}
	return
}

// GetAll get all record
func (t *AdminDao) GetAll() (ret []*Admin, err error) {
	if err = t.Db.Find(&ret).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// GetCount get count
func (t *AdminDao) GetCount() (ret int64) {
	t.Db.Model(&Admin{}).Count(&ret)
	return
}

type QueryAdminForm struct {
	CreatedAt *FieldData `json:"createdAt" form:"createdAt"`
	UpdatedAt *FieldData `json:"updatedAt" form:"updatedAt"`
	Age       *FieldData `json:"age" form:"age"`
	Email     *FieldData `json:"email" form:"email"`
	Order     []string   `json:"order" form:"order"`
	PageNum   int        `json:"pageNum" form:"pageNum"`
	PageSize  int        `json:"pageSize" form:"pageSize"`
}

//  GetList get list some field value or some condition
func (t *AdminDao) GetList(q *QueryAdminForm) (ret []*Admin, err error) {
	// order
	if len(q.Order) > 0 {
		for _, v := range q.Order {
			t.Db = t.Db.Order(v)
		}
	}
	// pageSize
	if q.PageSize != 0 {
		t.Db = t.Db.Limit(q.PageSize)
	}
	// pageNum
	if q.PageNum != 0 {
		q.PageNum = (q.PageNum - 1) * q.PageSize
		t.Db = t.Db.Offset(q.PageNum)
	}

	// CreatedAt
	if q.CreatedAt != nil {
		t.Db = t.Db.Where("created_at"+q.CreatedAt.Symbol+"?", q.CreatedAt.Value)
	}
	// UpdatedAt
	if q.UpdatedAt != nil {
		t.Db = t.Db.Where("updated_at"+q.UpdatedAt.Symbol+"?", q.UpdatedAt.Value)
	}
	// Age
	if q.Age != nil {
		t.Db = t.Db.Where("age"+q.Age.Symbol+"?", q.Age.Value)
	}
	// Email
	if q.Email != nil {
		t.Db = t.Db.Where("email"+q.Email.Symbol+"?", q.Email.Value)
	}
	if err = t.Db.Find(&ret).Error; err != nil {
		return
	}
	return
}

// QueryByID query cond by ID
func (t *AdminDao) SetQueryByID(id uint) *Admin {
	t.ID = id
	return t.Admin
}

// GetByID get one record by ID
func (t *AdminDao) GetByID() (err error) {
	if err = t.Db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *AdminDao) DeleteByID() (err error) {
	if err = t.Db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}

// QueryByName query cond by Name
func (t *AdminDao) SetQueryByName(name string) *Admin {
	t.Name = name
	return t.Admin
}

// GetByName get one record by Name
func (t *AdminDao) GetByName() (err error) {
	if err = t.Db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetAdmin
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *AdminDao) DeleteByName() (err error) {
	if err = t.Db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteAdmin
		return
	}
	return
}
