package spark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"errors"
)

const SparkApi = "https://api.ciscospark.com/v1"
const SparkRoomsApi = SparkApi + "/rooms"
const SparkMessagesApi = SparkApi + "/messages"


type Client struct {
	accessToken string
}

// NewClient creates and returns a Client instance.
func NewClient(accessToken string) *Client {
	return &Client{accessToken: accessToken}
}

// Post a message to spark
func (c *Client) PostMessage(msg *Message) error {
	body, _ := json.Marshal(msg)
	buf := bytes.NewReader(body)

	req, err := http.NewRequest("POST", SparkMessagesApi, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d - %s", resp.StatusCode, string(t))
	}
	return nil
}

// Post a message with attachment to spark
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

	req, err := http.NewRequest("POST", SparkMessagesApi, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("%d - %s", resp.StatusCode, string(t))
	}

	return err
}

// Find a spark room id has the specific room name
func (c *Client) FindRoomIdByName(roomName string) (string, error) {
	req, err := http.NewRequest("GET", SparkRoomsApi, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+c.accessToken)

	var h http.Client
	resp, err := h.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		t, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("%d - %s", resp.StatusCode, string(t))
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
	msg := fmt.Sprintf("spark room '%s' does not exist", roomName)
	println(msg)
	return "", errors.New(msg)
}
