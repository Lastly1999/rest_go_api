package param

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetParamInt(c *gin.Context, paramName string, defaultValue int) (int, error) {
	paramValue := c.Param(paramName)
	if paramValue == "" {
		return defaultValue, fmt.Errorf("param '%s' is missing", paramName)
	}
	intValue, err := strconv.Atoi(paramValue)
	if err != nil {
		return defaultValue, fmt.Errorf("param '%s' is not an integer", paramName)
	}
	return intValue, nil
}

func GetParamInt64(c *gin.Context, paramName string, defaultValue int64) (int64, error) {
	paramValue := c.Param(paramName)
	if paramValue == "" {
		return defaultValue, fmt.Errorf("param '%s' is missing", paramName)
	}
	intValue, err := strconv.ParseInt(paramValue, 10, 64)
	if err != nil {
		return defaultValue, fmt.Errorf("param '%s' is not a valid int64", paramName)
	}
	return intValue, nil
}
