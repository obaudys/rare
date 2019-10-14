package aggregation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleNumericalAggregation(t *testing.T) {
	aggr := NewNumericalAggregator()
	aggr.Sample(5)
	aggr.Sample(10)
	aggr.Sample(15)

	assert.Equal(t, uint64(3), aggr.Count())
	assert.Equal(t, 10.0, aggr.Mean())
	assert.Equal(t, 5.0, aggr.Min())
	assert.Equal(t, 15.0, aggr.Max())

	data := aggr.Analyze()

	assert.Equal(t, 10.0, data.Median())
	assert.Equal(t, 10.0, data.Quantile(0.5))
	assert.Equal(t, 5.0, data.Mode())
}

func TestSimpleMode(t *testing.T) {
	aggr := NewNumericalAggregator()
	aggr.Sample(5)
	aggr.Sample(10)
	aggr.Sample(15)
	aggr.Sample(5)
	aggr.Sample(10)
	aggr.Sample(5)

	data := aggr.Analyze()
	assert.Equal(t, 5.0, data.Mode())
	assert.Equal(t, 15.0, data.Quantile(0.9))
}