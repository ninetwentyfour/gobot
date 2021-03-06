package gobot

import "time"

type Driver struct {
	Adaptor  AdaptorInterface
	Interval time.Duration
	Pin      string
	Name     string
	Commands map[string]func(map[string]interface{}) interface{}
	Events   map[string]*Event
}

type DriverInterface interface {
	Start() bool
	Halt() bool
	adaptor() AdaptorInterface
	setInterval(time.Duration)
	interval() time.Duration
	setName(string)
	name() string
	commands() map[string]func(map[string]interface{}) interface{}
}

func (d *Driver) adaptor() AdaptorInterface {
	return d.Adaptor
}

func (d *Driver) setInterval(t time.Duration) {
	d.Interval = t
}

func (d *Driver) interval() time.Duration {
	return d.Interval
}

func (d *Driver) setName(s string) {
	d.Name = s
}

func (d *Driver) name() string {
	return d.Name
}

func (d *Driver) commands() map[string]func(map[string]interface{}) interface{} {
	return d.Commands
}

func (d *Driver) AddCommand(name string, f func(map[string]interface{}) interface{}) {
	d.Commands[name] = f
}
