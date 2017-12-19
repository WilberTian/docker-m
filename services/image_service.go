package services

import (
	"encoding/json"

	"strings"

	Err "errors"

	"docker-m/vos"

	"docker-m/utils"

	"docker-m/errors"
)

func GetImages() ([]vos.ImagesVo, error) {
	address := "/images"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, strings.NewReader(""))
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var images []vos.ImagesVo
	if err := json.NewDecoder(response.Body).Decode(&images); err != nil {
		return nil, err
	}

	return images, nil
}

func GetImageDetail(id string) (*vos.ImageDetailVo, error) {
	address := "/images/" + id

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var imageDetailVo vos.ImageDetailVo
	if err := json.NewDecoder(response.Body).Decode(&imageDetailVo); err != nil {
		return nil, err
	}

	return &imageDetailVo, nil
}

func GetImageHistory(id string) ([]vos.ImageHistoryVo, error) {
	address := "/images/" + id + "/history"

	connUtil := utils.DockerConnectionUtil{}
	if err := connUtil.InitConnection(); err != nil {
		return nil, err
	}

	response, err := connUtil.Do("GET", address, nil)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	var imageHistoryVo []vos.ImageHistoryVo
	if err := json.NewDecoder(response.Body).Decode(&imageHistoryVo); err != nil {
		return nil, err
	}

	return imageHistoryVo, nil
}

func DeleteImage(id string, params string) error {
	address := "/iamges/" + id + "?" + params

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
	case 404:
		return &errors.NoSuchImageError{ID: id}
	case 409:
		return &errors.DeleteConflictError{ID: id}
	case 500:
		return Err.New("Server error")
	}

	return nil
}