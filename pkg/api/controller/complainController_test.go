package controller

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"projeto-ra-api-go/pkg/api/controller/mocks"
	"projeto-ra-api-go/pkg/api/model"
	"strings"
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/mock"
)

func TestComplainController_SaveComplain(t *testing.T) {

	type fields struct {
		service *mocks.Complain
	}
	type args struct {
		body io.Reader
	}

	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				body: strings.NewReader(`{
				  "title": "Nenhuma atenção com o cliente",
				  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				  "locale": {
						"city":"Salvador",
						"state":"BA"
				  },
				  "company": {
					  	"name": "Enterprise90",
						"description": "Enterprise Developer 90 based"
				  }
			}`),
			},
			wantHttpStatusCode: http.StatusCreated,
			mock: func(fs *mocks.Complain) {
				fs.On("Save", mock.Anything, mock.Anything).Return(mock.Anything, nil).Once()
			},
		},
		{
			name: "return 422 when don't send the body",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				body: strings.NewReader(``),
			},
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			mock: func(fs *mocks.Complain) {
				fs.On("Save", mock.Anything, mock.Anything).Maybe().Times(0)
			},
		},
		{
			name: "return 422 when don't send the fild",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				body: strings.NewReader(`{"name":"", "locale":"mock", "company":"mock"}`),
			},
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			mock: func(fs *mocks.Complain) {
				fs.On("Save", mock.Anything, mock.Anything).Maybe().Times(0)
			},
		},
		{
			name: "return 422 when send the fild with number",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				body: strings.NewReader(`{"name":"", "locale":1, "company":1}`),
			},
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			mock: func(fs *mocks.Complain) {
				fs.On("Save", mock.Anything, mock.Anything).Maybe().Times(0)
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}

			request := httptest.NewRequest(http.MethodPost, "/complains", tt.args.body)
			recorder := httptest.NewRecorder()

			p.SaveComplain(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)

		})
	}
}

func TestComplainController_DeleteById(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		service *mocks.Complain
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{

				id: (id.Hex()),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			}},
		{name: "return 404 not found",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: mock.Anything,
			},
			wantHttpStatusCode: http.StatusNotFound,
			mock: func(fs *mocks.Complain) {
				fs.On("DeleteById", mock.Anything, mock.Anything).Return(errors.New("not found")).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodDelete, "/complains/"+id.Hex(), nil)
			recorder := httptest.NewRecorder()

			p.DeleteById(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestComplainController_GetAll(t *testing.T) {
	var complains []model.ComplainOut
	type fields struct {
		service *mocks.Complain
	}

	tests := []struct {
		name               string
		fields             fields
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("FindByParam", mock.Anything, mock.Anything).Return(complains, nil).Once()
			}},
		{name: "return 500 server error",
			fields: fields{
				service: new(mocks.Complain),
			},

			wantHttpStatusCode: http.StatusInternalServerError,
			mock: func(fs *mocks.Complain) {
				fs.On("FindByParam", mock.Anything, mock.Anything).Return(nil, errors.New("server error")).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/complains", nil)
			recorder := httptest.NewRecorder()

			p.FindByParam(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestComplainController_GetAll2(t *testing.T) {
	var complains []model.ComplainOut
	type fields struct {
		service *mocks.Complain
	}

	tests := []struct {
		name               string
		fields             fields
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("FindByParam", mock.Anything, mock.Anything).Return(complains, nil).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/complains?title=Tatooine&company=Solid", nil)
			recorder := httptest.NewRecorder()

			p.FindByParam(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestComplainController_Update(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		service *mocks.Complain
	}
	type args struct {
		id   string
		body io.Reader
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: (id.Hex()),
				body: strings.NewReader(`{
				  "title": "Nenhuma atenção com o cliente",
				  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				  "locale": {
						"city":"Salvador",
					"state":"BA"
				  },
				  "company": {
					  "name": "Enterprise90",
						"description": "Enterprise Developer 90 based"
				  }
			}`),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			}},
		{name: "return 404 not found",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: id.Hex(),
				body: strings.NewReader(`{
				  "title": "Nenhuma atenção com o cliente",
				  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				  "locale": {
						"city":"Salvador",
					"state":"BA"
				  },
				  "company": {
					  "name": "Enterprise90",
						"description": "Enterprise Developer 90 based"
				  }
			}`),
			},
			wantHttpStatusCode: http.StatusNotFound,
			mock: func(fs *mocks.Complain) {
				fs.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("not found")).Once()
			}},
		{name: "return 422 not unoprocessable entity",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: id.Hex(),
				body: strings.NewReader(`{
				  "title": "Nenhuma atenção com o cliente",
				  "description": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			}`),
			},
			wantHttpStatusCode: http.StatusUnprocessableEntity,
			mock: func(fs *mocks.Complain) {
				fs.On("Update", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("field null")).Maybe().Times(0)
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodPut, "/complains/"+id.Hex(), tt.args.body)
			recorder := httptest.NewRecorder()

			p.Update(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}

}

func TestComplainController_FindById(t *testing.T) {
	id := primitive.NewObjectID()
	complain := &model.ComplainOut{ID: id, Title: "", Description: ""}

	type fields struct {
		service *mocks.Complain
	}
	type args struct {
		id string
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: (id.Hex()),
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("FindById", mock.Anything, mock.Anything).Return(complain, nil).Once()
			}},
		{name: "return 404 not found",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				id: mock.Anything,
			},
			wantHttpStatusCode: http.StatusNotFound,
			mock: func(fs *mocks.Complain) {
				fs.On("FindById", mock.Anything, mock.Anything).Return(nil, errors.New("not found ")).Once()
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.service)
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/complains/"+id.Hex(), nil)
			recorder := httptest.NewRecorder()

			p.FindById(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}

func TestComplainController_FindByParam(t *testing.T) {
	complain := &model.ComplainIn{Title: "", Description: ""}
	type fields struct {
		service *mocks.Complain
	}
	type args struct {
		plan *model.ComplainIn
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantHttpStatusCode int
		mock               func(fs *mocks.Complain)
	}{
		{
			name: "sucesss",
			fields: fields{
				service: new(mocks.Complain),
			},
			args: args{
				plan: complain,
			},
			wantHttpStatusCode: http.StatusOK,
			mock: func(fs *mocks.Complain) {
				fs.On("FindByParam", mock.Anything, complain).Return([]model.ComplainOut{
					model.ComplainOut{
						ID:          primitive.NewObjectID(),
						Title:       mock.Anything,
						Description: mock.Anything,
						Locale: model.Locale{
							City:  mock.Anything,
							State: mock.Anything,
						},
						Company: model.Company{
							Title:       mock.Anything,
							Description: mock.Anything,
						},
						CountPageViews: 0,
						IsOnTop10BadRA: false,
					},
				}, nil).Once()
			}},
	}
	for _, tt := range tests {
		tt.mock(tt.fields.service)
		t.Run(tt.name, func(t *testing.T) {
			p := &ComplainCtrlImpl{
				service: tt.fields.service,
			}
			request := httptest.NewRequest(http.MethodGet, "/complains/?title="+complain.Title, nil)
			recorder := httptest.NewRecorder()

			p.FindByParam(recorder, request)

			assert.Equal(t, tt.wantHttpStatusCode, recorder.Code)

			tt.fields.service.AssertExpectations(t)
		})
	}
}
