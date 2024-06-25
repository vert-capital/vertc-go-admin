package vertc_go_admin

import "gorm.io/gorm"

type RepositoryUsers struct {
	DB *gorm.DB
}

func NewRepositoryUsers(db *gorm.DB) *RepositoryUsers {
	return &RepositoryUsers{DB: db}
}

func (u *RepositoryUsers) GetByEmail(email string) (user *UserSSO, err error) {
	err = u.DB.Table(user.TableName()).Where("email = ?", email).First(&user).Error
	return
}

func (u *RepositoryUsers) CreateOrUpdateUser(user *UserSSO) error {
	return u.DB.Table(user.TableName()).Save(user).Error
}
