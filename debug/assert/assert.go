package assert

import (
	"fmt"
	"log"
	"runtime/debug"

	"github.com/niyoh120/go-utility/inspect"
)

func Assert(exp bool, msg string) {
	if !exp {
		msg = fmt.Sprintf("Assert failed!\n", msg)
		debug.PrintStack()
		log.Fatal(msg)
	}
}

func True(exp bool) {
	Assert(exp, "")
}

func False(exp bool) {
	Assert(!exp, "")
}

func Equal(v1, v2 interface{}) {
	msg := fmt.Sprintf("%+v != %+v!", v1, v2)
	Assert(v1 == v2, msg)
}

func NotEqual(v1, v2 interface{}) {
	msg := fmt.Sprintf("%+v == %+v!", v1, v2)
	Assert(v1 != v2, msg)
}

func Nil(v interface{}) {
	msg := fmt.Sprintf("%+v is not nil!", v)
	Assert(v == nil, msg)
}

func NoError(error error){
    Nil(error)
}

func NotZero(v interface{}) {
	vv := inspect.UnderlyingValueOf(v)
	msg := fmt.Sprintf("%+v is zero!", v)
	Assert(vv.IsValid(), msg)
}

func InvalidBranch() {
	Assert(false, "Invalid Branch!")
}
