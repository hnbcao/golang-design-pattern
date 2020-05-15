package factorymethod

import "testing"

func createHuman(factory HumanFactory) Human {
	return factory.CreateHuman()
}

func TestHumanFactory_CreateHuman(t *testing.T) {
	blackHuman := createHuman(&BlackHumanFactory{})
	blackHuman.Color()
	blackHuman.Talk()
	yellowHuman := createHuman(&YellowHumanFactory{})
	yellowHuman.Color()
	yellowHuman.Talk()
}
