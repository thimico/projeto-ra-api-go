package service

import (
	"context"
	"errors"
	"projeto-ra-api-go/pkg/api/model"
	"projeto-ra-api-go/pkg/api/service/mocks"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_complain_Save(t *testing.T) {
	idd := primitive.NewObjectID()
	type fields struct {
		dao    *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx context.Context
		in  *model.ComplainIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ComplainOut
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{
			name: "save sucess",
			fields: fields{
				dao:    new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.ComplainIn{
					Title:       mock.Anything,
					Description: mock.Anything,
					Locale:      model.Locale{
						City:  mock.Anything,
						State: mock.Anything,
					},
					Company:     model.Company{
						Title:       mock.Anything,
						Description: mock.Anything,
					},
				},
			},
			want:    &model.ComplainOut{
			ID:          idd ,
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
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return(&model.Complain{
					ID:      idd,
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
				}, nil).Once()
			},
		},
		{
			name: "save error",
			fields: fields{
				dao:    new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.ComplainIn{
					Title:       mock.Anything,
					Description: mock.Anything,
					Locale:      model.Locale{
						City:  mock.Anything,
						State: mock.Anything,
					},
					Company:     model.Company{
						Title:       mock.Anything,
						Description: mock.Anything,
					},
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save"))
			},
		},
		{
			name: "save error2",
			fields: fields{
				dao:    new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				in: &model.ComplainIn{
					Title:       mock.Anything,
					Description: mock.Anything,
					Locale:      model.Locale{
						City:  mock.Anything,
						State: mock.Anything,
					},
					Company:     model.Company{
						Title:       mock.Anything,
						Description: mock.Anything,
					},
				},
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("Save", mock.Anything, mock.Anything).Return(nil, errors.New("error to save"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			got, err := s.Save(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Save() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_complain_DeleteById(t *testing.T) {
	type fields struct {
		dao   *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{
			name: "delete sucess",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("DeleteById", mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			if err := s.DeleteById(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

func Test_complain_FindById(t *testing.T) {
	idd := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ComplainOut
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{name: "success",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want: &model.ComplainOut{
				ID:          idd,
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
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(&model.Complain{
					ID:      idd,
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
				}, nil).Once()
			},
		},
		{
			name: "error",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				id:  mock.Anything,
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("FindById", mock.Anything, mock.Anything).Return(nil, errors.New("error ao buscar"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			got, err := s.FindById(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

//
func Test_complain_UpdateById(t *testing.T) {
	id := primitive.NewObjectID().Hex()
	//idd, _ := primitive.ObjectIDFromHex(id)

	type fields struct {
		dao   *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx context.Context
		p   model.ComplainIn
		id  string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    error
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{name: "success",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				p: model.ComplainIn{
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
				},
				id: id,
			},
			want:    nil,
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
			},
		},
		{
			name: "error",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				p: model.ComplainIn{
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
				},
				id: id,
			},
			want:    errors.New("erro ao fazer update"),
			wantErr: true,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("erro ao fazer update"))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			if err := s.Update(tt.args.ctx, tt.args.p, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("UpdateById() error = %v, wantErr %v", err, tt.wantErr)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}

//
func Test_complain_FindByParam(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.ComplainOut
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{
			name: "findALl success",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
			},
			want: []model.ComplainOut{

				model.ComplainOut{
					ID:                          id,
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
				model.ComplainOut{
					ID:                          id,
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
			},
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return([]model.Complain{
					model.Complain{

						ID:      id,
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
					},
					model.Complain{
						ID:      id,
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
					},
				}, nil).Once()
			},
		},
		{name: "findALl error",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
			},
			want:    nil,
			wantErr: true,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return(nil, errors.New("erro ao retornar"))
			}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			got, err := s.FindByParam(tt.args.ctx, nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByParam() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}
func Test_complain_FindByParam_of_arguments(t *testing.T) {
	id := primitive.NewObjectID()
	type fields struct {
		dao   *mocks.ComplainRepository
		extapi ReclameAquiApiInterface
	}
	type args struct {
		ctx  context.Context
		plan *model.ComplainIn
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []model.ComplainOut
		wantErr bool
		mock    func(repository *mocks.ComplainRepository)
	}{
		{
			name: "findByParam success",
			fields: fields{
				dao:   new(mocks.ComplainRepository),
				extapi: NewReclameAquiExternalApi(),
			},
			args: args{
				ctx: context.Background(),
				plan: &model.ComplainIn{
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
				},
			},
			want: []model.ComplainOut{

				model.ComplainOut{
					ID:                          id,
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
				model.ComplainOut{
					ID:                          id,
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
			},
			wantErr: false,
			mock: func(repository *mocks.ComplainRepository) {
				repository.On("FindByParam", mock.Anything, mock.Anything).Return([]model.Complain{
					model.Complain{

						ID:      id,
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
					},
					model.Complain{
						ID:      id,
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
					},
				}, nil).Once()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock(tt.fields.dao)
			s := &complain{
				repository: tt.fields.dao,
				externalRA: tt.fields.extapi,
			}
			got, err := s.FindByParam(tt.args.ctx, tt.args.plan)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindByParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindByParam() got = %v, want %v", got, tt.want)
			}
			tt.fields.dao.AssertExpectations(t)
		})
	}
}
