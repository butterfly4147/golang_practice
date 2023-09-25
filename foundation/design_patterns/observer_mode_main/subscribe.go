package main

// Observer 订阅者 （Subscriber） 接口声明了通知接口。
// 在绝大多数情况下， 该接口仅包含一个 update更新方法。
// 该方法可以拥有多个参数， 使发布者能在更新时传递事件的详细信息。
type Observer interface {
	update(s string)
	getID() string
}

// effect:
// - judge that the struct has implemented all the method definition of the interface
// - the pointer of the struct can be used as the interface type, but the struct without its pointer
var _ Observer = (*Customer)(nil)

type Customer struct {
	id string
}

// implement
func (c *Customer) update(s string) {
	println(s)
}

func (c *Customer) getID() string {
	return c.id
}
