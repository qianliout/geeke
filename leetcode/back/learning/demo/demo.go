package demo

type Name interface {
	ToName(string2 string)string
}

type Lq struct {
	
}

func (l Lq) ToName(string2 string) string {
	panic("implement me")
}
 