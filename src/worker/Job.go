package worker

import "fmt"

type Job interface {
	Do() error
}



//定义一个实现Job接口的数据
type PrintNum struct {
	Num int
}
//定义对数据的处理
func (s *PrintNum) Do() error {
	fmt.Println("num:", s.Num)
	return nil
}