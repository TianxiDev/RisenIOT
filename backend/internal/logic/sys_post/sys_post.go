package sysPost

import (
	"backend/api/v1/system"
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model/do"
	"backend/internal/model/entity"
	"backend/internal/service"
	"backend/library/liberr"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysPost struct {
}

func init() {
	service.RegisterSysPost(New())
}

func New() *sSysPost {
	return &sSysPost{}
}

// List 岗位列表
func (s *sSysPost) List(ctx context.Context, req *system.PostSearchReq) (res *system.PostSearchRes, err error) {
	res = new(system.PostSearchRes)
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.SysPost.Ctx(ctx)
		if req != nil {
			if req.PostCode != "" {
				m = m.Where("post_code like ?", "%"+req.PostCode+"%")
			}
			if req.PostName != "" {
				m = m.Where("post_name like ?", "%"+req.PostName+"%")
			}
			if req.Status != "" {
				m = m.Where("status", gconv.Uint(req.Status))
			}
		}
		res.Total, err = m.Count()
		liberr.ErrIsNil(ctx, err, "获取岗位失败")
		if req.PageNum == 0 {
			req.PageNum = 1
		}
		if req.PageSize == 0 {
			req.PageSize = consts.PageSize
		}
		res.CurrentPage = req.PageNum
		err = m.Page(req.PageNum, req.PageSize).Order("post_sort asc,post_id asc").Scan(&res.PostList)
		liberr.ErrIsNil(ctx, err, "获取岗位失败")
	})
	return
}

func (s *sSysPost) Add(ctx context.Context, req *system.PostAddReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysPost.Ctx(ctx).Insert(do.SysPost{
			PostCode:  req.PostCode,
			PostName:  req.PostName,
			PostSort:  req.PostSort,
			Status:    req.Status,
			Remark:    req.Remark,
			CreatedBy: service.UserCtx().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "添加岗位失败")
	})
	return
}

func (s *sSysPost) Edit(ctx context.Context, req *system.PostEditReq) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysPost.Ctx(ctx).WherePri(req.PostId).Update(do.SysPost{
			PostCode:  req.PostCode,
			PostName:  req.PostName,
			PostSort:  req.PostSort,
			Status:    req.Status,
			Remark:    req.Remark,
			UpdatedBy: service.UserCtx().GetUserId(ctx),
		})
		liberr.ErrIsNil(ctx, err, "修改岗位失败")
	})
	return
}

func (s *sSysPost) Delete(ctx context.Context, ids []int) (err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		_, err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().PostId+" in(?)", ids).Delete()
		liberr.ErrIsNil(ctx, err, "删除失败")
	})
	return
}

// GetUsedPost 获取正常状态的岗位
func (s *sSysPost) GetUsedPost(ctx context.Context) (list []*entity.SysPost, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		err = dao.SysPost.Ctx(ctx).Where(dao.SysPost.Columns().Status, 1).
			Order(dao.SysPost.Columns().PostSort + " ASC, " + dao.SysPost.Columns().PostId + " ASC ").Scan(&list)
		liberr.ErrIsNil(ctx, err, "获取岗位数据失败")
	})
	return
}
