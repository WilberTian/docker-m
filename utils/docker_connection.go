package utils

import (
	"net"

	"net/http"

	"net/url"

	"net/http/httputil"

	"io"

	"docker-m/errors"
)

type DockerConnectionUtil struct {
	conn net.Conn
}

func (cUtil *DockerConnectionUtil) InitConnection() error {
	DOCKER_UNIX_SOCKET := "unix:///var/run/docker.sock"
	u, err := url.Parse(DOCKER_UNIX_SOCKET)

	conn, err := net.Dial("unix", u.Path)
	if err != nil {
		return &errors.UnixConenctionError{u.Path, nil}
	}

	cUtil.conn = conn

	return nil
}

func (cUtil *DockerConnectionUtil) Do(method string, address string, reader io.Reader) (*http.Response, error) {
	request, err := http.NewRequest(method, address, reader)
	if err != nil {
		return nil, err
	}

	client := httputil.NewClientConn(cUtil.conn, nil)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	return response, nil
}
