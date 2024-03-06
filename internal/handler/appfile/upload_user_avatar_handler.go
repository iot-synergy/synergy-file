package appfile

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"

	"github.com/iot-synergy/synergy-file/internal/logic/appfile"
	"github.com/iot-synergy/synergy-file/internal/svc"
)

// swagger:route post /userAvatar/upload appfile UploadUserAvatar
//
// userAvatar file upload | 上传用户头像文件
//
// userAvatar file upload | 上传用户头像文件
//
// Responses:
//  200: CloudFileInfoResp

func UploadUserAvatarHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := appfile.NewUploadUserAvatarLogic(r, r.Context(), svcCtx)
		resp, err := l.UploadUserAvatar()
		if err != nil {
			err = svcCtx.Trans.TransError(r.Context(), err)
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
