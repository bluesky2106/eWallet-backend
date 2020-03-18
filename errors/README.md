
# Error handling

### To define a new error code (associated with a new error message), we need to:
- Define error code in `errors/error_code.go`. For example:
    ```
        // ECChangePasswordSame : new password and old password are the same
        ECChangePasswordSame
    ```

- Define error message in `errors/error_message.go`. For example:
    ```
        // EMChangePasswordSame : new password and old password are the same
	    EMChangePasswordSame string = "new password and old password are the same"
    ```

- Link the new error code to the new error message by adding a new key - value pair in `defaultErrors map[ErrorCode]string` in file `errors/error_message.go`:
    ```
        defaultErrors[ECChangePasswordSame] = EMChangePasswordSame
    ```
### Create new error in services (backend, cache, store, etc.)

- Import errors pacakge:

    ```
    import (
        errs "github.com/bluesky2106/eWallet-backend/errors"	
    )
    ```

- When a error <b>first</b> appears in our system (normally, golang standard library functions or third party functions threw this error), "wrap" this error with our `func New(code ErrorCode, messages ...string) error` as shown below:

    ```
    if err := c.ShouldBindJSON(&req); err != nil {
        err := errs.New(errs.ECInvalidArgument)
        ...
    }
    ```

    - In the above sample code, `err := errs.New(errs.ECInvalidArgument)` statement indicates that the default error message `EMInvalidArgument string = "invalid argument"` will be used. 
    - However, if we want to replace the default error message by the string in `err.Error()`, we can do like: `err := errs.New(errs.ECInvalidArgument, err.Error())`.

### Wrap the error message

- We can wrap the error message of original error by using `func WithMessage(err error, messages ...string) error`

    ```
    user, err := s.userSvc.Authenticate(&req)
	if err != nil {
		return nil, errs.WithMessage(err, "s.userSvc.Authenticate")
	}
    ```
