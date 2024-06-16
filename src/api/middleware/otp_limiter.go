package middleware

import (
	"CarBuyingAndSelling/src/api/helper"
	"CarBuyingAndSelling/src/config"
	"CarBuyingAndSelling/src/pkg/limiter"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limiter.NewIPRateLimiter(rate.Every(cfg.Otp.Limiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, errors.New("not allowed")))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
