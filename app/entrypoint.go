package app

import (
	"project/app/infrastructures/echo"
	"project/app/infrastructures/viper"
)

// StartApp entoru point of application
func StartApp() {
	// configuration
	viper.SetupAppConfig()
	echo.StartEcho()
}
