package servers

import (
	"math/rand"

	"github.com/bluesky2106/eWallet-backend/entry_store/models"
	pb "github.com/bluesky2106/eWallet-backend/protobuf"
)

const (
	keystoreLength   = 32
	letters          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	masterWalletName = "Master Wallet"
)

func (u *UserSrv) generateUserKeystore() string {
	keystore := u.randomKeystore()
	for {
		count, err := u.dao.CountByQuery(models.User{}, map[string]interface{}{
			"keystore = ?": keystore,
		})
		if err == nil && count == 0 {
			return keystore
		}
		keystore = u.randomKeystore()
	}

}

func (u *UserSrv) randomKeystore() string {
	length := keystoreLength
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (u *UserSrv) isValidUserRequest(req *pb.BaseReq) bool {
	if req.Action == pb.Action_ACTION_READ &&
		req.ObjectType == pb.Object_OBJECT_USER &&
		(req.Message == pb.Message_MESSAGE_READ_USER_BY_EMAIL || req.Message == pb.Message_MESSAGE_READ_USER_BY_ID) {
		return true
	}

	return false
}
