// Copyright 2024 xiexianbin<me@xiexianbin.cn>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

import (
	"context"
	"time"

	"gorm.io/gorm"
)

// BaseModel like gorm.Model
type BaseModel[T any] struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (model *BaseModel[T]) Create(ctx context.Context) error {
	return Pool().WithContext(ctx).Create(model).Error
}

func (model *BaseModel[T]) First(ctx context.Context, id uint) error {
	return Pool().WithContext(ctx).First(model, id).Error
}

func (model *BaseModel[T]) Update(ctx context.Context, column string, value interface{}) error {
	return Pool().WithContext(ctx).Model(model).Update(column, value).Error
}

func (model *BaseModel[T]) Updates(ctx context.Context, values interface{}) error {
	return Pool().WithContext(ctx).Model(model).Updates(values).Error
}

func (model *BaseModel[T]) Delete(ctx context.Context) error {
	return Pool().WithContext(ctx).Delete(model).Error
}

func Paginate[T any](ctx context.Context, model interface{}, page int, pageSize int, order interface{}, query interface{}, args ...interface{}) (results []T, err error) {
	offset := (page - 1) * pageSize
	err = Pool().WithContext(ctx).Offset(offset).Limit(pageSize).Where(query).Order(order).Find(&results, model).Error
	return
}
