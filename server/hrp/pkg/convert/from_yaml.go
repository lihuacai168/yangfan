package convert

import (
	"reflect"

	"github.com/pkg/errors"

	"github.com/test-instructor/yangfan/server/hrp"
	"github.com/test-instructor/yangfan/server/hrp/internal/builtin"
)

func NewYAMLCase(path string) (*hrp.TCase, error) {
	// load yaml case file
	caseJSON := new(hrp.TCase)
	err := builtin.LoadFile(path, caseJSON)
	if err != nil {
		return nil, errors.Wrap(err, "load yaml file failed")
	}
	if reflect.ValueOf(*caseJSON).IsZero() {
		return nil, errors.New("invalid yaml file")
	}

	err = caseJSON.MakeCompat()
	if err != nil {
		return nil, err
	}
	return caseJSON, nil
}
