package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/iot-synergy/oms-core/rpc/coreclient"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-file/internal/utils/cloud"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/iot-synergy/synergy-file/ent"
	"github.com/iot-synergy/synergy-file/internal/config"
	i18n2 "github.com/iot-synergy/synergy-file/internal/i18n"
	"github.com/iot-synergy/synergy-file/internal/middleware"
)

type ServiceContext struct {
	Config       config.Config
	DB           *ent.Client
	Casbin       *casbin.Enforcer
	Authority    rest.Middleware
	Trans        *i18n.Translator
	CoreRpc      coreclient.Core
	CloudStorage *cloud.CloudServiceGroup
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := ent.NewClient(
		ent.Log(logx.Info), // logger
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(), // debug mode
	)

	rds := c.RedisConf.MustNewUniversalRedis()

	cbn := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher(c.CasbinDatabaseConf.Type,
		c.CasbinDatabaseConf.GetDSN(), c.RedisConf)

	trans := i18n.NewTranslator(i18n2.LocaleFS)

	return &ServiceContext{
		Config:       c,
		DB:           db,
		Casbin:       cbn,
		CoreRpc:      coreclient.NewCore(zrpc.MustNewClient(c.CoreRpc)),
		Authority:    middleware.NewAuthorityMiddleware(cbn, rds, trans).Handle,
		Trans:        trans,
		CloudStorage: cloud.NewCloudServiceGroup(db),
	}
}
