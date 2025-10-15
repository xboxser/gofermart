package model

import "errors"

var (
	Error202 = errors.New("новый номер заказа принят в обработку")
	Error409 = errors.New("номер заказа уже был загружен другим пользователем")
)
