package validation

import (
	"errors"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/shopspring/decimal"
)

type DecimalThresholdRule struct {
	threshold decimal.Decimal
	operator  int
	err       ozzo.Error
}

const (
	greaterThan = iota
	greaterEqualThan
	lessThan
	lessEqualThan
)

func MinDecimal(min decimal.Decimal) DecimalThresholdRule {
	return DecimalThresholdRule{
		threshold: min,
		operator:  greaterEqualThan,
		err:       ozzo.ErrMinGreaterEqualThanRequired,
	}

}

func MaxDecimal(max decimal.Decimal) DecimalThresholdRule {
	return DecimalThresholdRule{
		threshold: max,
		operator:  lessEqualThan,
		err:       ozzo.ErrMaxLessEqualThanRequired,
	}
}

func (r DecimalThresholdRule) Exclusive() DecimalThresholdRule {
	if r.operator == greaterEqualThan {
		r.operator = greaterThan
		r.err = ozzo.ErrMinGreaterThanRequired
	} else if r.operator == lessEqualThan {
		r.operator = lessThan
		r.err = ozzo.ErrMaxLessThanRequired
	}
	return r
}

func (r DecimalThresholdRule) Validate(value interface{}) error {
	v, ok := value.(decimal.Decimal)
	if !ok {
		return errors.New("unsupported type")
	}

	if r.compareDecimal(r.threshold, v) {
		return nil
	}

	return r.err.SetParams(map[string]interface{}{"threshold": r.threshold})
}

func (r DecimalThresholdRule) compareDecimal(threshold, value decimal.Decimal) bool {
	switch r.operator {
	case greaterThan:
		return value.GreaterThan(threshold)
	case greaterEqualThan:
		return value.GreaterThanOrEqual(threshold)
	case lessThan:
		return value.LessThan(threshold)
	default:
		return value.LessThanOrEqual(threshold)
	}
}
