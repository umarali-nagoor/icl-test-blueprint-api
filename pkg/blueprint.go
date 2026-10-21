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

func (s *Service) BulkPublish(ctx context.Context, items []*Blueprint) error {
	for _, item := range items {
		s.log.Infof("publishing blueprint: %+v", item)
		if err := s.store.Save(ctx, item); err != nil {
			s.log.Errorf("save failed: %v", err)
			return err
		}
		s.log.Infof("blueprint published: %+v", item)
	}
	return nil
}
