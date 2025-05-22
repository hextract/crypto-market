package services

import "context"

type ISyntheticContinuousOrdersProcessor interface {
	StartProcessing(context.Context) error
}
