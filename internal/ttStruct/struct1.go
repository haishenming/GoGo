package ttStruct

import (
	"GoGo/pkg/thereAreSomeStruct"
	"fmt"
)

func ShowStruct1() {
	var V1, V2, V3 thereAreSomeStruct.Vertax
	var p1, p2 *thereAreSomeStruct.Vertax

	V1 = thereAreSomeStruct.Vertax{1, 2}
	V2 = thereAreSomeStruct.Vertax{X: 2}
	V3 = thereAreSomeStruct.Vertax{}
	p1 = &V1
	p2 = &thereAreSomeStruct.Vertax{3,4}

	fmt.Println(V1)
	fmt.Println(V2)
	fmt.Println(V3)
	fmt.Println(p1)
	fmt.Println(p2)
}


