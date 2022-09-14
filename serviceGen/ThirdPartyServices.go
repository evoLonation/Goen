package serviceGen

import (
	"Cocome/entityRepo"
	"time"
)

var ThirdPartyServicesInstance ThirdPartyServices

type ThirdPartyServices struct {
}

func (p *ThirdPartyServices) thirdPartyCardPaymentService(cardAccountNumber string, expiryDate time.Time, fee float64) (ret OperationResult[bool]) {
	defer func() {
		if err := entityRepo.Saver.Save(); err != nil {
			ret.Err = NewErrPostCondition(err)
			return
		}
	}()

	// precondition
	if !(true) {
		ret.Err = ErrPreConditionUnsatisfied
		return
	}
	ret.Value = true

	return
}
