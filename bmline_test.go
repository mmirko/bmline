package bmline

import (
	"fmt"
	"testing"
)

func TestMetaData(t *testing.T) {
	bb := new(BasmBody)
	bb.BasmMeta = bb.SetMeta("test1", "value1")

	bl := new(BasmLine)
	bl.BasmMeta = bl.SetMeta("test2", "value2")

	be := new(BasmElement)
	be.BasmMeta = be.SetMeta("test3", "value3")
	be.BasmMeta = be.SetMeta("test4", "value4")

	fmt.Println(bb.GetMeta("test1"))
	fmt.Println(bl.GetMeta("test2"))
	fmt.Println(be.GetMeta("test3"))

	bb.RmMeta("test1")
	bl.RmMeta("test1")
	be.RmMeta("test1")

	fmt.Println(bb.GetMeta("test1"))
	fmt.Println(bl.GetMeta("test2"))
	fmt.Println(be.GetMeta("test3"))

	fmt.Println(be.ListMeta())
}
