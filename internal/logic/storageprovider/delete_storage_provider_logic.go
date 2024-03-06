package storageprovider

import (
	"context"

	"github.com/iot-synergy/synergy-file/ent/cloudfile"
	"github.com/iot-synergy/synergy-file/ent/storageprovider"
	"github.com/iot-synergy/synergy-file/internal/svc"
	"github.com/iot-synergy/synergy-file/internal/types"
	"github.com/iot-synergy/synergy-file/internal/utils/cloud"
	"github.com/iot-synergy/synergy-file/internal/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteStorageProviderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStorageProviderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStorageProviderLogic {
	return &DeleteStorageProviderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStorageProviderLogic) DeleteStorageProvider(req *types.IDsReq) (*types.BaseMsgResp, error) {
	check, err := l.svcCtx.DB.CloudFile.Query().Where(cloudfile.HasStorageProvidersWith(storageprovider.IDIn(req.Ids...))).
		Count(l.ctx)

	if err != nil {
		if err != nil {
			return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
		}
	}

	if check != 0 {
		return nil, errorx.NewCodeInvalidArgumentError(l.svcCtx.Trans.Trans(l.ctx, "storage_provider.hasFileError"))
	}

	_, err = l.svcCtx.DB.StorageProvider.Delete().Where(storageprovider.IDIn(req.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	l.svcCtx.CloudStorage = cloud.NewCloudServiceGroup(l.svcCtx.DB)

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.DeleteSuccess)}, nil
}
