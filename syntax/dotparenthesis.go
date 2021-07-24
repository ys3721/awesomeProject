package syntax

import "fmt"

func DotParent() {
	var i interface{}
	i = int(42)

	a, ok := i.(int)
	fmt.Printf("a=%d, ok=%v\n", a, ok)

	b, ok := i.(string)
	fmt.Printf("b=%s, ok=%v\n", b, ok)

	type BuildingType int32

	const (
		BuildingType_farm       BuildingType = 1 // 灵田
		BuildingType_decoration BuildingType = 2 // 装饰物
		BuildingType_mine       BuildingType = 3 // 矿
		BuildingType_factory    BuildingType = 4 // 加工厂
		BuildingType_palace     BuildingType = 5 // 大殿
		BuildingType_house      BuildingType = 6 // 住宅
		BuildingType_warehouse  BuildingType = 7 // 仓库
		BuildingType_museplat   BuildingType = 8 // 练功台
	)

	var vv int32 = 9
	bb:= BuildingType(vv)
	bbb := int32(bb)
	fmt.Printf("b:%T, b:%v, bbb:%T, bbb:%v\n", bb, bb, bbb, bbb)
	fmt.Printf("%v\n", bb == BuildingType_warehouse + 2)

}
