package main

import (
	"context"
	"fmt"
	"time"
)

// Span represents a distributed trace span
type Span struct {
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	Attributes map[string]string
	Parent     *Span
	Children   []*Span
}

// contextKey is a private type to avoid collisions
type contextKey string

const spanKey contextKey = "span"

// Tracer provides distributed tracing capabilities
type Tracer struct {
	activeSpan *Span
}

// NewTracer creates a new tracer
func NewTracer() *Tracer {
	return &Tracer{}
}

// StartSpan starts a new span
func (t *Tracer) StartSpan(ctx context.Context, name string) context.Context {
	// Get parent span from context if it exists
	var parent *Span
	if parentSpan := ctx.Value(spanKey); parentSpan != nil {
		parent = parentSpan.(*Span)
	}

	// Create a new span
	span := &Span{
		Name:       name,
		StartTime:  time.Now(),
		Attributes: make(map[string]string),
		Parent:     parent,
	}

	// Track the current span internally
	t.activeSpan = span

	// Add to parent's children if there's a parent
	if parent != nil {
		parent.Children = append(parent.Children, span)
	}

	// Store the span in context
	return context.WithValue(ctx, spanKey, span)
}

// EndSpan ends the current span
func (t *Tracer) EndSpan(ctx context.Context) {
	// Get the span from context
	spanVal := ctx.Value(spanKey)
	if spanVal == nil {
		return
	}

	span := spanVal.(*Span)
	span.EndTime = time.Now()
	t.activeSpan = nil
}

// ActiveSpan returns the currently active span, if any
func (t *Tracer) ActiveSpan() *Span {
	return t.activeSpan
}

// AddAttribute adds an attribute to the current span
func (t *Tracer) AddAttribute(ctx context.Context, key, value string) {
	// Get the span from context
	spanVal := ctx.Value(spanKey)
	if spanVal == nil {
		return
	}

	span := spanVal.(*Span)
	span.Attributes[key] = value
}

// InjectSpanContext injects the span context into a carrier for propagation
func (t *Tracer) InjectSpanContext(ctx context.Context, carrier map[string]string) {
	// Extract span from context
	spanVal := ctx.Value(spanKey)
	if spanVal == nil {
		return
	}

	span := spanVal.(*Span)

	// Add tracing information to carrier
	carrier["trace_id"] = span.Name
	carrier["start_time"] = span.StartTime.Format(time.RFC3339Nano)

	if span.Parent != nil {
		carrier["parent_id"] = span.Parent.Name
	}

	// Add attributes
	for k, v := range span.Attributes {
		carrier[fmt.Sprintf("attr_%s", k)] = v
	}
}

// ExtractSpanContext extracts the span context from a carrier
func (t *Tracer) ExtractSpanContext(ctx context.Context, carrier map[string]string) context.Context {
	// Extract tracing information from carrier
	traceID := carrier["trace_id"]
	if traceID == "" {
		return ctx
	}

	// Create a span from the extracted information
	parentSpan := &Span{
		Name: traceID,
	}

	// Store the parent span in context
	return context.WithValue(ctx, spanKey, parentSpan)
}

func main() {}