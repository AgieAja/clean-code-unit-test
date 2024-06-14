package checkHealthUsecase

import "clean-code-unit-test/src/checkHealth"

type checkHealthUC struct {
	checkHealthRepo checkHealth.CheckHealthRepository
}

func NewCheckHealthUsecase(checkHealthRepo checkHealth.CheckHealthRepository) checkHealth.CheckHealthUsecase {
	return &checkHealthUC{checkHealthRepo}
}

//implement GetVersion
func (c *checkHealthUC) GetVersion() (string, error) {
	version, err := c.checkHealthRepo.RetrieveVersion()
	if err != nil {
		return "", err
	}

	return version, nil
}
