// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildlimits/buildlimits.go - DO NOT EDIT.

package config

import (
	"time"

	"github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"
	"github.com/elastic/go-ucfg/yaml"
	"github.com/pbnjay/memory"
	"github.com/pkg/errors"
)

const (
	defaultCacheNumCounters = 500000           // 10x times expected count
	defaultCacheMaxCost     = 50 * 1024 * 1024 // 50MiB cache size

	defaultMaxConnections = 0 // no limit
	defaultPolicyThrottle = time.Millisecond * 5

	defaultCheckinInterval = time.Millisecond
	defaultCheckinBurst    = 1000
	defaultCheckinMax      = 0
	defaultCheckinMaxBody  = 1024 * 1024

	defaultArtifactInterval = time.Millisecond * 5
	defaultArtifactBurst    = 25
	defaultArtifactMax      = 50
	defaultArtifactMaxBody  = 0

	defaultEnrollInterval = time.Millisecond * 10
	defaultEnrollBurst    = 100
	defaultEnrollMax      = 50
	defaultEnrollMaxBody  = 1024 * 512

	defaultAckInterval = time.Millisecond * 10
	defaultAckBurst    = 100
	defaultAckMax      = 50
	defaultAckMaxBody  = 1024 * 1024 * 2
)

type valueRange struct {
	Min int `config:"min"`
	Max int `config:"max"`
}

type envLimits struct {
	RAM    valueRange           `config:"ram"`
	Server *serverLimitDefaults `config:"server_limits"`
	Cache  *cacheLimits         `config:"cache_limits"`
}

func defaultEnvLimits() *envLimits {
	return &envLimits{
		RAM: valueRange{
			Min: 0,
			Max: 17179869184,
		},
		Server: defaultserverLimitDefaults(),
		Cache:  defaultCacheLimits(),
	}
}

type cacheLimits struct {
	NumCounters int64 `config:"num_counters"`
	MaxCost     int64 `config:"max_cost"`
}

func defaultCacheLimits() *cacheLimits {
	return &cacheLimits{
		NumCounters: defaultCacheNumCounters,
		MaxCost:     defaultCacheMaxCost,
	}
}

type limit struct {
	Interval time.Duration `config:"interval"`
	Burst    int           `config:"burst"`
	Max      int64         `config:"max"`
	MaxBody  int64         `config:"max_body_byte_size"`
}

type serverLimitDefaults struct {
	PolicyThrottle time.Duration `config:"policy_throttle"`
	MaxConnections int           `config:"max_connections"`

	CheckinLimit  limit `config:"checkin_limit"`
	ArtifactLimit limit `config:"artifact_limit"`
	EnrollLimit   limit `config:"enroll_limit"`
	AckLimit      limit `config:"ack_limit"`
}

func defaultserverLimitDefaults() *serverLimitDefaults {
	return &serverLimitDefaults{
		PolicyThrottle: defaultCacheNumCounters,
		MaxConnections: defaultCacheMaxCost,

		CheckinLimit: limit{
			Interval: defaultCheckinInterval,
			Burst:    defaultCheckinBurst,
			Max:      defaultCheckinMax,
			MaxBody:  defaultCheckinMaxBody,
		},
		ArtifactLimit: limit{
			Interval: defaultArtifactInterval,
			Burst:    defaultArtifactBurst,
			Max:      defaultArtifactMax,
			MaxBody:  defaultArtifactMaxBody,
		},
		EnrollLimit: limit{
			Interval: defaultEnrollInterval,
			Burst:    defaultEnrollBurst,
			Max:      defaultEnrollMax,
			MaxBody:  defaultEnrollMaxBody,
		},
		AckLimit: limit{
			Interval: defaultAckInterval,
			Burst:    defaultAckBurst,
			Max:      defaultAckMax,
			MaxBody:  defaultAckMaxBody,
		},
	}
}

var defaults []*envLimits

