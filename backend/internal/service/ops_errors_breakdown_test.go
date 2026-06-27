package service

import (
	"context"
	"testing"
	"time"
)

func TestOpsService_GetErrorBreakdown_DelegatesToRepo(t *testing.T) {
	repo := &opsRepoMock{}
	svc := NewOpsService(repo, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	filter := &OpsDashboardFilter{StartTime: time.Unix(0, 0).UTC(), EndTime: time.Unix(3600, 0).UTC()}

	got, err := svc.GetErrorBreakdown(context.Background(), filter, "user", 20)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got == nil || got.Dimension != "user" {
		t.Fatalf("expected dimension passthrough, got %#v", got)
	}
}

func TestOpsService_GetErrorBreakdown_RequiresTimeRange(t *testing.T) {
	repo := &opsRepoMock{}
	svc := NewOpsService(repo, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if _, err := svc.GetErrorBreakdown(context.Background(), &OpsDashboardFilter{}, "user", 20); err == nil {
		t.Fatalf("expected error for missing time range")
	}
}
