package services

import (
	"encoding/json"

	"strings"

	"docker-m/vos"

	"docker-m/utils"
)

func GetImages() ([]vos.ImagesVo, error) {
	address := "/images/json"

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
