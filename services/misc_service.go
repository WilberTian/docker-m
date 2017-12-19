package services

import (
	"encoding/json"

	Err "errors"

	"docker-m/vos"

	"docker-m/utils"
)

func GetVersion() (interface{}, error) {
	address := "/version"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var version interface{}
	if err := json.NewDecoder(response.Body).Decode(&version); err != nil {
		return nil, err
	}

	return version, nil
}

func GetInfo() (*vos.DockerInfo, error) {
	address := "/info"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var dockerInfo vos.DockerInfo
	if err := json.NewDecoder(response.Body).Decode(&dockerInfo); err != nil {
		return nil, err
	}

	return &dockerInfo, nil
}

func Ping() (error) {
	address := "/_ping"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 500:
		return Err.New("Server error")
	}

	return nil
}