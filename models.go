package main

import "gorm.io/gorm"

type PingPong struct {
	gorm.Model
	Counter int64
}
