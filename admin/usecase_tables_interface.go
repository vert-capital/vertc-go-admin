package vertc_go_admin

type IRepositoryTable interface {
	List(table Table, filters Filters) (response ResponseList, err error)
	Get(table Table, id int) (response Fields, err error)
	Create(table Table, fields Fields) (response ResponseCreateUpdateDelete, err error)
	Update(table Table, fields Fields, id int) (response ResponseCreateUpdateDelete, err error)
	Delete(table Table, ids []int) (response ResponseCreateUpdateDelete, err error)
}

type IUseCaseTable interface {
	List(table Table, filters Filters) (response ResponseList, err error)
	Get(table Table, id int) (response Fields, err error)
	Create(table Table, fields Fields) (response ResponseCreateUpdateDelete, err error)
	Update(table Table, fields Fields, id int) (response ResponseCreateUpdateDelete, err error)
	Delete(table Table, ids []int) (response ResponseCreateUpdateDelete, err error)
}
