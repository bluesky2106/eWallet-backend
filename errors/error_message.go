package errors

const (
	// Prefix EM = ErrorMessage

	// EMOK : successful
	EMOK string = "no error"
	// EMCanceled : operation was canceled
	EMCanceled string = "operation was canceled"
	// EMUnknown : Unknown error
	EMUnknown string = "unknown error"
	// EMDeadlineExceeded : operation expired
	EMDeadlineExceeded string = "operation expired"

	// EMInvalidEmail : error message invalid email
	EMInvalidEmail string = "invalid email"
	// EMInvalidPassword : error message invalid password
	EMInvalidPassword string = "invalid password"
	// EMEmailNotExists : error message email doesn't exist
	EMEmailNotExists string = "email doesn't exist"
	// EMEmailAlreadyExists : error message email already exists
	EMEmailAlreadyExists string = "email already exists"
	// EMPasswordMismatch : perror message assword and confirm password must match
	EMPasswordMismatch string = "password and confirm password must match"
	// EMInvalidCredentials : error message invalid credentials
	EMInvalidCredentials string = "invalid credentials"
	// EMInvalidName : error message invalid name
	EMInvalidName string = "invalid name"
	// EMUpdateUser : error message cannot update user profile
	EMUpdateUser string = "cannot update user profile"
	// EMChangePasswordSame : new password and old password are the same
	EMChangePasswordSame string = "new password and old password are the same"
	// EMChangePasswordOldPwdNotSame : old password doesn't match
	EMChangePasswordOldPwdNotSame string = "old password doesn't match"
	// EMForgotPasswordCode : invalid forgot password code
	EMForgotPasswordCode string = "invalid forgot password code"

	// EMAddressExisted : error message user address is existed
	EMAddressExisted string = "user address is existed"
	// EMAddressSaveFail : user address save fail
	EMAddressSaveFail string = "user address save fail"

	// EMInvalidMessage : invalid message
	EMInvalidMessage string = "invalid message"
	// EMInvalidArgument : invalid argument
	EMInvalidArgument string = "invalid argument"
	// EMInternalServerError : internal server error
	EMInternalServerError string = "internal server error"
	// EMInvalidLimit : invalid limit
	EMInvalidLimit string = "invalid limit"
	// EMInvalidPage : invalid page
	EMInvalidPage string = "invalid page"
	// EMSystemError : system error
	EMSystemError string = "system error"
	// EMPermissionDenied : permission denied
	EMPermissionDenied string = "permission denied"

	// EMWalletNotMaster : not a master wallet
	EMWalletNotMaster string = "not a master wallet"
	// EMWalletNotFound : not found wallet
	EMWalletNotFound string = "not found wallet"
	// EMWalletAddressDataInvalid : wallet address data invalid
	EMWalletAddressDataInvalid string = "wallet address data invalid"

	// EMAssetEmptyRawTx : invalid raw transaction
	EMAssetEmptyRawTx string = "invalid raw transaction"

	// EMTnxExisted : tnx is existed
	EMTnxExisted string = "tnx is existed"
	// EMInvalidTnxHash : tnxHash is invalid
	EMInvalidTnxHash string = "tnxHash is invalid"
	// EMAssetNotExist : asset is not exist
	EMAssetNotExist string = "asset is not exist"
	// EMAssetIDRequire : idAsset is required
	EMAssetIDRequire string = "idAsset is required"
	// EMInvalidID : id is invalid
	EMInvalidID string = "id is invalid"

	// EMEventIDInvalid : id is invalid
	EMEventIDInvalid string = "id is invalid"
	// EMEventIDNotFound : id is not found
	EMEventIDNotFound string = "id is not found"
	// EMEventTypeInvalid : eventType is invalid
	EMEventTypeInvalid string = "eventType is invalid"
	// EMEventTypeNotFound : eventType is not found
	EMEventTypeNotFound string = "eventType is not found"
	// EMEventBeginDateInvalid : beginDate is invalid
	EMEventBeginDateInvalid string = "beginDate is invalid"
	// EMEventEndDateInvalid : endDate is invalid
	EMEventEndDateInvalid string = "endDate is invalid"
	// EMEventEndDateLessThanBeginDate : endDate is less than beginDate
	EMEventEndDateLessThanBeginDate string = "endDate is less than beginDate"
	// EMEventIndexInvalid : index is invalid
	EMEventIndexInvalid string = "index is invalid"

	// EMInvalidAdminKey : admin key is invalid
	EMInvalidAdminKey string = "admin key is invalid"

	// EMInvalidNotifyID : id is required
	EMInvalidNotifyID string = "id is required"

	// EMMongoConnection : mongodb connection error
	EMMongoConnection string = "mongodb connection error"
	// EMMongoCreate : mongodb create model error
	EMMongoCreate string = "mongodb create model error"
	// EMMongoRead : mongodb read model error
	EMMongoRead string = "mongodb read model error"
	// EMMongoUpdate : mongodb update model error
	EMMongoUpdate string = "mongodb update model error"
	// EMMongoDelete : mongodb delete model error
	EMMongoDelete string = "mongodb delete model error"
)

