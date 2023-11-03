package optional

import (
	"errors"
	"fmt"
)

/**
  @author : Jerbe - The porter from Earth
  @time : 2023/10/23 18:34
  @describe :
*/

var (
	empty          = &Optional{value: nil}
	NoValuePresent = errors.New("no value present")
	Nil            = errors.New("nil value")
)

func newEmpty(err error) *Optional {
	return &Optional{value: nil, err: err}
}

type Consumer func(interface{})

type Predicate func(interface{}) bool

type MapFunction func(interface{}) interface{}

type Supplier interface {
	Get() (interface{}, error)
}

func Of(val interface{}) *Optional {
	return &Optional{value: val}
}

func OfNullable(val interface{}) *Optional {
	if val == nil {
		return newEmpty(Nil)
	}
	return Of(val)
}

type Optional struct {
	value interface{}
	err   error
}

func (o *Optional) errEmpty() bool {
	return o.value == nil && o.err != nil
}

func (o *Optional) Err() error {
	return o.err
}

func (o *Optional) Get() (interface{}, error) {
	if o.err != nil {
		return nil, o.err
	}

	if o.value == nil {
		return nil, NoValuePresent
	}
	return o.value, nil
}

func (o *Optional) IsPresent() bool {
	return o.err == nil && o.value != nil
}

func (o *Optional) IfPresent(consumer Consumer) {
	if o.err == nil && o.value != nil {
		consumer(o.value)
	}
}

func (o *Optional) Filter(predicate Predicate) *Optional {
	if o.errEmpty() {
		return o
	}

	if !o.IsPresent() {
		return o
	}

	if predicate(o.value) {
		return o
	}

	return empty
}

func (o *Optional) Map(fn MapFunction) *Optional {
	if !o.IsPresent() {
		return empty
	}
	return OfNullable(fn(o.value))
}

func (o *Optional) FlatMap(fn MapFunction) *Optional {
	if !o.IsPresent() {
		return empty
	}

	res := fn(o.value)
	if res == nil {
		return &Optional{value: nil, err: Nil}
	}
	return res.(*Optional)
}

func (o *Optional) OrElse(value interface{}) (interface{}, error) {
	if o.err != nil {
		return nil, o.err
	}

	if o.value != nil {
		return o.value, nil
	}
	return value, nil
}

func (o *Optional) OrElseGet(supplier Supplier) (interface{}, error) {
	if o.err != nil {
		return nil, o.err
	}
	if o.value != nil {
		return o.value, nil
	}

	return supplier.Get()
}

func (o *Optional) Equals(obj interface{}) bool {
	if o == obj {
		return true
	}

	objOpt, ok := obj.(*Optional)
	if !ok {
		return false
	}

	aOpt, ok := o.value.(*Optional)
	if ok {
		return aOpt.Equals(objOpt.value)
	}
	return o.value == objOpt.value
}

func (o *Optional) String() string {
	var t string
	if opt, ok := o.value.(*Optional); ok {
		t = fmt.Sprintf("%s", opt.String())
	} else {
		t = fmt.Sprintf("%T(%v)", o.value, o.value)
	}

	return fmt.Sprintf("%T[%s]", o, t)
}
