package querybuilder

type QueryBuilder struct {
	table string
}

type QueryBuilderInterface interface {
	Where(column string, operator string, value interface{}) QueryBuilderInterface
	Limit(limit int) QueryBuilderInterface
	Offset(offset int) QueryBuilderInterface
	Get() ([]interface{}, error)
}

func NewQueryBuilder(table string) *QueryBuilder {
	return &QueryBuilder{
		table: table,
	}
}

func (qb *QueryBuilder) Where(column string, operator string, value interface{}) QueryBuilderInterface {
	return qb
}

func (qb *QueryBuilder) Limit(limit int) QueryBuilderInterface {
	return qb
}

func (qb *QueryBuilder) Offset(offset int) QueryBuilderInterface {
	return qb
}

func (qb *QueryBuilder) Get() ([]interface{}, error) {
	return nil, nil
}