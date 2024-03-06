package cloudfiletag

import (
	"context"

	"github.com/iot-synergy/synergy-file/internal/svc"
	"github.com/iot-synergy/synergy-file/internal/types"
	"github.com/iot-synergy/synergy-file/internal/utils/dberrorhandler"

	"github.com/iot-synergy/synergy-common/i18n"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCloudFileTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCloudFileTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCloudFileTagLogic {
	return &UpdateCloudFileTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCloudFileTagLogic) UpdateCloudFileTag(req *types.CloudFileTagInfo) (*types.BaseMsgResp, error) {
	err := l.svcCtx.DB.CloudFileTag.UpdateOneID(*req.Id).
		SetNotNilStatus(req.Status).
		SetNotNilName(req.Name).
		SetNotNilRemark(req.Remark).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
