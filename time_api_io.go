package timeapiio

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

var DefaultBaseURL *url.URL
var DefaultTimeAPIIO TimeAPIIO

func init() {
	var err error
	DefaultBaseURL, err = url.Parse("https://www.timeapi.io/api")
	if err != nil {
		panic(err)
	}

	timeAPIIO, err := New()
	if err != nil {
		panic(err)
	}
	DefaultTimeAPIIO = *timeAPIIO
}

type TimeAPIIO struct {
	Config
}

type Config struct {
	HTTPClient http.Client
	BaseURL    *url.URL
}

type Option interface {
	Apply(*Config) error
}

type OptionHTTPClient http.Client

func (opt OptionHTTPClient) Apply(cfg *Config) error {
	cfg.HTTPClient = (http.Client)(opt)
	return nil
}

type OptionBaseURL string

func (opt OptionBaseURL) Apply(cfg *Config) error {
	baseURL, err := url.Parse((string)(opt))
	if err != nil {
		return fmt.Errorf("unable to parse BaseURL '%s': %w", opt, err)
	}
	cfg.BaseURL = baseURL
	return nil
}

type Options []Option

func (opts Options) Apply(cfg *Config) error {
	var errs []error
	for _, opt := range opts {
		if err := opt.Apply(cfg); err != nil {
			errs = append(errs, fmt.Errorf("unable to apply option %#+v: %w", opt, err))
		}
	}

	return errors.Join(errs...)
}

func ptr[T any](in T) *T {
	return &in
}

func (opts Options) Config() (Config, error) {
	cfg := Config{
		BaseURL: ptr(*DefaultBaseURL),
	}
	if err := opts.Apply(&cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func New(opts ...Option) (*TimeAPIIO, error) {
	cfg, err := Options(opts).Config()
	if err != nil {
		return nil, err
	}
	return &TimeAPIIO{
		Config: cfg,
	}, nil
}
