package filters

import (
	"fmt"

	"github.com/astaxie/beego/context"
)

func BeforeRouterFilter(ctx *context.Context) {
	fmt.Println("Before Router filter")
}

func BeforeExecFilter(ctx *context.Context) {
	fmt.Println("Before Exec filter")
}

func AfterExecFilter(ctx *context.Context) {
	fmt.Println("After Exec filter")
}

func FinishRouterFilter(ctx *context.Context) {
	fmt.Println("Finish Router filter")
}
