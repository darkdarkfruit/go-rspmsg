//|--------+--------+-----------+-----------+------------+-------------------------------------------------------|
//| Field  | type   | Required? | Optional? | value      | Meaning                                               |
//|--------+--------+-----------+-----------+------------+-------------------------------------------------------|
//| status | string | *         |           | "StatusSuccessful" or "StatusFailed" | Is the response successful?                           |
//| code   | any    |           | *         |            | CODE for application logic(Normally it is an integer) |
//| data   | any    |           | *         |            | Data(payload) of the response                         |
//| desc   | any    |           | *         |            | Description: normally it's a helping infomation       |
//| meta   | any    |           | *         |            | Meta info. eg: servers/ips chain in distributed env.  |
//|        |        |           |           |            |                                                       |
//|--------+--------+-----------+-----------+------------+-------------------------------------------------------|
//
//    Field:status is always in state: "StatusSuccessful" or "StatusFailed"(represents "Successful", "Failed"), no 3th state.
package rspmsg

import (
	"encoding/json"
	"fmt"
)

const VERSION = "0.1.0"

type RspStatus string

const StatusSuccessful RspStatus = "S"
const StatusFailed RspStatus = "F"

type M map[string]interface{}

type RspMsg struct {
	Status RspStatus   `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Code   interface{} `json:"code,omitempty"`
	Desc   interface{} `json:"desc,omitempty"`
	Meta   interface{} `json:"meta,omitempty"`
}

func (rspmsg *RspMsg) String() string {
	s := fmt.Sprintf("<RspMsg>: status:%#v", rspmsg.Status)
	if rspmsg.Data != nil {
		s += fmt.Sprintf(", %#v", rspmsg.Data)
	}
	if rspmsg.Code != nil {
		s += fmt.Sprintf(", %#v", rspmsg.Code)
	}
	if rspmsg.Desc != nil {
		s += fmt.Sprintf(", %#v", rspmsg.Desc)
	}
	if rspmsg.Meta != nil {
		s += fmt.Sprintf(", %#v", rspmsg.Meta)
	}
	return s

}

// new successful RspMsg by default
func New() *RspMsg {
	return NewS()
}

// new successful RspMsg
func NewS() *RspMsg {
	return &RspMsg{
		Status: StatusSuccessful,
	}
}

// new failed RspMsg
func NewF() *RspMsg {
	return &RspMsg{
		Status: StatusFailed,
	}
}

func NewSWithDataMap() *RspMsg {
	return &RspMsg{
		Status: StatusSuccessful,
		Data:   make(M),
	}
}

func NewFWithDataMap() *RspMsg {
	return &RspMsg{
		Status: StatusFailed,
		Data:   make(M),
	}
}

func (rspmsg *RspMsg) IsSuccessful() bool {
	return rspmsg.Status == StatusSuccessful
}

func (rspmsg *RspMsg) IsFailed() bool {
	return !rspmsg.IsSuccessful()
}

func (rspmsg *RspMsg) ToJson() (bs []byte, err error) {
	return json.Marshal(rspmsg)
}

// warning: this method will set "Data" part to be type: "M"
func (rspmsg *RspMsg) SetDataMap(key string, value interface{}) bool {
	t, ok := rspmsg.Data.(M)
	if !ok{
		rspmsg.Data = make(M)
		t = rspmsg.Data.(M)
	}
	t[key] = value
	return true
}
