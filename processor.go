package main

type Processor interface {
	Authorize() (string, error)
}
