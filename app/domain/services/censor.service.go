package services

import (
	"censor.services/app/domain/dtos"
	"censor.services/app/domain/models/mongo_models"
	"censor.services/app/domain/repositories"
	"censor.services/pkg/utils"
	"context"
	"fmt"
	ac_automaton "github.com/henrion-y/base.services/infra/ac.automaton"
	"github.com/henrion-y/base.services/infra/censor"
	"time"
)

type CensorService interface {
	CensorText(c context.Context, dto *dtos.CensorTextDto) (*censor.CensorResult, error)
	CensorImage(c context.Context, dto *dtos.CensorImageDto) (*censor.CensorResult, error)
}

type censorService struct {
	censorClient            censor.Client
	acAutoMachine           *ac_automaton.AcAutoMachine
	sensitiveWordRepository repositories.SensitiveWordRepository
}

func (s *censorService) CensorText(c context.Context, dto *dtos.CensorTextDto) (*censor.CensorResult, error) {
	censorResult := &censor.CensorResult{
		CensorType:    censor.CensorTypeText,
		CensorContent: dto.Text,
	}
	result := s.acAutoMachine.Query(dto.Text)
	if len(result) != 0 {
		censorResult.InterceptStatus = true
		return censorResult, nil
	}
	var err error
	censorResult, err = s.censorClient.CensorText(dto.Text)
	if err != nil {
		return nil, err
	}

	if censorResult.InterceptStatus {
		go s.acAutoMachineAddPattern(censorResult)
	}

	return censorResult, nil
}

func (s *censorService) CensorImage(c context.Context, dto *dtos.CensorImageDto) (*censor.CensorResult, error) {
	return s.censorClient.CensorImage(dto.ImageUrl)
}

func (s *censorService) acAutoMachineAddPattern(censorResult *censor.CensorResult) {
	var acNeedBuild bool
	for i := range censorResult.ReviewLabel {
		s.acAutoMachine.AddPattern(censorResult.ReviewLabel[i].ReviewContent)
		acNeedBuild = true
		sensitiveWord := &mongo_models.TSensitiveWord{
			ReviewContent: censorResult.ReviewLabel[i].ReviewContent,
			Label:         censorResult.ReviewLabel[i].Label,
			Rate:          censorResult.ReviewLabel[i].Rate,
			CreatedAt:     time.Now(),
		}
		sensitiveWord.WordId, _ = utils.GenerateEntityID(sensitiveWord.WordId)
		_, err := s.sensitiveWordRepository.Create(context.Background(), sensitiveWord)
		if err != nil {
			fmt.Println("acAutoMachineAddPattern err : ", err)
		}
	}

	if acNeedBuild {
		s.acAutoMachine.Build()
	}
	return
}

func (s *censorService) initAcAutoMachine() error {
	words, err := s.sensitiveWordRepository.FindByRate(context.Background(), 0)
	if err != nil {
		return err
	}
	for i := range words {
		s.acAutoMachine.AddPattern(words[i].ReviewContent)
	}
	s.acAutoMachine.Build()
	return nil
}

func NewCensorService(censorClient censor.Client, sensitiveWordRepository repositories.SensitiveWordRepository) (CensorService, error) {
	acAutoMachine := ac_automaton.NewAcAutoMachine()

	service := &censorService{
		acAutoMachine:           acAutoMachine,
		censorClient:            censorClient,
		sensitiveWordRepository: sensitiveWordRepository,
	}

	err := service.initAcAutoMachine()
	if err != nil {
		return nil, err
	}

	return service, nil
}
