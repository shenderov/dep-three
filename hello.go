package depthree // import "github.com/shenderov/dep-thee"

import (
	"fmt"
	"github.com/shenderov/dep-two"
)

func Hello() {
	fmt.Println("This is dep three")
	deptwo.Hello()
}
