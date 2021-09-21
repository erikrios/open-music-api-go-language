package validation

import "gopkg.in/validator.v2"

func Init() {
	validator.SetPrintJSON(true)
}
