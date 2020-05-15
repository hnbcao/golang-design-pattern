package factorymethod

import "github.com/sirupsen/logrus"

type Human interface {
	Color()
	Talk()
}

type BlackHuman struct {
	color string
}

func (b BlackHuman) Color() {
	logrus.Infoln("color is black")
}

func (b BlackHuman) Talk() {
	logrus.Infoln("i am black")
}

type YellowHuman struct {
	color string
}

func (y YellowHuman) Color() {
	logrus.Infoln("color is yellow")
}

func (y YellowHuman) Talk() {
	logrus.Infoln("i am yellow")
}

type HumanFactory interface {
	CreateHuman() Human
}

type BlackHumanFactory struct {
}

func (BlackHumanFactory) CreateHuman() Human {
	return &BlackHuman{
		color: "black",
	}
}

type YellowHumanFactory struct {
}

func (YellowHumanFactory) CreateHuman() Human {
	return &YellowHuman{
		color: "yellow",
	}
}
