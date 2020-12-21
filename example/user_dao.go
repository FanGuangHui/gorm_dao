package example

import (
	"errors"

	"gorm.io/gorm"
)

var (
	ErrCreateUser = errors.New("create User failed")
	ErrDeleteUser = errors.New("delete User failed")
	ErrGetUser    = errors.New("get User failed")
	ErrUpdateUser = errors.New("update User failed")
)

// UserDao
type UserDao struct {
	Db *gorm.DB
	*User
}

// Add add one record
func (t *UserDao) Add() (err error) {
	if err = t.Db.Create(t).Error; err != nil {

		err = ErrCreateUser
		return
	}
	return
}

// Delete delete record
func (t *UserDao) Delete() (err error) {
	if err = t.Db.Delete(t).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}

// Updates update record
func (t *UserDao) Updates(m map[string]interface{}) (err error) {
	if err = t.Db.Model(&User{}).Where("id = ?", t.ID).Updates(m).Error; err != nil {

		err = ErrUpdateUser
		return
	}
	return
}

// GetAll get all record
func (t *UserDao) GetAll() (ret []*User, err error) {
	if err = t.Db.Find(&ret).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// GetCount get count
func (t *UserDao) GetCount() (ret int64) {
	t.Db.Model(&User{}).Count(&ret)
	return
}

type QueryUserForm struct {
	CreatedAt *FieldData `json:"createdAt" form:"createdAt"`
	UpdatedAt *FieldData `json:"updatedAt" form:"updatedAt"`
	Age       *FieldData `json:"age" form:"age"`
	Email     *FieldData `json:"email" form:"email"`
	Order     []string   `json:"order" form:"order"`
	PageNum   int        `json:"pageNum" form:"pageNum"`
	PageSize  int        `json:"pageSize" form:"pageSize"`
}

//  GetList get list some field value or some condition
func (t *UserDao) GetList(q *QueryUserForm) (ret []*User, err error) {
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
func (t *UserDao) SetQueryByID(id uint) *User {
	t.ID = id
	return t.User
}

// GetByID get one record by ID
func (t *UserDao) GetByID() (err error) {
	if err = t.Db.First(t, "id = ?", t.ID).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByID delete record by ID
func (t *UserDao) DeleteByID() (err error) {
	if err = t.Db.Delete(t, "id = ?", t.ID).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}

// QueryByName query cond by Name
func (t *UserDao) SetQueryByName(name string) *User {
	t.Name = name
	return t.User
}

// GetByName get one record by Name
func (t *UserDao) GetByName() (err error) {
	if err = t.Db.First(t, "name = ?", t.Name).Error; err != nil {

		err = ErrGetUser
		return
	}
	return
}

// DeleteByName delete record by Name
func (t *UserDao) DeleteByName() (err error) {
	if err = t.Db.Delete(t, "name = ?", t.Name).Error; err != nil {

		err = ErrDeleteUser
		return
	}
	return
}
