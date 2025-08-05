package internal

import (
	"bytes"
	"errors"

	"github.com/dchest/captcha"
	"github.com/google/uuid"
)

type CaptchaService struct {
	store captcha.Store
}

const DefaultLen = 4

func NewCaptchaService() *CaptchaService {
	return &CaptchaService{
		store: captcha.NewMemoryStore(captcha.CollectNum, captcha.Expiration),
	}
}

func (s *CaptchaService) RandomImage() (string, []byte, error) {
	randomId := uuid.New().String()
	digits := captcha.RandomDigits(DefaultLen)
	s.store.Set(randomId, digits)
	var buf bytes.Buffer
	_, err := captcha.NewImage(randomId, digits, 160, 80).WriteTo(&buf)
	if err != nil {
		return "", nil, errors.New("generate captcha error")
	}

	return randomId, buf.Bytes(), nil
}

func (s *CaptchaService) Verify(captchaID, digits string) error {
	if len(captchaID) == 0 || len(digits) == 0 {
		return errors.New("params fail")
	}
	realId := s.store.Get(captchaID, true)
	if realId == nil {
		return errors.New("captcha verify fail")
	}

	ns := make([]byte, len(digits))
	for i := range ns {
		d := digits[i]
		switch {
		case '0' <= d && d <= '9':
			ns[i] = d - '0'
		case d == ' ' || d == ',':
			// ignore
		default:
			return errors.New("非法参数")
		}
	}

	if !bytes.Equal(realId, ns) {
		return errors.New("验证码错误")
	}

	return nil
}
