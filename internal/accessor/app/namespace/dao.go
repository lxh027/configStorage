package namespace

type Dao struct{}

func (dao *Dao) NewNamespace() error {
	return nil
}

func (dao *Dao) GetUserNamespace(id int) ([]Namespace, error) {
	return nil, nil
}
