package pkg

import "go.uber.org/zap"

func (b *Blueprints) Publish(id string) error {
	b.logger.Info("Entering Publish")
	if err := b.repo.Publish(id); err != nil {
		b.logger.Error("publish failed", zap.String("blueprintId", id), zap.Error(err))
		return err
	}
	b.logger.Info("blueprint published", zap.String("blueprintId", id))
	b.logger.Info("Exiting Publish")
	return nil
}
