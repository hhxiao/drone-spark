package spark

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const spark_api = "https://api.ciscospark.com/v1"
const spark_rooms_api = spark_api + "/rooms"
const spark_messages_api = spark_api + "/messages"

type Client struct {
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{accessToken: accessToken}
}

func (c *Client) PostMessage(msg *Message) error {
	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	req, err := http.NewRequest("POST", spark_messages_api, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer " + c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, string(t)))
	}
	return nil
}

func (c *Client) PostFileMessage(msg *Message) error {
	file, err := os.Open(msg.Files)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("files", filepath.Base(msg.Files))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	if msg.RoomId != "" {
		writer.WriteField("roomId", msg.RoomId)
	}
	if msg.ToPersonEmail != "" {
		writer.WriteField("toPersonEmail", msg.ToPersonEmail)
	}
	if msg.Text != "" {
		writer.WriteField("text", msg.Text)
	}
	if msg.Markdown != "" {
		writer.WriteField("markdown", msg.Markdown)
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", spark_messages_api, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer " + c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, string(t)))
	}

	return nil
}

func (c *Client) FindRoomIdByName(roomName string) (string, error) {
	req, err := http.NewRequest("GET", spark_rooms_api, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer " + c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		//t, _ := ioutil.ReadAll(resp.Body)
		return "", errors.New(fmt.Sprintf("%d - %s", resp.StatusCode, c.accessToken))
	}

	var rooms Rooms
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(body, &rooms)
	if err != nil {
		return "", err
	}
	for _, room := range rooms.Items {
		if room.Title == roomName {
			return room.Id, nil
		}
	}
	return "", errors.New(fmt.Sprintf("spark room '%s' does not exist", roomName))
}
