package itsm

import "testing"


func TestItsm(t *testing.T) {
	RequestItsm("/api/misc/get_change_list",ItsmOpenApiHost,MethodStrGet,nil)
}