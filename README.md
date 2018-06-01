# go-rspmsg
simple response message for golang

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

## install
go get github.com/darkdarkfruit/go-rspmsg

## usage
import "github.com/darkdarkfruit/go-rspmsg/rspmsg"

