package project

import (
	"fmt"
	"io"
)

func renderTemplate(rw io.ReadWriteCloser) error {
	goconstructTmplRe.FindStringSubmatch(`2015-05-27`)
	fmt.Printf("%#v\n", goconstructTmplRe.SubexpNames())

	return nil
}
