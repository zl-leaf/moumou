package system

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type InitTblPipeline struct {
	db           *gorm.DB
	model        interface{}
	beforeCreate IBeforeCreate
	nextPipeline INext
	errorHandler IErrorHandler
}

type IBeforeCreate func(modelObj interface{})
type INext func(modelObj interface{}) []*InitTblPipeline
type IErrorHandler func(err error) error

func NewPipeline(db *gorm.DB) *InitTblPipeline {
	return &InitTblPipeline{db: db}
}

func (p *InitTblPipeline) SetModel(m interface{}) *InitTblPipeline {
	p.model = m
	return p
}

func (p *InitTblPipeline) Before(fn IBeforeCreate) *InitTblPipeline {
	p.beforeCreate = fn
	return p
}

func (p *InitTblPipeline) Next(fn INext) *InitTblPipeline {
	p.nextPipeline = fn
	return p
}

func (p *InitTblPipeline) Execute(ctx context.Context) error {
	var err error
	db := p.db.WithContext(ctx)

	defer func() {
		err = p.handleError(err)
	}()

	result := db.WithContext(ctx).Where(p.model).First(p.model)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		p.beforeCreate(p.model)
		if result = db.Create(p.model); result.Error != nil {
			err = result.Error
			return err
		}
	} else if result.Error != nil {
		err = result.Error
		return err
	}

	if p.nextPipeline != nil {
		for _, next := range p.nextPipeline(p.model) {
			if err = next.Execute(ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (p *InitTblPipeline) handleError(err error) error {
	if err == nil {
		return nil
	}
	if p.errorHandler == nil {
		panic(err)
	}
	return p.errorHandler(err)
}
