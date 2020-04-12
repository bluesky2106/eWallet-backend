package servers

import (
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

func (e *EmailSrv) isValidEmailInfo(data *pb.EmailInfo) bool {
	if data == nil {
		return false
	}

	return true
}
