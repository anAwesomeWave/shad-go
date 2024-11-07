//go:build !solution

package retryupdate

import (
	"errors"
	"github.com/gofrs/uuid"
	"gitlab.com/slon/shad-go/retryupdate/kvapi"
)

func UpdateValue(c kvapi.Client, key string, updateFn func(oldValue *string) (newValue string, err error)) error {
	for {
		getResp, err := c.Get(&kvapi.GetRequest{Key: key})
		var version uuid.UUID
		var auErr *kvapi.AuthError
		var cErr *kvapi.ConflictError

		var value *string

		if err != nil {
			switch {
			case errors.As(err, &auErr):
				return err
			case errors.Is(err, kvapi.ErrKeyNotFound):
				value = nil
				version = uuid.Nil
			default:
				continue
			}
		} else {
			value = &getResp.Value
			version = getResp.Version
		}

		var newVal string
		var setReq *kvapi.SetRequest
		var newVersion uuid.UUID
		var prevNewVersion uuid.UUID

	keyVanished:
		newVal, err = updateFn(value)
		if err != nil {
			return err
		}
		prevNewVersion = newVersion
		newVersion = uuid.Must(uuid.NewV4())
		setReq = &kvapi.SetRequest{
			Key:        key,
			Value:      newVal,
			OldVersion: version,
			NewVersion: newVersion,
		}
		_, err = c.Set(setReq)
		if err != nil {
			switch {
			case errors.Is(err, kvapi.ErrKeyNotFound):
				value = nil
				version = uuid.Nil
				goto keyVanished
			case errors.As(err, &cErr):
				if cErr.ExpectedVersion != prevNewVersion {
					continue
				} else {
					break
				}
			case errors.As(err, &auErr):
				return err
			default:
				goto keyVanished
			}
		}
		return nil
	}
}

// в последнем тесте. 2 запроса к set, оба возвращают ошибку, но
// 2 запрос возвращает oldVersion == NewVersion 1 запроса. Можем сделать вывод
// что 1 запрос отработал успешно