func init() {
	// Packed Files
	// internal/pkg/config/defaults/1024_limits.yml
	// internal/pkg/config/defaults/2048_limits.yml
	// internal/pkg/config/defaults/4096_limits.yml
	// internal/pkg/config/defaults/8192_limits.yml
	// internal/pkg/config/defaults/base_limits.yml
	// internal/pkg/config/defaults/max_limits.yml
	unpacked := packer.MustUnpack("eJzsll9vqkgYxu/3Y/R6swUUc9zkXIyOIDQzRoP8u9kwYBEO/1KxMGz2u28GwbZCFZtNerOXTvCZd97nN887fz8ESb57SZzoMfvlP7pp8hz4j97u2TlG+eGR54TxX1EQB/nhDxpHD38+uLGUr/zUV2Q+I2HqIw1QDOrf+x1MfRQCbgWy0DJwauvTg2PwmbdEk3kAfJLouSngV88QOdtUq1UAKIKgQO3/pWlIRpg7rf8aIW1LEcgq21QF21SfiRzljrmp6v2Xs1ciR+FOn3LOUn31lpuDraU+1ma5W7zTE8SjbWDOMaZHt0p9XLX74dQ28ItjiI3uuq5Rmc8ykmwiN7H3BKY+NviqWQ88U608tqaBolnLLdOfKHBbILh4UuazvZtsMjuWQk9qzg5T/+u6iwJBwHQjkqivRC7PtSrsDHPgO6deCpZRTBSoFMRET2xdkVXeTepe8qj5lhjSeBWAstHcW0J+/ezy4Bp/Pvx+nSSBG/+4QVJNzhtJGAIR+4NJ4j4hif8qSegGSUgDrTuhI+uhI0RHUygzYkRcrdtxRz3pzYFvJTp1R+t3DoO3c1fbAq3ZvhL15Ci2DMyddYP22+joyTr1YunAbkuP68zlC6eUO2la9NFUXtaLwtM+loFf/ovzMz0Eft4kasxNJ9eJOhH0RtQKLkQ0nKjxJ0SVCK7HWHMFVrkrMBd02p7crVLW4YIIZWaNomdP3lMy2nBE3k4UqF8nKryXKKmno0q3ozWp9+UT6smnkzMfsoRrKbMNkRKhPNwgtUQ9mYI6maI0fZdChz/X2avXQyjtkr8elFE/+KlwnagTQe+mnWZVK7h+sgSJOfYxV2SRJzqjRq8rPmWGVbTVsQ4y4lrXVppbYo2d/M6J1zjV6pFYjKzRJiOCyCiuzq4NnXgaKLxuV7luVxcUa6Cs776pco5h763R5kambAu01LtTBXbIGn8zWTyCM757u9ad27WCw/KKOIfdp3S195TiO99OV0i6M5uUou3gJySV92XTtm96ULzuf4sMziUN9E+lyzdO7bJ+dOPpgXxNk7985wybcr1Tk17UR4cQEzvlrcf398cPir8tfngE/4+fc/z889u/AQAA///e0qUb")

	for f, v := range unpacked {
		cfg, err := yaml.NewConfig(v, DefaultOptions...)
		if err != nil {
			panic(errors.Wrap(err, "Cannot read spec from "+f))
		}

		l := defaultEnvLimits()
		if err := cfg.Unpack(&l, DefaultOptions...); err != nil {
			panic(errors.Wrap(err, "Cannot unpack spec from "+f))
		}

		defaults = append(defaults, l)
	}
}

func loadLimits() *envLimits {
	ramSize := int(memory.TotalMemory() / 1024 / 1024)
	return loadLimitsForRam(ramSize)
}

func loadLimitsForRam(currentRAM int) *envLimits {
	for _, l := range defaults {
		// get max possible config for current env
		if l.RAM.Min < currentRAM && currentRAM <= l.RAM.Max {
			return l
		}
	}

	return defaultEnvLimits()
}
