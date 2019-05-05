package main

import (
	"github.com/spf13/viper"
	"github.com/mtbdeano/rpi-go-led-sign/cmd"
)

func setup() {
	viper.SetDefault("GPIO_PIN", 24)

}

func main() {
	cmd.Execute()
	
}