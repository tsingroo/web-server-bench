// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"gozero_bench/internal/svc"
	"gozero_bench/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Gozero_benchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGozero_benchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Gozero_benchLogic {
	return &Gozero_benchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Gozero_benchLogic) Gozero_bench(req *types.Request) (resp *types.Response, err error) {
	resp = &types.Response{
		Message: "my name is " + req.Name,
	}

	return
}
