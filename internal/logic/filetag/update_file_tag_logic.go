package filetag

import (
	"context"

	"github.com/iot-synergy/synergy-file/internal/svc"
	"github.com/iot-synergy/synergy-file/internal/types"
	"github.com/iot-synergy/synergy-file/internal/utils/dberrorhandler"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFileTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFileTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFileTagLogic {
	return &UpdateFileTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateFileTagLogic) UpdateFileTag(req *types.FileTagInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.FileTag.UpdateOneID(*req.Id).
		SetNotNilStatus(req.Status).
		SetNotNilName(req.Name).
		SetNotNilRemark(req.Remark).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
