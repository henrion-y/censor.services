package services

import (
	"context"
	"fmt"
	"time"

	"censor.services/app/domain/dtos"
	"censor.services/app/domain/models/mongo_models"
	"censor.services/app/domain/repositories"
	"censor.services/pkg/utils"

	ac_automaton "github.com/henrion-y/base.services/infra/ac.automaton"
	"github.com/henrion-y/base.services/infra/censor"
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
	for i := range censorResult.ReviewLabel {
		s.acAutoMachine.AddPatternAndBuild(censorResult.ReviewLabel[i].ReviewContent)
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

	return
}

func (s *censorService) initAcAutoMachine() error {
	page := 1
	for {
		words, err := s.sensitiveWordRepository.FindByRate(context.Background(), page, 100, 0)
		if err != nil {
			return err
		}

		var patterns []string
		for i := range words {
			patterns = append(patterns, words[i].ReviewContent)
		}
		err = s.acAutoMachine.LoadPatterns(patterns)
		if err != nil {
			return err
		}

		if len(words) < 100 {
			break
		}

		page++
	}

	s.acAutoMachine.StopLoad()
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
