package container

import (
	"github.com/pkg/errors"
	"sync"
)

type once struct {
	sync.Mutex
	values  map[string]*value
	ongoing map[string]bool
}

type value struct {
	sync.Mutex
	v interface{}
}

func (o *once) Do(key string, fn func() interface{}) interface{} {
	o.Lock()
	if o.values == nil {
		o.values = make(map[string]*value)
		o.ongoing = make(map[string]bool)
	}
	val := o.values[key]
	if val != nil {
		o.Unlock()
		// Lock/Unlock it just so we sync up with the previous call if it takes longer than expected.
		val.Lock()
		val.Unlock()
		return val.v
	}
	val = &value{}
	o.values[key] = val
	if o.ongoing[key] {
		panic(errors.Errorf(`container: recursive service call for "%s"`, key))
	}
	o.ongoing[key] = true
	val.Lock()
	o.Unlock()
	val.v = fn()
	o.ongoing[key] = false
	val.Unlock()
	return val.v
}
