package grpc

import (
	"censor.services/app/domain/dtos"
	censorPb "censor.services/app/domain/proto"
	"censor.services/app/domain/services"
	"context"
	"github.com/henrion-y/base.services/infra/censor"
)

type CensorGrpcService struct {
	censorService services.CensorService
	censorPb.UnimplementedCensorGrpcServiceServer
}

func NewCensorGrpcService(censorService services.CensorService) censorPb.CensorGrpcServiceServer {
	return &CensorGrpcService{censorService: censorService}
}

func (g *CensorGrpcService) CensorText(c context.Context, censorTextRequest *censorPb.CensorTextRequest) (*censorPb.CensorResultResponse, error) {
	result := &censorPb.CensorResultResponse{
		Code:    censorPb.RET_CODE_SUCCESS,
		Message: "",
		Data:    nil,
	}

	censorResult, err := g.censorService.CensorText(c, &dtos.CensorTextDto{Text: censorTextRequest.GetText()})
	if err != nil {
		result.Code = censorPb.RET_CODE_ERROR
		result.Message = err.Error()
		return result, err
	}

	result.Data = toCensorResultResponse(censorResult)
	return result, nil
}
func (g *CensorGrpcService) CensorImage(c context.Context, censorImageRequest *censorPb.CensorImageRequest) (*censorPb.CensorResultResponse, error) {
	result := &censorPb.CensorResultResponse{
		Code:    censorPb.RET_CODE_SUCCESS,
		Message: "",
		Data:    nil,
	}

	censorResult, err := g.censorService.CensorImage(c, &dtos.CensorImageDto{ImageUrl: censorImageRequest.GetImage()})
	if err != nil {
		result.Code = censorPb.RET_CODE_ERROR
		result.Message = err.Error()
		return result, err
	}

	result.Data = toCensorResultResponse(censorResult)
	return result, nil
}

func toCensorResultResponse(censorResult *censor.CensorResult) *censorPb.CensorResult {
	censorResultResponse := &censorPb.CensorResult{
		CensorContent:   censorResult.CensorContent,
		CensorType:      censorResult.CensorType,
		Suggestion:      censorResult.Suggestion,
		InterceptStatus: censorResult.InterceptStatus,
	}
	for i := range censorResult.ReviewLabel {
		censorResultResponse.ReviewLabel = append(censorResultResponse.GetReviewLabel(), &censorPb.ReviewLabel{
			ReviewContent: censorResult.ReviewLabel[i].ReviewContent,
			Label:         censorResult.ReviewLabel[i].Label,
			Rate:          censorResult.ReviewLabel[i].Rate,
		})
	}
	return censorResultResponse
}
