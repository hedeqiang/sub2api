package service

import (
	"strings"
	"testing"
)

func TestAlertErrorTileLabel(t *testing.T) {
	if got := alertErrorTileLabel("upstream_error_rate"); got != "上游失败尝试" {
		t.Fatalf("upstream tile label = %q", got)
	}
	if got := alertErrorTileLabel("error_rate"); got != "错误数" {
		t.Fatalf("default tile label = %q", got)
	}
}

func TestAlertInsightLine_UpstreamWording(t *testing.T) {
	bd := &OpsAlertBreakdown{MetricType: "upstream_error_rate", TotalErrors: 15, Client4xx: 2, Server5xx: 1, OtherErrors: 12}
	got := alertInsightLine(bd)
	if !strings.Contains(got, "上游") {
		t.Fatalf("upstream insight should mention 上游: %s", got)
	}
	if !strings.Contains(got, "12") {
		t.Fatalf("upstream insight should surface 其他余量 12: %s", got)
	}
}

func TestAlertInsightLine_UpstreamAllOtherStillRenders(t *testing.T) {
	// 全部是无明确上游状态码的救回行:4xx=5xx=0 但 TotalErrors>0,不能返回空串。
	bd := &OpsAlertBreakdown{MetricType: "upstream_error_rate", TotalErrors: 9, OtherErrors: 9}
	if got := alertInsightLine(bd); strings.TrimSpace(got) == "" {
		t.Fatalf("upstream insight must not be empty when TotalErrors>0")
	}
}

func TestAlertInsightLine_DefaultUnchanged(t *testing.T) {
	bd := &OpsAlertBreakdown{MetricType: "error_rate", Client4xx: 0, Server5xx: 0}
	if got := alertInsightLine(bd); got != "" {
		t.Fatalf("error_rate with no 4xx/5xx should stay empty: %s", got)
	}
}

func TestEmailBreakdown_UpstreamLabel(t *testing.T) {
	bd := &OpsAlertBreakdown{MetricType: "upstream_error_rate", WindowRequests: 92, TotalErrors: 15, Client4xx: 2, Server5xx: 1, OtherErrors: 12, WindowMinutes: 5}
	html := buildOpsAlertEmailBreakdownHTML(bd)
	if !strings.Contains(html, "上游失败尝试") || !strings.Contains(html, "含重试救回") {
		t.Fatalf("email upstream wording missing: %s", html)
	}
}
