package servers

import (
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

func (u *UserSrv) isValidUserRequest(req *pb.BaseReq) bool {
	if req.GetObjectType() != pb.Object_OBJECT_USER {
		return false
	}

	switch req.GetAction() {
	case pb.Action_ACTION_CREATE:
		if req.GetMessage() == pb.Message_MESSAGE_CREATE_USER {
			return true
		}
	case pb.Action_ACTION_READ:
		if req.Message == pb.Message_MESSAGE_READ_USER_BY_EMAIL ||
			req.Message == pb.Message_MESSAGE_READ_USER_BY_ID {
			return true
		}
	}

	return false
}
