package services

import (
	"context"
	"fmt"
	"generator/internal/application/use_cases"
	"generator/internal/domain/services"
	"log"
	"os"
	"time"
)

var iOrderPollerSatisfiedConstraint services.ISyntheticContinuousOrdersProcessor = (*SyntheticContinuousOrdersProcessor)(nil)

type SyntheticContinuousOrdersProcessor struct {
	sendContinuousOrderUseCase   *use_cases.SendContinuousOrderUseCase
	getSnapshotUseCase           *use_cases.GetSnapshotUseCase
	createContinuousOrderUseCase *use_cases.CreateContinuousOrderFromOrderBookSnapshotUseCase
}

/*func (s SyntheticOrdersGenerator) PollOrdersGenerator(ctx context.Context) error {
	for {
		snapshot, err := s.getSnapshotUseCase.Execute()
		if err != nil {
			return err
		}
		order, err := s.createContinuousOrderUseCase.Execute(snapshot)
		if err != nil {
			return err
		}
		err = s.sendContinuousOrderUseCase.Execute(order)
		if err != nil {
			return err
		}
	}

	return nil
}*/

func (processor *SyntheticContinuousOrdersProcessor) StartProcessing(ctx context.Context) error {
	go processor.startPollingLoop(ctx)
	return nil
}

// error must be returned
func (processor *SyntheticContinuousOrdersProcessor) startPollingLoop(ctx context.Context) {
	sleepTimeString := os.Getenv("GENERATE_ORDER_SLEEP_TIME")
	if sleepTimeString == "" {
		log.Fatal("GENERATE_ORDER_SLEEP_TIME must be set")
	}

	sleepTime, errParse := time.ParseDuration(sleepTimeString)
	if errParse != nil {
		log.Fatal("GENERATE_ORDER_SLEEP_TIME must be float number")
	}

	ticker := time.NewTicker(sleepTime)
	defer ticker.Stop()

	log.Println("[INFO] Starting background polling loop")

	for {
		select {
		case <-ticker.C:
			if err := processor.runOnce(); err != nil {
				log.Printf("[ERROR] Failed to generate/send order: %v", err)
			}
		case <-ctx.Done():
			log.Println("[INFO] Stopping polling loop")
			return
		}
	}
}

func (processor *SyntheticContinuousOrdersProcessor) runOnce() error {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("[PANIC] Recovered from panic: %v", r)
		}
	}()

	snapshot, err := processor.getSnapshotUseCase.Execute()
	if err != nil {
		return fmt.Errorf("failed to get snapshot: %w", err)
	}

	order, err := processor.createContinuousOrderUseCase.Execute(snapshot)
	if err != nil {
		return fmt.Errorf("failed to create order: %w", err)
	}

	err = processor.sendContinuousOrderUseCase.Execute(order)
	if err != nil {
		return fmt.Errorf("failed to send order: %w", err)
	}

	log.Printf("[SUCCESS] Sent order for pair: %v", order.Pair)
	return nil
}

func NewSyntheticContinuousOrdersProcessor(sendContinuousOrderUseCase *use_cases.SendContinuousOrderUseCase,
	getSnapshotUseCase *use_cases.GetSnapshotUseCase,
	createContinuousOrderUseCase *use_cases.CreateContinuousOrderFromOrderBookSnapshotUseCase) (*SyntheticContinuousOrdersProcessor, error) {
	syntheticOrdersGenerator := &SyntheticContinuousOrdersProcessor{}

	syntheticOrdersGenerator.sendContinuousOrderUseCase = sendContinuousOrderUseCase
	syntheticOrdersGenerator.getSnapshotUseCase = getSnapshotUseCase
	syntheticOrdersGenerator.createContinuousOrderUseCase = createContinuousOrderUseCase

	return syntheticOrdersGenerator, nil
}
