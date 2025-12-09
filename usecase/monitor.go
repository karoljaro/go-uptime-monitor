package usecase

import (
	"context"
	"fmt"

	"github.com/karoljaro/go-uptime-monitor/domain"
)

type MonitorUseCase struct {
	targetRepo domain.TargetRepository
	resultRepo domain.ResultRepository
	alertRepo  domain.AlertRepository
	httpClient domain.HTTPClient
	idGenerator domain.IDGenerator
}

func NewMonitorUseCase(
	targetRepo domain.TargetRepository,
	resultRepo domain.ResultRepository,
	alertRepo domain.AlertRepository,
	httpClient domain.HTTPClient,
	idGenerator domain.IDGenerator,
) *MonitorUseCase {
	return &MonitorUseCase{
		targetRepo: targetRepo,
		resultRepo: resultRepo,
		alertRepo:  alertRepo,
		httpClient: httpClient,
		idGenerator: idGenerator,
	}
}

func (u *MonitorUseCase) CheckTarget(ctx context.Context, targetID string) error {
	target, err := u.targetRepo.FindByID(targetID)
	if err != nil {
		return err
	}

	httpResp, err := u.httpClient.Check(ctx, target.URL)
	if err != nil {
		return err
	}

	var status string
	if httpResp.Error != nil {
		status = "ERROR"
	} else if httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
		status = "OK"
	} else if httpResp.StatusCode >= 500 {
		status = "SERVER_ERROR"
	} else {
		status = "CLIENT_ERROR"
	}

	genResultID := u.idGenerator.Generate()

	result := domain.NewResult(
		genResultID,
		targetID,
		status,
		httpResp.StatusCode,
		httpResp.ResponseTime,
	)

	prevResult, err := u.resultRepo.GetLastByTargetID(targetID)
	if err != nil {
		prevResult = nil
	}

	u.resultRepo.Save(result)

	if prevResult != nil && httpResp.StatusCode != 200 && prevResult.StatusCode == 200 {
		genAlertID := u.idGenerator.Generate()
		newAlert := domain.NewAlert(
			genAlertID,
			targetID,
			status,
			fmt.Sprintf("Target %s is %s", target.URL, status),
		)

		u.alertRepo.Save(newAlert)
	}

	if httpResp.StatusCode == 200 {
		unresolvedAlerts, _ := u.alertRepo.GetUnresolvedByTargetID(targetID)

		for _, alert := range unresolvedAlerts {
			alert.Resolve()
			u.alertRepo.Update(alert)
		}
	}

	return nil
}
