package notification_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgdraganov/noti-fire/internal/http/handler/notification"
	"github.com/dgdraganov/noti-fire/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type actionMock struct {
	execute func(context.Context, string) error
}

func (a *actionMock) Execute(ctx context.Context, message string) error {
	return a.execute(ctx, message)
}

func Test_NotificationHandler_ServeHTTP_Success(t *testing.T) {
	action := &actionMock{
		execute: func(ctx context.Context, s string) error {
			return nil
		},
	}
	notificationHandler := notification.NewNotificationHandler(http.MethodPost, action, zap.NewNop().Sugar())

	reqMessage := model.NotificationRequest{
		Message: "test message",
	}

	var reqestBody bytes.Buffer
	err := json.NewEncoder(&reqestBody).Encode(reqMessage)
	if err != nil {
		t.Fatal("failed to encode NotificationRequest")
	}

	endpoint := "/notify"
	request, err := http.NewRequest(http.MethodPost, endpoint, &reqestBody)
	if err != nil {
		t.Fatalf("creating new request failed: %s", endpoint)
	}

	requestID := uuid.New()
	ctx := context.WithValue(request.Context(), model.RequestID, requestID.String())
	request = request.WithContext(ctx)

	responseWriter := httptest.NewRecorder()

	notificationHandler.ServeHTTP(responseWriter, request)

	body, err := io.ReadAll(responseWriter.Body)
	if err != nil {
		t.Fatal("failed to read response body")
	}

	respStruct := model.ResponseMessage{
		Message: "notification request submitted successfully",
	}
	expectedBody, err := json.Marshal(respStruct)
	if err != nil {
		t.Fatal("failed to marshal ResponseMessage")
	}

	got := string(body)
	expected := string(expectedBody)

	gotStatusCode := responseWriter.Code
	expectedStatusCode := http.StatusOK

	if string(expectedBody) != string(got) {
		t.Fatalf("response does not match, expected: %s, got: %s", expected, got)
	}
	if expectedStatusCode != gotStatusCode {
		t.Fatalf("response code does not match, expected: %d, got: %d", expectedStatusCode, gotStatusCode)
	}
}

func Test_NotificationHandler_ServeHTTP_WrongMethod(t *testing.T) {

	action := &actionMock{}
	handlerMethod := http.MethodPost
	notificationHandler := notification.NewNotificationHandler(handlerMethod, action, zap.NewNop().Sugar())

	endpoint := "/notify"
	testMethod := http.MethodGet
	request, err := http.NewRequest(testMethod, endpoint, nil)
	if err != nil {
		t.Fatalf("creating new request failed: %s", endpoint)
	}

	requestID := uuid.New()
	ctx := context.WithValue(request.Context(), model.RequestID, requestID.String())
	request = request.WithContext(ctx)

	responseWriter := httptest.NewRecorder()

	notificationHandler.ServeHTTP(responseWriter, request)

	body, err := io.ReadAll(responseWriter.Body)
	if err != nil {
		t.Fatal("failed to read response body")
	}

	respStruct := model.ResponseMessage{
		Message: fmt.Sprintf("invalid request method %s, expected method is %s", testMethod, handlerMethod),
	}
	expectedBody, err := json.Marshal(respStruct)
	if err != nil {
		t.Fatal("failed to marshal ResponseMessage")
	}

	got := string(body)
	expected := string(expectedBody)

	gotStatusCode := responseWriter.Code
	expectedStatusCode := http.StatusBadRequest

	if string(expectedBody) != string(got) {
		t.Fatalf("response does not match, expected: %s, got: %s", expected, got)
	}
	if expectedStatusCode != gotStatusCode {
		t.Fatalf("response code does not match, expected: %d, got: %d", expectedStatusCode, gotStatusCode)
	}
}

func Test_NotificationHandler_ServeHTTP_InvalidRequestBody(t *testing.T) {

	action := &actionMock{}
	notificationHandler := notification.NewNotificationHandler(http.MethodPost, action, zap.NewNop().Sugar())

	reqestBody := strings.NewReader("some invalid JSON")

	endpoint := "/notify"
	request, err := http.NewRequest(http.MethodPost, endpoint, reqestBody)
	if err != nil {
		t.Fatalf("creating new request failed: %s", endpoint)
	}

	requestID := uuid.New()
	ctx := context.WithValue(request.Context(), model.RequestID, requestID.String())
	request = request.WithContext(ctx)

	responseWriter := httptest.NewRecorder()

	notificationHandler.ServeHTTP(responseWriter, request)

	body, err := io.ReadAll(responseWriter.Body)
	if err != nil {
		t.Fatal("failed to read response body")
	}

	respStruct := model.ResponseMessage{
		Message: "invalid JSON request body",
	}
	expectedBody, err := json.Marshal(respStruct)
	if err != nil {
		t.Fatal("failed to marshal ResponseMessage")
	}

	got := string(body)
	expected := string(expectedBody)

	gotStatusCode := responseWriter.Code
	expectedStatusCode := http.StatusBadRequest

	if string(expectedBody) != string(got) {
		t.Fatalf("response does not match, expected: %s, got: %s", expected, got)
	}
	if expectedStatusCode != gotStatusCode {
		t.Fatalf("response code does not match, expected: %d, got: %d", expectedStatusCode, gotStatusCode)
	}
}

func Test_NotificationHandler_ServeHTTP_ActionExecuteFailed(t *testing.T) {

	action := &actionMock{
		execute: func(ctx context.Context, s string) error {
			return errors.New("test error")
		},
	}
	notificationHandler := notification.NewNotificationHandler(http.MethodPost, action, zap.NewNop().Sugar())

	reqMessage := model.NotificationRequest{
		Message: "test message",
	}

	var reqestBody bytes.Buffer
	err := json.NewEncoder(&reqestBody).Encode(reqMessage)
	if err != nil {
		t.Fatal("failed to encode NotificationRequest")
	}

	endpoint := "/notify"
	request, err := http.NewRequest(http.MethodPost, endpoint, &reqestBody)
	if err != nil {
		t.Fatalf("creating new request failed: %s", endpoint)
	}

	requestID := uuid.New()
	ctx := context.WithValue(request.Context(), model.RequestID, requestID.String())
	request = request.WithContext(ctx)

	responseWriter := httptest.NewRecorder()

	notificationHandler.ServeHTTP(responseWriter, request)

	body, err := io.ReadAll(responseWriter.Body)
	if err != nil {
		t.Fatal("failed to read response body")
	}

	respStruct := model.ResponseMessage{
		Message: "Something went wrong on our end!",
	}
	expectedBody, err := json.Marshal(respStruct)
	if err != nil {
		t.Fatal("failed to marshal ResponseMessage")
	}

	got := string(body)
	expected := string(expectedBody)

	gotStatusCode := responseWriter.Code
	expectedStatusCode := http.StatusInternalServerError

	if string(expectedBody) != string(got) {
		t.Fatalf("response does not match, expected: %s, got: %s", expected, got)
	}
	if expectedStatusCode != gotStatusCode {
		t.Fatalf("response code does not match, expected: %d, got: %d", expectedStatusCode, gotStatusCode)
	}
}
