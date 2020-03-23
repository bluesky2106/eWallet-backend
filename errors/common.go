package errors

// GRPCDialError : returns grpc dial error
func GRPCDialError(err error) error {
	if err == nil {
		return nil
	}
	err = New(ECSystemError, err.Error())
	return WithMessage(err, "grpc.Dial")
}

// BCRYPTGenerateFromPasswordError : bcrypt.GenerateFromPassword
func BCRYPTGenerateFromPasswordError(err error) error {
	if err == nil {
		return nil
	}

	err = New(ECSystemError, err.Error())
	return WithMessage(err, "bcrypt.GenerateFromPassword")
}
