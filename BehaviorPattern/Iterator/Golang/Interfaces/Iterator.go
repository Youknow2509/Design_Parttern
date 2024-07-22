
package interfaces

type Iterator interface {
	HasNext() bool
	Next() interface{}
}
   