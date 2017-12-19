package services

import (
	"encoding/json"

	Err "errors"

	"docker-m/vos"

	"docker-m/utils"

	"docker-m/errors"
)

func GetContainers(params string) ([]vos.ContainersVo, error) {
	address := "/containers/json" + "?" + params

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var containers []vos.ContainersVo
	if err := json.NewDecoder(response.Body).Decode(&containers); err != nil {
		return nil, err
	}

	return containers, nil
}

func GetContainerDetail(id string) (*vos.ContainerDetailVo, error) {
	address := "/containers/" + id + "/json"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var containerDetail vos.ContainerDetailVo
	if err := json.NewDecoder(response.Body).Decode(&containerDetail); err != nil {
		return nil, err
	}

	return &containerDetail, nil
}

func GetContainerProcesses(id string) (*vos.ContainerProcesses, error) {
	address := "/containers/" + id + "/top"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var containerProcesses vos.ContainerProcesses
	if err := json.NewDecoder(response.Body).Decode(&containerProcesses); err != nil {
		return nil, err
	}

	return &containerProcesses, nil
}

func StartContainer(id string) error {
	address := "/containers/" + id + "/start"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return err
	}

	response, err := connUtil.Do("POST", address, nil)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 304:
		return &errors.AlreadyStartedError{ID: id}
	case 404:
		return &errors.NoSuchContainerError{ID: id}
	case 500:
		return Err.New("Server error")
	}

	return nil
}

func StopContainer(id string) error {
	address := "/containers/" + id + "/stop"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return err
	}

	response, err := connUtil.Do("POST", address, nil)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 304:
		return &errors.AlreadyStoppedError{ID: id}
	case 404:
		return &errors.NoSuchContainerError{ID: id}
	case 500:
		return Err.New("Server error")
	}

	return nil
}

func DeleteContainer(id string, params string) error {
	address := "/containers/" + id + "?" + params

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return err
	}

	response, err := connUtil.Do("DELETE", address, nil)
	defer response.Body.Close()

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 400:
		return &errors.DeleteParamError{ID: id}
	case 404:
		return &errors.NoSuchContainerError{ID: id}
	case 409:
		return &errors.DeleteConflictError{ID: id}
	case 500:
		return Err.New("Server error")
	}

	return nil
}

func GetContainerStats(id string, params string) (*vos.ContainerStats, error) {
	address := "/containers/" + id + "/stats" + "?" + params

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var containerStats vos.ContainerStats
	if err := json.NewDecoder(response.Body).Decode(&containerStats); err != nil {
		return nil, err
	}

	return &containerStats, nil
}
