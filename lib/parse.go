package lib

type CreateSubpage struct{}

func (c CreateSubpage) Parse(variable string) []string {
	return []string{"鈴木", "佐藤", "田中", "堀井"}
}
