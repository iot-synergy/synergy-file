package appfile

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/duke-git/lancet/v2/datetime"
	"github.com/iot-synergy/synergy-common/enum/errorcode"
	"github.com/iot-synergy/synergy-common/i18n"
	"github.com/iot-synergy/synergy-common/utils/pointy"
	"github.com/iot-synergy/synergy-common/utils/uuidx"
	"github.com/iot-synergy/synergy-file/internal/svc"
	"github.com/iot-synergy/synergy-file/internal/types"
	"github.com/iot-synergy/synergy-file/internal/utils/dberrorhandler"
	"github.com/iot-synergy/synergy-file/internal/utils/filex"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserAvatarLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadUserAvatarLogic(r *http.Request, ctx context.Context, svcCtx *svc.ServiceContext) *UploadUserAvatarLogic {
	return &UploadUserAvatarLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserAvatarLogic) UploadUserAvatar() (resp *types.CloudFileInfoResp, err error) {
	err = l.r.ParseMultipartForm(l.svcCtx.Config.UploadConf.MaxVideoSize)
	if err != nil {
		logx.Error("fail to parse the multipart form")
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "file.parseFormFailed"))
	}

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		logx.Error("the value of file cannot be found")
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "file.parseFormFailed"))
	}
	defer file.Close()

	// judge if the suffix is legal
	// 校验后缀是否合法
	dotIndex := strings.LastIndex(handler.Filename, ".")
	// if there is no suffix, reject it
	// 拒绝无后缀文件
	if dotIndex == -1 {
		logx.Errorw("reject the file which does not have suffix")
		return nil, errorx.NewCodeError(errorcode.InvalidArgument,
			l.svcCtx.Trans.Trans(l.ctx, "file.wrongTypeError"))
	}

	fileName, fileSuffix := handler.Filename[:dotIndex], handler.Filename[dotIndex+1:]
	fileUUID := uuidx.NewUUID()
	storeFileName := fileUUID.String() + "." + fileSuffix
	userId := l.ctx.Value("uid").(string)

	// judge if the file size is over max size
	// 判断文件大小是否超过设定值
	fileType := strings.Split(handler.Header.Get("Content-Type"), "/")[0]
	if fileType != "image" && fileType != "video" && fileType != "audio" {
		fileType = "other"
	}
	err = filex.CheckOverSize(l.ctx, l.svcCtx, fileType, handler.Size)
	if err != nil {
		logx.Errorw("the file is over size", logx.Field("type", fileType),
			logx.Field("userId", userId), logx.Field("size", handler.Size),
			logx.Field("fileName", handler.Filename))
		return nil, err
	}

	provider := "userAvatar"
	// if l.r.MultipartForm.Value["provider"] != nil && l.r.MultipartForm.Value["provider"][0] != "" {
	// 	provider = l.r.MultipartForm.Value["provider"][0]
	// } else {
	// 	provider = l.svcCtx.CloudStorage.DefaultProvider
	// }
	relativeSrc := fmt.Sprintf("%s/%s/%s/%s/%s",
		l.svcCtx.CloudStorage.ProviderData[provider].Folder, userId,
		datetime.FormatTimeToStr(time.Now(), "yyyy-mm-dd"),
		fileType,
		storeFileName)

	url, err := l.UploadToProvider(file, relativeSrc, provider)
	if err != nil {
		return nil, err
	}

	// store to database
	data, err := l.svcCtx.DB.CloudFile.Create().
		SetName(fileName).
		SetFileType(filex.ConvertFileTypeToUint8(fileType)).
		SetStorageProvidersID(l.svcCtx.CloudStorage.ProviderData[provider].Id).
		SetURL(url).
		SetSize(uint64(handler.Size)).
		SetUserID(userId).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, nil)
	}

	return &types.CloudFileInfoResp{
		BaseDataInfo: types.BaseDataInfo{
			Code: 0,
			Msg:  i18n.Success,
			Data: "",
		},
		Data: types.CloudFileInfo{
			BaseUUIDInfo: types.BaseUUIDInfo{
				Id:        pointy.GetPointer(data.ID.String()),
				CreatedAt: pointy.GetPointer(data.CreatedAt.UnixMilli()),
			},
			State:       pointy.GetPointer(data.State),
			Name:        pointy.GetPointer(data.Name),
			Url:         pointy.GetPointer(data.URL),
			RelativeSrc: pointy.GetPointer(relativeSrc),
			Size:        pointy.GetPointer(data.Size),
			FileType:    pointy.GetPointer(data.FileType),
			UserId:      pointy.GetPointer(data.UserID),
		},
	}, nil
}

func (l *UploadUserAvatarLogic) UploadToProvider(file multipart.File, fileName, provider string) (url string, err error) {
	if client, ok := l.svcCtx.CloudStorage.CloudStorage[provider]; ok {
		_, err := client.PutObjectWithContext(l.ctx, &s3.PutObjectInput{
			Bucket: aws.String(l.svcCtx.CloudStorage.ProviderData[provider].Bucket),
			Key:    aws.String(fileName),
			Body:   file,
		})
		if err != nil {
			logx.Errorw("failed to upload object", logx.Field("detail", err))
			var aerr awserr.Error
			if errors.As(err, &aerr) && aerr.Code() == request.CanceledErrorCode {
				return url, errorx.NewCodeInternalError("upload canceled due to timeout")
			} else {
				return url, errorx.NewCodeInternalError("failed to upload object")
			}
		}

		return fmt.Sprintf("https://%s.%s%s",
			l.svcCtx.CloudStorage.ProviderData[provider].Bucket,
			l.svcCtx.CloudStorage.ProviderData[provider].Endpoint, fileName), nil
	}

	return url, nil
}