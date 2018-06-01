package rspmsg

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *RspMsg
	}{
		{name: "basic", want: &RspMsg{Status: StatusSuccessful}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewS(t *testing.T) {
	tests := []struct {
		name string
		want *RspMsg
	}{
		{name: "basic", want: &RspMsg{Status: StatusSuccessful}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewS(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewF(t *testing.T) {
	tests := []struct {
		name string
		want *RspMsg
	}{
		{name: "basic", want: &RspMsg{Status: StatusFailed}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewF(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSWithDataMap(t *testing.T) {
	data := make(M)
	tests := []struct {
		name string
		want *RspMsg
	}{
		{name: "basic", want: &RspMsg{Status: StatusSuccessful, Data: data}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSWithDataMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSWithDataMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFWithDataMap(t *testing.T) {
	tests := []struct {
		name string
		want *RspMsg
	}{
		{name: "basic", want: &RspMsg{Status: StatusFailed, Data: make(M)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFWithDataMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFWithDataMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRspMsg_IsSuccessful(t *testing.T) {
	type fields struct {
		Status RspStatus
		Data   interface{}
		Code   interface{}
		Desc   interface{}
		Meta   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "basic", fields: fields{Status: StatusSuccessful}, want: true},
		{name: "basic", fields: fields{Status: StatusFailed}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rspmsg := &RspMsg{
				Status: tt.fields.Status,
				Data:   tt.fields.Data,
				Code:   tt.fields.Code,
				Desc:   tt.fields.Desc,
				Meta:   tt.fields.Meta,
			}
			if got := rspmsg.IsSuccessful(); got != tt.want {
				t.Errorf("RspMsg.IsSuccessful() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRspMsg_IsFailed(t *testing.T) {
	type fields struct {
		Status RspStatus
		Data   interface{}
		Code   interface{}
		Desc   interface{}
		Meta   interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "basic", fields: fields{Status: StatusSuccessful}, want: false},
		{name: "basic", fields: fields{Status: StatusFailed}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rspmsg := &RspMsg{
				Status: tt.fields.Status,
				Data:   tt.fields.Data,
				Code:   tt.fields.Code,
				Desc:   tt.fields.Desc,
				Meta:   tt.fields.Meta,
			}
			if got := rspmsg.IsFailed(); got != tt.want {
				t.Errorf("RspMsg.IsFailed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRspMsg_ToJson(t *testing.T) {
	type fields struct {
		Status RspStatus
		Data   interface{}
		Code   interface{}
		Desc   interface{}
		Meta   interface{}
	}
	rspS := New()
	bsS, _ := json.Marshal(rspS)

	rspF := NewF()
	bsF, _ := json.Marshal(rspF)

	rsp1 := NewSWithDataMap()
	bs1, _ := json.Marshal(rsp1)

	rsp2 := NewFWithDataMap()
	bs2, _ := json.Marshal(rsp2)

	tests := []struct {
		name    string
		fields  fields
		wantBs  []byte
		wantErr bool
	}{
		{name: "basic-S", fields: fields{Status: rspS.Status}, wantBs: bsS, wantErr: false},
		{name: "basic-F", fields: fields{Status: StatusFailed}, wantBs: bsF, wantErr: false},
		{name: "basic-datamap-S", fields: fields{Status: rsp1.Status, Data: rsp1.Data}, wantBs: bs1, wantErr: false},
		{name: "basic-datamap-F", fields: fields{Status: rsp2.Status, Data: rsp2.Data}, wantBs: bs2, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rspmsg := &RspMsg{
				Status: tt.fields.Status,
				Data:   tt.fields.Data,
				Code:   tt.fields.Code,
				Desc:   tt.fields.Desc,
				Meta:   tt.fields.Meta,
			}
			gotBs, err := rspmsg.ToJson()
			if (err != nil) != tt.wantErr {
				t.Errorf("RspMsg.ToJson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBs, tt.wantBs) {
				t.Errorf("RspMsg.ToJson() = %v, want %v", gotBs, tt.wantBs)
			}
		})
	}
}

func TestRspMsg_SetDataMap(t *testing.T) {
	type fields struct {
		Status RspStatus
		Data   interface{}
		Code   interface{}
		Desc   interface{}
		Meta   interface{}
	}
	type args struct {
		key   string
		value interface{}
	}

	rspS := New()
	rspF := NewF()
	rsp1 := NewSWithDataMap()
	rsp2 := NewFWithDataMap()

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "basic-S", fields: fields{Status: rspS.Status}, args: args{key: "a", value: "b"}, want: true},
		{name: "basic-F", fields: fields{Status: rspF.Status}, args: args{key: "a", value: 3}, want: true},
		{name: "basic-datamap-S", fields: fields{Status: rsp1.Status, Data: rsp1.Data}, args: args{key: "a", value: []int{2, 3, 4}}, want: true},
		{name: "basic-datamap-F", fields: fields{Status: rsp2.Status, Data: rsp2.Data}, args: args{key: "a", value: rspS}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rspmsg := &RspMsg{
				Status: tt.fields.Status,
				Data:   tt.fields.Data,
				Code:   tt.fields.Code,
				Desc:   tt.fields.Desc,
				Meta:   tt.fields.Meta,
			}
			if got := rspmsg.SetDataMap(tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("RspMsg.SetDataMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
