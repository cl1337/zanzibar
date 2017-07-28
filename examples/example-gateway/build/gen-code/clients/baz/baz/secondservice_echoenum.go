// Code generated by thriftrw v1.3.0
// @generated

package baz

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type SecondService_EchoEnum_Args struct {
	Arg *Fruit `json:"arg,omitempty"`
}

func _Fruit_ptr(v Fruit) *Fruit {
	return &v
}

func (v *SecondService_EchoEnum_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		v.Arg = _Fruit_ptr(FruitApple)
	}
	{
		w, err = v.Arg.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Fruit_Read(w wire.Value) (Fruit, error) {
	var v Fruit
	err := v.FromWire(w)
	return v, err
}

func (v *SecondService_EchoEnum_Args) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TI32 {
				var x Fruit
				x, err = _Fruit_Read(field.Value)
				v.Arg = &x
				if err != nil {
					return err
				}
			}
		}
	}
	if v.Arg == nil {
		v.Arg = _Fruit_ptr(FruitApple)
	}
	return nil
}

func (v *SecondService_EchoEnum_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Arg != nil {
		fields[i] = fmt.Sprintf("Arg: %v", *(v.Arg))
		i++
	}
	return fmt.Sprintf("SecondService_EchoEnum_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Fruit_EqualsPtr(lhs, rhs *Fruit) bool {
	if lhs != nil && rhs != nil {
		x := *lhs
		y := *rhs
		return x.Equals(y)
	}
	return lhs == nil && rhs == nil
}

func (v *SecondService_EchoEnum_Args) Equals(rhs *SecondService_EchoEnum_Args) bool {
	if !_Fruit_EqualsPtr(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *SecondService_EchoEnum_Args) MethodName() string {
	return "echoEnum"
}

func (v *SecondService_EchoEnum_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var SecondService_EchoEnum_Helper = struct {
	Args           func(arg *Fruit) *SecondService_EchoEnum_Args
	IsException    func(error) bool
	WrapResponse   func(Fruit, error) (*SecondService_EchoEnum_Result, error)
	UnwrapResponse func(*SecondService_EchoEnum_Result) (Fruit, error)
}{}

func init() {
	SecondService_EchoEnum_Helper.Args = func(arg *Fruit) *SecondService_EchoEnum_Args {
		return &SecondService_EchoEnum_Args{Arg: arg}
	}
	SecondService_EchoEnum_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	SecondService_EchoEnum_Helper.WrapResponse = func(success Fruit, err error) (*SecondService_EchoEnum_Result, error) {
		if err == nil {
			return &SecondService_EchoEnum_Result{Success: &success}, nil
		}
		return nil, err
	}
	SecondService_EchoEnum_Helper.UnwrapResponse = func(result *SecondService_EchoEnum_Result) (success Fruit, err error) {
		if result.Success != nil {
			success = *result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type SecondService_EchoEnum_Result struct {
	Success *Fruit `json:"success,omitempty"`
}

func (v *SecondService_EchoEnum_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("SecondService_EchoEnum_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *SecondService_EchoEnum_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TI32 {
				var x Fruit
				x, err = _Fruit_Read(field.Value)
				v.Success = &x
				if err != nil {
					return err
				}
			}
		}
	}
	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("SecondService_EchoEnum_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *SecondService_EchoEnum_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", *(v.Success))
		i++
	}
	return fmt.Sprintf("SecondService_EchoEnum_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *SecondService_EchoEnum_Result) Equals(rhs *SecondService_EchoEnum_Result) bool {
	if !_Fruit_EqualsPtr(v.Success, rhs.Success) {
		return false
	}
	return true
}

func (v *SecondService_EchoEnum_Result) MethodName() string {
	return "echoEnum"
}

func (v *SecondService_EchoEnum_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}