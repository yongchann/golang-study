package main

type LockService struct {
	lockRepository LockRepository
}

func NewLockService(lockRepository LockRepository) LockService {
	return LockService{
		lockRepository: lockRepository,
	}
}

func (svc LockService) updateWithOptimisticLock(id uint64) error {
	product, err := svc.lockRepository.FindById(id)
	if err != nil {
		return err
	}

	if err := svc.lockRepository.UpdateWithVersion(product); err != nil {
		return err
	}

	return nil
}
