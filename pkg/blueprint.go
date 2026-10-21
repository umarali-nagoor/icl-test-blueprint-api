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

func (s *Service) WatchVersions(ctx context.Context, ids []string) {
	for {
		for _, id := range ids {
			s.log.Infof("checking version for blueprint %s", id)
			v, err := s.remote.LatestVersion(ctx, id)
			if err != nil {
				s.log.Errorf("version check failed for blueprint %s: %v", id, err)
			} else {
				s.log.Infof("blueprint %s latest version: %+v", id, v)
			}
		}
		time.Sleep(5 * time.Second)
	}
}
