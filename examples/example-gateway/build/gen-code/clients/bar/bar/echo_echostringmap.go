// Code generated by thriftrw v1.6.0. DO NOT EDIT.
// @generated

package bar

import (
	"errors"
	"fmt"
	"go.uber.org/thriftrw/wire"
	"strings"
)

type Echo_EchoStringMap_Args struct {
	Arg map[string]*BarResponse `json:"arg,required"`
}

type _Map_String_BarResponse_MapItemList map[string]*BarResponse

func (m _Map_String_BarResponse_MapItemList) ForEach(f func(wire.MapItem) error) error {
	for k, v := range m {
		if v == nil {
			return fmt.Errorf("invalid [%v]: value is nil", k)
		}
		kw, err := wire.NewValueString(k), error(nil)
		if err != nil {
			return err
		}
		vw, err := v.ToWire()
		if err != nil {
			return err
		}
		err = f(wire.MapItem{Key: kw, Value: vw})
		if err != nil {
			return err
		}
	}
	return nil
}

func (m _Map_String_BarResponse_MapItemList) Size() int {
	return len(m)
}

func (_Map_String_BarResponse_MapItemList) KeyType() wire.Type {
	return wire.TBinary
}

func (_Map_String_BarResponse_MapItemList) ValueType() wire.Type {
	return wire.TStruct
}

func (_Map_String_BarResponse_MapItemList) Close() {
}

func (v *Echo_EchoStringMap_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Arg == nil {
		return w, errors.New("field Arg of Echo_EchoStringMap_Args is required")
	}
	w, err = wire.NewValueMap(_Map_String_BarResponse_MapItemList(v.Arg)), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _Map_String_BarResponse_Read(m wire.MapItemList) (map[string]*BarResponse, error) {
	if m.KeyType() != wire.TBinary {
		return nil, nil
	}
	if m.ValueType() != wire.TStruct {
		return nil, nil
	}
	o := make(map[string]*BarResponse, m.Size())
	err := m.ForEach(func(x wire.MapItem) error {
		k, err := x.Key.GetString(), error(nil)
		if err != nil {
			return err
		}
		v, err := _BarResponse_Read(x.Value)
		if err != nil {
			return err
		}
		o[k] = v
		return nil
	})
	m.Close()
	return o, err
}

func (v *Echo_EchoStringMap_Args) FromWire(w wire.Value) error {
	var err error
	argIsSet := false
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TMap {
				v.Arg, err = _Map_String_BarResponse_Read(field.Value.GetMap())
				if err != nil {
					return err
				}
				argIsSet = true
			}
		}
	}
	if !argIsSet {
		return errors.New("field Arg of Echo_EchoStringMap_Args is required")
	}
	return nil
}

func (v *Echo_EchoStringMap_Args) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	fields[i] = fmt.Sprintf("Arg: %v", v.Arg)
	i++
	return fmt.Sprintf("Echo_EchoStringMap_Args{%v}", strings.Join(fields[:i], ", "))
}

func _Map_String_BarResponse_Equals(lhs, rhs map[string]*BarResponse) bool {
	if len(lhs) != len(rhs) {
		return false
	}
	for lk, lv := range lhs {
		rv, ok := rhs[lk]
		if !ok {
			return false
		}
		if !lv.Equals(rv) {
			return false
		}
	}
	return true
}

func (v *Echo_EchoStringMap_Args) Equals(rhs *Echo_EchoStringMap_Args) bool {
	if !_Map_String_BarResponse_Equals(v.Arg, rhs.Arg) {
		return false
	}
	return true
}

func (v *Echo_EchoStringMap_Args) MethodName() string {
	return "echoStringMap"
}

func (v *Echo_EchoStringMap_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

var Echo_EchoStringMap_Helper = struct {
	Args           func(arg map[string]*BarResponse) *Echo_EchoStringMap_Args
	IsException    func(error) bool
	WrapResponse   func(map[string]*BarResponse, error) (*Echo_EchoStringMap_Result, error)
	UnwrapResponse func(*Echo_EchoStringMap_Result) (map[string]*BarResponse, error)
}{}

func init() {
	Echo_EchoStringMap_Helper.Args = func(arg map[string]*BarResponse) *Echo_EchoStringMap_Args {
		return &Echo_EchoStringMap_Args{Arg: arg}
	}
	Echo_EchoStringMap_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}
	Echo_EchoStringMap_Helper.WrapResponse = func(success map[string]*BarResponse, err error) (*Echo_EchoStringMap_Result, error) {
		if err == nil {
			return &Echo_EchoStringMap_Result{Success: success}, nil
		}
		return nil, err
	}
	Echo_EchoStringMap_Helper.UnwrapResponse = func(result *Echo_EchoStringMap_Result) (success map[string]*BarResponse, err error) {
		if result.Success != nil {
			success = result.Success
			return
		}
		err = errors.New("expected a non-void result")
		return
	}
}

type Echo_EchoStringMap_Result struct {
	Success map[string]*BarResponse `json:"success"`
}

func (v *Echo_EchoStringMap_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)
	if v.Success != nil {
		w, err = wire.NewValueMap(_Map_String_BarResponse_MapItemList(v.Success)), error(nil)
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}
	if i != 1 {
		return wire.Value{}, fmt.Errorf("Echo_EchoStringMap_Result should have exactly one field: got %v fields", i)
	}
	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func (v *Echo_EchoStringMap_Result) FromWire(w wire.Value) error {
	var err error
	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TMap {
				v.Success, err = _Map_String_BarResponse_Read(field.Value.GetMap())
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
		return fmt.Errorf("Echo_EchoStringMap_Result should have exactly one field: got %v fields", count)
	}
	return nil
}

func (v *Echo_EchoStringMap_Result) String() string {
	if v == nil {
		return "<nil>"
	}
	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}
	return fmt.Sprintf("Echo_EchoStringMap_Result{%v}", strings.Join(fields[:i], ", "))
}

func (v *Echo_EchoStringMap_Result) Equals(rhs *Echo_EchoStringMap_Result) bool {
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && _Map_String_BarResponse_Equals(v.Success, rhs.Success))) {
		return false
	}
	return true
}

func (v *Echo_EchoStringMap_Result) MethodName() string {
	return "echoStringMap"
}

func (v *Echo_EchoStringMap_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
