package error_code

import "errors"

var ErrInvCommand = errors.New("InvalidCommand")
var ErrInvParams = errors.New("InvalidParams")
var ErrParsing = errors.New("ParsingError")
var ErrQueueEmpty = errors.New("EmptyQueue")
var ErrKeyNotFound = errors.New("KeyNotFound")