var defaultErrors map[ErrorCode]string

func init() {
	defaultErrors = make(map[ErrorCode]string, 0)

	defaultErrors[ECOK] = EMOK
	defaultErrors[ECCanceled] = EMCanceled
	defaultErrors[ECUnknown] = EMUnknown
	defaultErrors[ECDeadlineExceeded] = EMDeadlineExceeded

	defaultErrors[ECInvalidEmail] = EMInvalidEmail
	defaultErrors[ECInvalidPassword] = EMInvalidPassword
	defaultErrors[ECEmailNotExists] = EMEmailNotExists
	defaultErrors[ECEmailAlreadyExists] = EMEmailAlreadyExists
	defaultErrors[ECPasswordMismatch] = EMPasswordMismatch
	defaultErrors[ECInvalidCredentials] = EMInvalidCredentials
	defaultErrors[ECInvalidName] = EMInvalidName
	defaultErrors[ECUpdateUser] = EMUpdateUser
	defaultErrors[ECChangePasswordSame] = EMChangePasswordSame
	defaultErrors[ECChangePasswordOldPwdNotSame] = EMChangePasswordOldPwdNotSame
	defaultErrors[ECForgotPasswordCode] = EMForgotPasswordCode

	defaultErrors[ECAddressExisted] = EMAddressExisted
	defaultErrors[ECAddressSaveFail] = EMAddressSaveFail

	defaultErrors[ECInvalidMessage] = EMInvalidMessage
	defaultErrors[ECInvalidArgument] = EMInvalidArgument
	defaultErrors[ECInternalServerError] = EMInternalServerError
	defaultErrors[ECInvalidLimit] = EMInvalidLimit
	defaultErrors[ECInvalidPage] = EMInvalidPage
	defaultErrors[ECSystemError] = EMSystemError
	defaultErrors[ECPermissionDenied] = EMPermissionDenied

	defaultErrors[ECWalletNotMaster] = EMWalletNotMaster
	defaultErrors[ECWalletNotFound] = EMWalletNotFound
	defaultErrors[ECWalletAddressDataInvalid] = EMWalletAddressDataInvalid

	defaultErrors[ECAssetEmptyRawTx] = EMAssetEmptyRawTx

	defaultErrors[ECTnxExisted] = EMTnxExisted
	defaultErrors[ECInvalidTnxHash] = EMInvalidTnxHash
	defaultErrors[ECAssetNotExist] = EMAssetNotExist
	defaultErrors[ECAssetIDRequire] = EMAssetIDRequire
	defaultErrors[ECInvalidID] = EMInvalidID

	defaultErrors[ECEventIDInvalid] = EMEventIDInvalid
	defaultErrors[ECEventIDNotFound] = EMEventIDNotFound
	defaultErrors[ECEventTypeInvalid] = EMEventTypeInvalid
	defaultErrors[ECEventTypeNotFound] = EMEventTypeNotFound
	defaultErrors[ECEventBeginDateInvalid] = EMEventBeginDateInvalid
	defaultErrors[ECEventEndDateInvalid] = EMEventEndDateInvalid
	defaultErrors[ECEventEndDateLessThanBeginDate] = EMEventEndDateLessThanBeginDate
	defaultErrors[ECEventIndexInvalid] = EMEventIndexInvalid

	defaultErrors[ECInvalidAdminKey] = EMInvalidAdminKey

	defaultErrors[ECInvalidNotifyID] = EMInvalidNotifyID

	defaultErrors[ECMongoConnection] = EMMongoConnection
	defaultErrors[ECMongoCreate] = EMMongoCreate
	defaultErrors[ECMongoRead] = EMMongoRead
	defaultErrors[ECMongoUpdate] = EMMongoUpdate
	defaultErrors[ECMongoDelete] = EMMongoDelete
}
