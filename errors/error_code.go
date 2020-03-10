package errors

type (
	// ErrorCode : error code
	ErrorCode int32
)

const (
	// Prefix EC = ErrorCode

	// ECOK is returned on success.
	// Error code starts from 100 in order not to duplicate grpc error codes which starts from 0 to 15.
	ECOK ErrorCode = iota + 100
	// ECCanceled indicates the operation was canceled (typically by the caller).
	ECCanceled
	// ECUnknown error. An example of where this error may be returned is
	// if a Status value received from another address space belongs to
	// an error-space that is not known in this address space. Also
	// errors raised by APIs that do not return enough error information
	// may be converted to this error.
	ECUnknown
	// ECDeadlineExceeded means operation expired before completion.
	// For operations that change the state of the system, this error may be
	// returned even if the operation has completed successfully. For
	// example, a successful response from a server could have been delayed
	// long enough for the deadline to expire.
	ECDeadlineExceeded

	// ECInvalidEmail : invalid email
	ECInvalidEmail
	// ECInvalidPassword : invalid password
	ECInvalidPassword
	// ECEmailNotExists : email doesn't exist
	ECEmailNotExists
	// ECEmailAlreadyExists : email already exists
	ECEmailAlreadyExists
	// ECPasswordMismatch : password and confirm password must match
	ECPasswordMismatch
	// ECInvalidCredentials : invalid credentials
	ECInvalidCredentials
	// ECInvalidName : invalid name
	ECInvalidName
	// ECUpdateUser : cannot update user profile
	ECUpdateUser
	// ECChangePasswordSame : new password and old password are the same
	ECChangePasswordSame
	// ECChangePasswordOldPwdNotSame : old password doesn't match
	ECChangePasswordOldPwdNotSame
	// ECForgotPasswordCode : invalid forgot password code
	ECForgotPasswordCode

	// ECAddressExisted : error message user address is existed
	ECAddressExisted
	// ECAddressSaveFail : user address save fail
	ECAddressSaveFail

	// ECInvalidMessage : invalid message
	ECInvalidMessage
	// ECInvalidArgument : invalid argument
	ECInvalidArgument
	// ECInternalServerError : internal server error
	ECInternalServerError
	// ECInvalidLimit : invalid limit
	ECInvalidLimit
	// ECInvalidPage : invalid page
	ECInvalidPage
	// ECSystemError : system error
	ECSystemError
	// ECPermissionDenied : permission denied
	ECPermissionDenied

	// ECWalletNotMaster : not a master wallet
	ECWalletNotMaster
	// ECWalletNotFound : not found wallet
	ECWalletNotFound
	// ECWalletAddressDataInvalid : wallet address data invalid
	ECWalletAddressDataInvalid

	// ECAssetEmptyRawTx : invalid raw transaction
	ECAssetEmptyRawTx

	// ECTnxExisted : tnx is existed
	ECTnxExisted
	// ECInvalidTnxHash : tnxHash is invalid
	ECInvalidTnxHash
	// ECAssetNotExist : asset is not exist
	ECAssetNotExist
	// ECAssetIDRequire : idAsset is require
	ECAssetIDRequire
	// ECInvalidID : id is invalid
	ECInvalidID

	// ECEventIDInvalid : event id is invalid
	ECEventIDInvalid
	// ECEventIDNotFound : event id is not found
	ECEventIDNotFound
	// ECEventTypeInvalid : eventType is invalid
	ECEventTypeInvalid
	// ECEventTypeNotFound : eventType is not found
	ECEventTypeNotFound
	// ECEventBeginDateInvalid : beginDate is invalid
	ECEventBeginDateInvalid
	// ECEventEndDateInvalid : endDate is invalid
	ECEventEndDateInvalid
	// ECEventEndDateLessThanBeginDate : endDate is less than beginDate
	ECEventEndDateLessThanBeginDate
	// ECEventIndexInvalid : event index is invalid
	ECEventIndexInvalid

	// ECInvalidAdminKey : admin key is invalid
	ECInvalidAdminKey

	// ECInvalidNotifyID : id is required
	ECInvalidNotifyID

	// ECMongoConnection : mongodb connection error
	ECMongoConnection
	// ECMongoCreate : mongodb create model error
	ECMongoCreate
	// ECMongoRead : mongodb read model error
	ECMongoRead
	// ECMongoUpdate : mongodb update model error
	ECMongoUpdate
	// ECMongoDelete : mongodb delete model error
	ECMongoDelete
)
