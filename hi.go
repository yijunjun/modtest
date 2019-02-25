package modtest

import "fmt"

func Hi(greeting string)string{
	return fmt.Sprintf("echo hi %v", greeting)
}