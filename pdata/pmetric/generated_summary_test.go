// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetric

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"go.opentelemetry.io/collector/pdata/internal"
	otlpmetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/metrics/v1"
)

func TestSummary_MoveTo(t *testing.T) {
	ms := generateTestSummary()
	dest := NewSummary()
	ms.MoveTo(dest)
	assert.Equal(t, NewSummary(), ms)
	assert.Equal(t, generateTestSummary(), dest)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.MoveTo(newSummary(&otlpmetrics.Summary{}, &sharedState)) })
	assert.Panics(t, func() { newSummary(&otlpmetrics.Summary{}, &sharedState).MoveTo(dest) })
}

func TestSummary_CopyTo(t *testing.T) {
	ms := NewSummary()
	orig := NewSummary()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	orig = generateTestSummary()
	orig.CopyTo(ms)
	assert.Equal(t, orig, ms)
	sharedState := internal.StateReadOnly
	assert.Panics(t, func() { ms.CopyTo(newSummary(&otlpmetrics.Summary{}, &sharedState)) })
}

func TestSummary_DataPoints(t *testing.T) {
	ms := NewSummary()
	assert.Equal(t, NewSummaryDataPointSlice(), ms.DataPoints())
	fillTestSummaryDataPointSlice(ms.DataPoints())
	assert.Equal(t, generateTestSummaryDataPointSlice(), ms.DataPoints())
}

func generateTestSummary() Summary {
	tv := NewSummary()
	fillTestSummary(tv)
	return tv
}

func fillTestSummary(tv Summary) {
	fillTestSummaryDataPointSlice(newSummaryDataPointSlice(&tv.orig.DataPoints, tv.state))
}
