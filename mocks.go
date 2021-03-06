package websocket

import (
	"net/http"
	"time"

	"github.com/stretchr/testify/mock"
)

//RawConnectionMock is the websocket connection mock
type RawConnectionMock struct {
	mock.Mock
}

//Close is a mocked method
func (c *RawConnectionMock) Close() error {
	args := c.Called()
	return args.Error(0)
}

//PingHandler is a mocked method
func (c *RawConnectionMock) PingHandler() func(appData string) error {
	args := c.Called()
	return args.Get(0).(func(appData string) error)
}

//PongHandler is a mocked method
func (c *RawConnectionMock) PongHandler() func(appData string) error {
	args := c.Called()
	return args.Get(0).(func(appData string) error)
}

//ReadJSON is a mocked method
func (c *RawConnectionMock) ReadJSON(v interface{}) error {
	args := c.Called(v)
	return args.Error(0)
}

//ReadMessage is a mocked method
func (c *RawConnectionMock) ReadMessage() (messageType int, p []byte, err error) {
	args := c.Called()
	return args.Int(0), args.Get(1).([]byte), args.Error(2)
}

//SetPingHandler is a mocked method
func (c *RawConnectionMock) SetPingHandler(h func(appData string) error) {
	c.Called(h)
}

//SetPongHandler is a mocked method
func (c *RawConnectionMock) SetPongHandler(h func(appData string) error) {
	c.Called(h)
}

//SetCloseHandler is a mocked method
func (c *RawConnectionMock) SetCloseHandler(h func(code int, text string) error) {
	c.Called(h)
}

//SetReadDeadline is a mocked method
func (c *RawConnectionMock) SetReadDeadline(t time.Time) error {
	args := c.Called(t)
	return args.Error(0)
}

//SetReadLimit is a mocked method
func (c *RawConnectionMock) SetReadLimit(limit int64) {
	c.Called(limit)
}

//SetWriteDeadline is a mocked method
func (c *RawConnectionMock) SetWriteDeadline(t time.Time) error {
	args := c.Called(t)
	return args.Error(0)
}

//WriteMessage is a mocked method
func (c *RawConnectionMock) WriteMessage(messageType int, data []byte) error {
	args := c.Called(messageType, data)
	return args.Error(0)
}

//ListenerMock is a ConnListener mock
type ListenerMock struct {
	mock.Mock
}

//Handle is a mocked method
func (m *ListenerMock) Handle(msg []byte) {
	m.Called(msg)
}

//Supports is a mocked method
func (m *ListenerMock) Supports(t int) bool {
	args := m.Called(t)
	return args.Bool(0)
}

//FactoryMock is a connection factory mock
type FactoryMock struct {
	mock.Mock
}

//UpgradeConnection is a mocked method
func (u *FactoryMock) UpgradeConnection(writer http.ResponseWriter, req *http.Request, channels []string) (*Connection, error) {
	args := u.Called(writer, req, channels)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Connection), args.Error(1)
}

//HubMock is the Hub interface mock
type HubMock struct {
	mock.Mock
}

//Broadcast is a mocked method
func (h *HubMock) Broadcast(msg Message, channel string) {
	h.Called(msg, channel)
}

//RegisterConnection is a mocked method
func (h *HubMock) RegisterConnection(writer http.ResponseWriter, req *http.Request, channels []string) (*Connection, error) {
	args := h.Called(writer, req, channels)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Connection), args.Error(1)
}

//RegisterListener is a mocked method
func (h *HubMock) RegisterListener(channel string, l ConnListener) {
	h.Called(channel, l)
}

//RegisterOnConnectListener is a mocked method
func (h *HubMock) RegisterOnConnectListener(l OnConnectListener) {
	h.Called(l)
}

//RegisterOnChannels is a mocked method
func (h *HubMock) RegisterOnChannels(c *Connection, channels []string) {
	h.Called(c, channels)
}
