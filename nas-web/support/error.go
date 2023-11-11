package support

import "errors"

type NasError error

var InitRedisError NasError = errors.New("Init Redis Failed")
var InitServiceError NasError = errors.New("Init Service Failed")
var InitAppError NasError = errors.New("Init App Failed")
