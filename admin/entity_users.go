package vertc_go_admin

type UserSSO struct {
	Id          int     `json:"id"`
	Name        *string `json:"name"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       *string `json:"email"`
	IsSuperuser bool    `json:"is_superuser"`
	IsStaff     bool    `json:"is_staff"`
	IsActive    bool    `json:"is_active"`
	Image       *string `json:"image"`
}

func (UserSSO) TableName() string {
	return "vertadmin_users"
}
