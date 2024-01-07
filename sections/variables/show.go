package variables

import (
	"fmt"
)

func Show() {
	fmt.Printf("globalString: %v, globalInt: %v, globalBool: %v\n", globalString, globalInt, globalBool)
	fmt.Printf("globalStringHello: %v\n", globalStringHello)
	fmt.Printf("gX: %v, gY: %v, gA: %v, gB: %v, gZ: %v, gC: %v, gD: %v\n", gX, gY, gA, gB, gZ, gC, gD)

	localInt := ExportedFunction()
	fmt.Printf("localInt: %v\n", localInt)
}
