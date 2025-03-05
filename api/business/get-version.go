package business

import "github.com/GPA-Gruppo-Progetti-Avanzati-SRL/go-core-app"

func (l Logic) GetVersion() string {
	return core.BuildVersion
}
