package main

import (
	_ "backend/internal/packed"

	_ "backend/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"backend/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
