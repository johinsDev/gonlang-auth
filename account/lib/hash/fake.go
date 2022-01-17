package hash

type FakeConfig struct {
	BaseConfig
}

type Fake struct {
	Config *FakeConfig
}

func (hassher *Fake) Make(value string) (string, error) {
	return value, nil
}

func (hasher *Fake) Verify(hashedValue string, plainValue string) (bool, error) {
	return hashedValue == plainValue, nil
}

func (hasher *Fake) NeedReHash(hashedValue string) (bool, error) {
	return false, nil
}

func NewFake(config *FakeConfig) *Fake {
	return &Fake{
		Config: &FakeConfig{
			BaseConfig{
				Driver: "fake",
			},
		},
	}
}
