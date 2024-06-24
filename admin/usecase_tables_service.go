package vertc_go_admin

type UseCaseTable struct {
	repo IRepositoryTable
}

func NewUseCaseTable(repo IRepositoryTable) *UseCaseTable {
	return &UseCaseTable{
		repo: repo,
	}
}

func (uc *UseCaseTable) List(table Table, filters Filters) (response ResponseList, err error) {
	return uc.repo.List(table, filters)
}

func (uc *UseCaseTable) Get(table Table, id int) (response map[string]interface{}, err error) {
	return uc.repo.Get(table, id)
}

func (uc *UseCaseTable) Create(table Table, fields Fields) (response ResponseCreateUpdateDelete, err error) {
	return uc.repo.Create(table, fields)
}

func (uc *UseCaseTable) Update(table Table, fields Fields, id int) (response ResponseCreateUpdateDelete, err error) {
	return uc.repo.Update(table, fields, id)
}

func (uc *UseCaseTable) Delete(table Table, ids []int) (response ResponseCreateUpdateDelete, err error) {
	return uc.repo.Delete(table, ids)
}
