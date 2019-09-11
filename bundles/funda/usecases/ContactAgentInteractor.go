package usecases

import (
	"github.com/fspace/ecm/bundles/funda/entities"
	"github.com/fspace/ecm/bundles/funda/util"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

// 原始实现中有验证引擎依赖 这里可选：
// - https://github.com/asaskevich/govalidator
// - https://github.com/go-playground/validator/
// - https://github.com/go-ozzo/ozzo-validation  强哥的
// 上面两个通过注解tag来写验证规则 强哥的是从yii中移植过来的  在方法上做手脚 更灵活可控 还能进行特殊验证
// NOTE 真正的验证逻辑应该出现在领域层 属于领域知识

type ContactAgentInteractor struct {
	HouseRepo entities.IHouseRepo
}

func NewContactAgentInteractor(repo entities.IHouseRepo) *ContactAgentInteractor {
	return &ContactAgentInteractor{
		HouseRepo: repo,
	}
}

func (it ContactAgentInteractor) Handle(request ContactAgentRequestMessage) ContactAgentResponseMessage {
	err := request.Validate()
	if err != nil {
		// 验证未通过！
		return ContactAgentResponseMessage{ValidationResult: err}
	}

	h := it.HouseRepo.Get(request.HouseId)
	h.RegisterInterest(entities.Interest{
		CustomerEmailAddress: request.CustomerEmailAddress,
		CustomerPhoneNumber:  request.CustomerPhoneNumber,
		CreationDate:         util.GetTime(),
	})

	it.HouseRepo.Save(h)

	return ContactAgentResponseMessage{
		ValidationResult: err,
		HouseId:          request.HouseId,
	}
}

type ContactAgentRequestMessage struct {
	CustomerEmailAddress string
	CustomerPhoneNumber  string // 原始用long 类型
	HouseId              int64
}

// @see https://github.com/go-ozzo/ozzo-validation#validatable-types
func (c ContactAgentRequestMessage) Validate() error {
	return validation.ValidateStruct(&c,
		//// Street cannot be empty, and the length must between 5 and 50
		//validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		//// City cannot be empty, and the length must between 5 and 50
		//validation.Field(&a.City, validation.Required, validation.Length(5, 50)),
		//// State cannot be empty, and must be a string consisting of two letters in upper case
		//validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		//// State cannot be empty, and must be a string consisting of five digits
		//validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),

		validation.Field(&c.CustomerEmailAddress, validation.Required, validation.Length(5, 50)),
		validation.Field(&c.CustomerEmailAddress, is.Email),
	)
}

type ContactAgentResponseMessage struct {
	ValidationResult error
	HouseId          int64
}
