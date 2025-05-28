package area

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"money/internal/core"
	mocks "money/internal/repo/mocks/area"
)

func TestCreatePost_TableDriven(t *testing.T) {
	testCases := []struct {
		name             string
		formValues       url.Values
		mockSession      *core.Session
		mockSessionErr   error
		mockInsertAreaFn func(ctx context.Context, area *core.Area) error
		expectedCode     int
		expectedLocation string // optional
	}{
		{
			name: "успешное создание области",
			formValues: url.Values{
				"name":    {"Тестовая область"},
				"levelid": {"1"},
			},
			mockSession: &core.Session{},
			mockInsertAreaFn: func(ctx context.Context, area *core.Area) error {
				return nil
			},
			expectedCode: http.StatusSeeOther,
		},
		{
			name: "некорректный levelid",
			formValues: url.Values{
				"name":    {"Тестовая область"},
				"levelid": {""}, // пустой levelid
			},
			mockSession: &core.Session{},
			mockInsertAreaFn: func(ctx context.Context, area *core.Area) error {
				return nil
			},
			expectedCode: http.StatusTemporaryRedirect,
		},
		{
			name: "levelid не число",
			formValues: url.Values{
				"name":    {"Тестовая область"},
				"levelid": {"abc"},
			},
			mockSession: &core.Session{},
			mockInsertAreaFn: func(ctx context.Context, area *core.Area) error {
				return nil
			},
			expectedCode: http.StatusTemporaryRedirect,
		},
		{
			name: "ошибка вставки в БД",
			formValues: url.Values{
				"name":    {"Тестовая область"},
				"levelid": {"1"},
			},
			mockSession: &core.Session{},
			mockInsertAreaFn: func(ctx context.Context, area *core.Area) error {
				return assert.AnError // произвольная ошибка
			},
			expectedCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockSessionManager := mocks.NewMockSessionManager(ctrl)
			mockAreaStorager := mocks.NewMockAreaStorager(ctrl)

			// Настройка мока сессии
			mockSessionManager.EXPECT().GetSession(gomock.Any()).Return(tc.mockSession, tc.mockSessionErr).AnyTimes()

			// Настройка мока InsertArea
			mockAreaStorager.EXPECT().InsertArea(gomock.Any(), gomock.Any()).DoAndReturn(tc.mockInsertAreaFn).AnyTimes()

			handler := CreatePost(mockSessionManager, mockAreaStorager)

			// Создаём запрос
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.formValues.Encode()))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()

			assert.Equal(t, tc.expectedCode, resp.StatusCode)

			if tc.expectedLocation != "" {
				assert.Equal(t, tc.expectedLocation, resp.Header.Get("Location"))
			}
		})
	}
}
