package service

import (
	"context"
	"log"
	"projeto-ra-api-go/pkg/api/model"
	"sync"

	"golang.org/x/sync/errgroup"
)

type Complain interface {
	Save(parentContext context.Context, complain *model.ComplainIn) (*model.ComplainOut, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, p model.ComplainIn, id string) error
	FindById(ctx context.Context, id string) (*model.ComplainOut, error)
	FindByIdWithExternal(ctx context.Context, id string) (*model.ComplainOut, error)
	FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error)
	FindByParamWithExternal(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error)
}

type ComplainRepository interface {
	Save(parentContext context.Context, complain *model.Complain) (string, error)
	DeleteById(ctx context.Context, id string) error
	Update(ctx context.Context, p *model.Complain, id string) error
	FindById(ctx context.Context, id string) (*model.Complain, error)
	FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.Complain, error)
}

type complain struct {
	repository ComplainRepository
	externalRA ReclameAquiApiInterface
}

func NewComplain(complains ComplainRepository, extapi ReclameAquiApiInterface) Complain {
	return &complain{repository: complains, externalRA: extapi}
}

func (s *complain) Save(ctx context.Context, in *model.ComplainIn) (*model.ComplainOut, error) {
	complainModel := in.ToComplain()
	HexID, err := s.repository.Save(ctx, complainModel)

	if err != nil {
		return nil, err
	}
	complainOut, err := s.FindById(ctx, HexID)

	return complainOut, nil
}

func (s *complain) DeleteById(ctx context.Context, id string) error {

	err := s.repository.DeleteById(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *complain) Update(ctx context.Context, p model.ComplainIn, id string) error {
	complain := p.ToComplain()

	err := s.repository.Update(ctx, complain, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *complain) FindById(ctx context.Context, id string) (*model.ComplainOut, error) {
	complain, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	complainOut := complain.ToComplainOut()

	complainOut.CountPageViews = 0

	return complainOut, nil
}

func (s *complain) FindByIdWithExternal(ctx context.Context, id string) (*model.ComplainOut, error) {
	complain, err := s.repository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	complainOut := complain.ToComplainOut()

	appearances, err := s.externalRA.CountPageViewsRaExternalApi(ctx, complainOut.Company.Title)
	if err != nil {
		return nil, err
	}

	complainOut.CountPageViews = appearances

	isOnTop10BadRA, err := s.externalRA.IsInTopBad10RaExternalApi(ctx, complainOut.Company.Title)
	if err != nil {
		return nil, err
	}

	complainOut.IsOnTop10BadRA = isOnTop10BadRA

	return complainOut, nil
}

func (s *complain) FindByParam_exemplo1(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {

	complains, err := s.repository.FindByParam(ctx, param)

	if err != nil {
		return nil, err
	}
	var complainOuts []model.ComplainOut
	ch := make(chan model.ComplainOut, len(complains))
	//O waitgroup serve para fazer a thread principal esperar todas as goroutines terminarem
	var w sync.WaitGroup
	for _, complain := range complains {
		// significa que a concorrencia sera feia de um em um ele diz quantos devem esperar
		w.Add(1)
		complainOut := complain.ToComplainOut()
		// uma goroutines
		go s.getCountComplain(ctx, *complainOut, ch, &w)
	}
	// fechando a goroutines
	go waitTOClose(&w, ch)

	for complainOut := range ch {
		complainOuts = append(complainOuts, complainOut)
	}

	return complainOuts, nil
}
func waitTOClose(w *sync.WaitGroup, ch chan model.ComplainOut) {
	w.Wait()
	close(ch)
}
func (s *complain) getCountComplain(ctx context.Context, p model.ComplainOut, ch chan model.ComplainOut, w *sync.WaitGroup) {
	// e executado no retorno do metodo
	defer w.Done()
	appearances, err := s.externalRA.CountPageViewsRaExternalApi(ctx, p.Company.Title)

	if err != nil {
		return
	}

	p.CountPageViews = appearances
	ch <- p
}

func (s *complain) FindByParam(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {
	complains, err := s.repository.FindByParam(ctx, param)
	if err != nil {
		return nil, err
	}
	var complainOuts []model.ComplainOut
	for _, complain := range complains {
		complainOut := complain.ToComplainOut()
		complainOuts = append(complainOuts, *complainOut)
	}
	return complainOuts, nil
}

func (s *complain) FindByParam_exemplo2(ctx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {
	complains, err := s.repository.FindByParam(ctx, param)
	if err != nil {
		return nil, err
	}
	var complainOuts []model.ComplainOut
	//ch := make(chan model.ComplainOut, len(complains))
	var wg sync.WaitGroup
	wg.Add(len(complains))
	for _, complain := range complains {
		complainOut := complain.ToComplainOut()
		go func(complain model.ComplainOut) error {
			defer wg.Done()
			appearances, err := s.externalRA.CountPageViewsRaExternalApi(ctx, complain.Company.Title)
			if err != nil {
				log.Fatalf("Error on calling external api: ", err)
				return err
			}
			complain.CountPageViews = appearances

			isOnTop10BadRA, err := s.externalRA.IsInTopBad10RaExternalApi(ctx, complain.Company.Title)
			if err != nil {
				return err
			}

			complain.IsOnTop10BadRA = isOnTop10BadRA

			return nil
		}(*complainOut)
		complainOuts = append(complainOuts, *complainOut)
	}
	wg.Wait()
	// n usei channel
	return complainOuts, nil
}

func (s *complain) FindByParamWithExternal(parentCtx context.Context, param *model.ComplainIn) ([]model.ComplainOut, error) {

	complains, err := s.repository.FindByParam(parentCtx, param)
	if err != nil {
		return nil, err
	}
	// group
	g, ctx := errgroup.WithContext(parentCtx)
	defer ctx.Done()
	// channel
	ch := make(chan *model.ComplainOut, len(complains))

	g.Go(func() error {

		defer close(ch)
		childG, _ := errgroup.WithContext(ctx)

		for _, complain := range complains {
			complain := complain

			childG.Go(func() error {
				p := complain.ToComplainOut()

				appearances, err := s.externalRA.CountPageViewsRaExternalApi(ctx, p.Company.Title)
				if err != nil {
					return err
				}

				p.CountPageViews = appearances

				isOnTop10BadRA, err := s.externalRA.IsInTopBad10RaExternalApi(ctx, p.Company.Title)
				if err != nil {
					return err
				}

				p.IsOnTop10BadRA = isOnTop10BadRA

				ch <- p

				return nil
			})
		}

		return childG.Wait()
	})

	var complainOuts []model.ComplainOut
	g.Go(func() error {
		for p := range ch {
			complainOuts = append(complainOuts, *p)
		}

		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return complainOuts, nil
}
