package servers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bluesky2106/eWallet-backend/bo_controller/api"
	"github.com/bluesky2106/eWallet-backend/bo_controller/serializers"
	boModels "github.com/bluesky2106/eWallet-backend/bo_entry_store/models"
	"github.com/bluesky2106/eWallet-backend/integration_test/config"
)

// User : struct
type User struct {
	conf *config.Config
}

// NewUser : new user server
func NewUser(conf *config.Config) *User {
	return &User{
		conf: conf,
	}
}

// AdminUserRegister : admin user register
func (uSrv *User) AdminUserRegister(uReq *serializers.UserRegisterReq) (*boModels.User, error) {
	jsonValue, _ := json.Marshal(uReq)
	url := fmt.Sprintf("%s/auth/register", uSrv.conf.URLBO)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Read body failed")
	}

	var respAPI api.Resp
	err = json.Unmarshal(body, &respAPI)
	if err != nil {
		return nil, errors.New("Unmarshal body failed")
	}

	jsonValue, err = json.Marshal(respAPI.Result)
	if err != nil {
		return nil, errors.New("Marshal respAPI.Result failed")
	}

	var adminUser boModels.User
	err = json.Unmarshal(jsonValue, &adminUser)
	if err != nil {
		return nil, errors.New("Unmarshal respAPI.Result failed")
	}

	return &adminUser, nil
}

// AdminUserLogin : admin user login
func (uSrv *User) AdminUserLogin(uReq *serializers.UserLoginReq) (*serializers.UserLoginResp, error) {
	url := fmt.Sprintf("%s/auth/login", uSrv.conf.URLBO)
	jsonValue, _ := json.Marshal(uReq)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Read body failed")
	}

	var respAPI api.Resp
	err = json.Unmarshal(body, &respAPI)
	if err != nil {
		return nil, errors.New("Unmarshal body failed")
	}

	jsonValue, err = json.Marshal(respAPI.Result)
	if err != nil {
		return nil, errors.New("Marshal respAPI.Result failed")
	}

	var respLogin serializers.UserLoginResp
	err = json.Unmarshal(jsonValue, &respLogin)
	if err != nil {
		return nil, errors.New("Unmarshal respAPI.Result failed")
	}

	return &respLogin, nil
}

// AdminUserProfile : admin user profile
func (uSrv *User) AdminUserProfile(token string) (*boModels.User, error) {
	url := fmt.Sprintf("%s/auth/user-profile", uSrv.conf.URLBO)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status != 200")
	}

	defer func(r *http.Response) {
		err := r.Body.Close()
		if err != nil {
			fmt.Println("Close body failed", err.Error())
		}
	}(res)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("Read body failed")
	}

	var respAPI api.Resp
	err = json.Unmarshal(body, &respAPI)
	if err != nil {
		return nil, errors.New("Unmarshal body failed")
	}

	jsonValue, err := json.Marshal(respAPI.Result)
	if err != nil {
		return nil, errors.New("Marshal respAPI.Result failed")
	}

	var adminUser boModels.User
	err = json.Unmarshal(jsonValue, &adminUser)
	if err != nil {
		return nil, errors.New("Unmarshal respAPI.Result failed")
	}

	return &adminUser, nil
}
