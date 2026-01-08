package core

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Retry[T any](ctx context.Context, config *Config, f func() (*ApiResp, error)) (*T, error) {

	var (
		resp    Response[T]
		err     error
		apiResp *ApiResp
		maxTry  = config.Retry + 1
	)

	for range maxTry {
		if config.Limiter != nil {
			if err = config.Limiter.Wait(ctx); err != nil {
				return nil, err
			}
		}

		apiResp, err = f()
		if err != nil {
			return nil, err
		}

		if apiResp.StatusCode != http.StatusOK {
			slog.Error("http error", "status", apiResp.StatusCode)
			return nil, fmt.Errorf("http error, status=%d", apiResp.StatusCode)
		}

		if err = config.Serializer.Unmarshal(apiResp.RawBody, &resp); err != nil {
			slog.Error("unmarshal error", "err", err)
			return nil, err
		}

		if resp.Status != 0 {
			slog.Error("api error", "status", resp.Status, "message", resp.Message)
			err = fmt.Errorf("api error, status=%v, message=%v", resp.Status, resp.Message)

			// retry
			if resp.Status == 100 {
				if resp.Message == "超过每分钟最大调用频率限制，请稍后重试" {
					time.Sleep(3 * time.Second)
					continue
				}
			}
		}

		err = nil
		break
	}

	return &resp.Data, err
}
