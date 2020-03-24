package errors

// Prefix EM = ErrorMessage

// common errors
const (
	// EMOK : successful
	EMOK string = "no error"
	// EMCanceled : operation was canceled
	EMCanceled string = "operation was canceled"
	// EMUnknown : Unknown error
	EMUnknown string = "unknown error"
	// EMDeadlineExceeded : operation expired
	EMDeadlineExceeded string = "operation expired"
)

// user errors
const (
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
)

// server errors
const (
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

	// EMInvalidAdminKey : admin key is invalid
	EMInvalidAdminKey string = "admin key is invalid"
)

// wallet errors
const (
	// EMWalletNotMaster : not a master wallet
	EMWalletNotMaster string = "not a master wallet"
	// EMWalletNotFound : not found wallet
	EMWalletNotFound string = "not found wallet"
	// EMWalletAddressDataInvalid : wallet address data invalid
	EMWalletAddressDataInvalid string = "wallet address data invalid"

	// EMAddressExisted : error message user address is existed
	EMAddressExisted string = "user address is existed"
	// EMAddressSaveFail : user address save fail
	EMAddressSaveFail string = "user address save fail"
)

// asset errors
const (
	// EMAssetEmptyRawTx : invalid raw transaction
	EMAssetEmptyRawTx string = "invalid raw transaction"
)

// tnx errors
const (
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
)

// event errors
const (
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
)

// notification errors
const (
	// EMInvalidNotifyID : id is required
	EMInvalidNotifyID string = "id is required"
)

// third - party errors
const (
	// EMRedisConnection : redis connection error
	EMRedisConnection string = "redis connection error"

	// EMRabbitmqConnection : rabbitmq connection error
	EMRabbitmqConnection string = "rabbitmq connection error"

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

	// EMMySQLConnection : mysql connection error
	EMMySQLConnection string = "mysql connection error"
	// EMMySQLDBEmpty : mysql db is empty
	EMMySQLDBEmpty string = "mysql db is empty error"
	// EMMySQLDBAutoMigrate : mysql db auto migrate error
	EMMySQLDBAutoMigrate string = "mysql db auto migrate error"
	// EMMySQLCreate : mysql create model error
	EMMySQLCreate string = "mysql create model error"
	// EMMySQLRead : mysql read model error
	EMMySQLRead string = "mysql read model error"
	// EMMySQLUpdate : mysql update model error
	EMMySQLUpdate string = "mysql update model error"
	// EMMySQLDelete : mysql delete model error
	EMMySQLDelete string = "mysql delete model error"
)

var defaultErrors map[ErrorCode]string

func init() {
	defaultErrors = make(map[ErrorCode]string, 0)

	initDefaultCommonErrors()
	initDefaultUserErrors()
	initDefaultServerErrors()
	initDefaultWalletErrors()
	initDefaultAssetErrors()
	initDefaultTnxErrors()
	initDefaultEventErrors()
	initDefaultNotificationErrors()
	initDefaultThirdPartyErrors()
}

func initDefaultCommonErrors() {
	defaultErrors[ECOK] = EMOK
	defaultErrors[ECCanceled] = EMCanceled
	defaultErrors[ECUnknown] = EMUnknown
	defaultErrors[ECDeadlineExceeded] = EMDeadlineExceeded
}

func initDefaultUserErrors() {
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
}

func initDefaultServerErrors() {
	defaultErrors[ECInvalidMessage] = EMInvalidMessage
	defaultErrors[ECInvalidArgument] = EMInvalidArgument
	defaultErrors[ECInternalServerError] = EMInternalServerError
	defaultErrors[ECInvalidLimit] = EMInvalidLimit
	defaultErrors[ECInvalidPage] = EMInvalidPage
	defaultErrors[ECSystemError] = EMSystemError
	defaultErrors[ECPermissionDenied] = EMPermissionDenied

	defaultErrors[ECInvalidAdminKey] = EMInvalidAdminKey
}

func initDefaultWalletErrors() {
	defaultErrors[ECWalletNotMaster] = EMWalletNotMaster
	defaultErrors[ECWalletNotFound] = EMWalletNotFound
	defaultErrors[ECWalletAddressDataInvalid] = EMWalletAddressDataInvalid
	defaultErrors[ECAddressExisted] = EMAddressExisted
	defaultErrors[ECAddressSaveFail] = EMAddressSaveFail
}

func initDefaultAssetErrors() {
	defaultErrors[ECAssetEmptyRawTx] = EMAssetEmptyRawTx
}

func initDefaultTnxErrors() {
	defaultErrors[ECTnxExisted] = EMTnxExisted
	defaultErrors[ECInvalidTnxHash] = EMInvalidTnxHash
	defaultErrors[ECAssetNotExist] = EMAssetNotExist
	defaultErrors[ECAssetIDRequire] = EMAssetIDRequire
	defaultErrors[ECInvalidID] = EMInvalidID
}

func initDefaultEventErrors() {
	defaultErrors[ECEventIDInvalid] = EMEventIDInvalid
	defaultErrors[ECEventIDNotFound] = EMEventIDNotFound
	defaultErrors[ECEventTypeInvalid] = EMEventTypeInvalid
	defaultErrors[ECEventTypeNotFound] = EMEventTypeNotFound
	defaultErrors[ECEventBeginDateInvalid] = EMEventBeginDateInvalid
	defaultErrors[ECEventEndDateInvalid] = EMEventEndDateInvalid
	defaultErrors[ECEventEndDateLessThanBeginDate] = EMEventEndDateLessThanBeginDate
	defaultErrors[ECEventIndexInvalid] = EMEventIndexInvalid
}

func initDefaultNotificationErrors() {
	defaultErrors[ECInvalidNotifyID] = EMInvalidNotifyID
}

func initDefaultThirdPartyErrors() {
	defaultErrors[ECRedisConnection] = EMRedisConnection
	defaultErrors[ECRabbitmqConnection] = EMRabbitmqConnection

	defaultErrors[ECMongoConnection] = EMMongoConnection
	defaultErrors[ECMongoCreate] = EMMongoCreate
	defaultErrors[ECMongoRead] = EMMongoRead
	defaultErrors[ECMongoUpdate] = EMMongoUpdate
	defaultErrors[ECMongoDelete] = EMMongoDelete

	defaultErrors[ECMySQLConnection] = EMMySQLConnection
	defaultErrors[ECMySQLDBEmpty] = EMMySQLDBEmpty
	defaultErrors[ECMySQLDBAutoMigrate] = EMMySQLDBAutoMigrate
	defaultErrors[ECMySQLCreate] = EMMySQLCreate
	defaultErrors[ECMySQLRead] = EMMySQLRead
	defaultErrors[ECMySQLUpdate] = EMMySQLUpdate
	defaultErrors[ECMySQLDelete] = EMMySQLDelete
}
