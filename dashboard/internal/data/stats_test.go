package data

import (
	"testing"

	"github.com/santifer/career-ops/dashboard/internal/model"
)

func TestCanonicalizeArchetype(t *testing.T) {
	tests := []struct {
		raw      string
		expected string
	}{
		{"Technical AI PM (primary) + AI Platform / LLMOps", "Agentic & Automation"}, // if agentic/automation keyword or AI PM
		{"Technical AI PM", "Technical AI PM"},
		{"Senior AI Product Manager", "Technical AI PM"},
		{"AI Platform / LLMOps", "AI Platform & LLMOps"},
		{"Agentic Automation Engineer", "Agentic & Automation"},
		{"Solutions Architect AI", "AI Solutions & FDE"},
		{"ML Engineer / Applied AI", "AI & ML Engineering"},
		{"Digital Transformation Consultant", "AI Transformation & Governance"},
		{"Data Governance Specialist", "AI Transformation & Governance"},
		{"Senior Data Engineer", "Data & Analytics"},
		{"IT Support Specialist", "IT & Technical Operations"},
		{"Wissenschaftliche Mitarbeiterin", "Research & Academia"},
		{"None", "Unclassified"},
		{"Unknown", "Unclassified"},
		{"random other role", "Other / Cross-Functional"},
	}

	for _, tt := range tests {
		got := CanonicalizeArchetype(tt.raw)
		if tt.raw == "Technical AI PM (primary) + AI Platform / LLMOps" {
			// This has Technical AI PM and AI Platform, checks Technical AI PM
			if got != "Technical AI PM" {
				t.Errorf("CanonicalizeArchetype(%q) = %q; expected %q", tt.raw, got, "Technical AI PM")
			}
			continue
		}
		if got != tt.expected {
			t.Errorf("CanonicalizeArchetype(%q) = %q; expected %q", tt.raw, got, tt.expected)
		}
	}
}

func TestCanonicalizeLocation(t *testing.T) {
	tests := []struct {
		raw      string
		expected string
	}{
		{"berlin", "Berlin"},
		{"Berlin", "Berlin"},
		{"Munich", "Munich"},
		{"münchen", "Munich"},
		{"Req, ID", ""},
		{"Social Sciences, IN", ""},
		{"Department of CS", ""},
		{"Austin, TX", "Austin, TX"},
		{"madrid", "Madrid"},
		{"", ""},
		{"—", ""},
	}

	for _, tt := range tests {
		got := CanonicalizeLocation(tt.raw)
		if got != tt.expected {
			t.Errorf("CanonicalizeLocation(%q) = %q; expected %q", tt.raw, got, tt.expected)
		}
	}
}

func TestComputeStatsMetrics(t *testing.T) {
	apps := []model.CareerApplication{
		{
			Archetype: "Technical AI PM",
			Score:     4.5,
			WorkMode:  "Remote",
			Location:  "Berlin",
			PayMax:    180000,
			PaySource: "POSTED",
		},
		{
			Archetype: "Senior AI Product Manager",
			Score:     4.0,
			WorkMode:  "Remote",
			Location:  "Berlin",
			PayMax:    200000,
			PaySource: "POSTED",
		},
		{
			Archetype: "Solutions Architect AI",
			Score:     4.2,
			WorkMode:  "Hybrid",
			Location:  "Munich",
			PayMax:    150000,
			PaySource: "est",
		},
	}

	metrics := ComputeStatsMetrics(apps)

	// Test Archetypes (Technical AI PM and AI Solutions & FDE)
	if len(metrics.Archetypes) != 2 {
		t.Fatalf("expected 2 canonical archetypes, got %d", len(metrics.Archetypes))
	}
	if metrics.Archetypes[0].Label != "Technical AI PM" || metrics.Archetypes[0].Count != 2 {
		t.Errorf("expected Technical AI PM count 2, got %+v", metrics.Archetypes[0])
	}
	if metrics.Archetypes[0].AvgScore != 4.25 {
		t.Errorf("expected avg score 4.25, got %f", metrics.Archetypes[0].AvgScore)
	}

	// Test WorkModes
	if len(metrics.WorkModes) != 2 {
		t.Fatalf("expected 2 work modes, got %d", len(metrics.WorkModes))
	}
	if metrics.WorkModes[0].Label != "Remote" || metrics.WorkModes[0].Count != 2 {
		t.Errorf("expected Remote count 2, got %+v", metrics.WorkModes[0])
	}

	// Test Locations
	if len(metrics.Locations) != 2 {
		t.Fatalf("expected 2 locations, got %d", len(metrics.Locations))
	}

	// Test Pay Stats
	if metrics.Pay.Count != 3 {
		t.Errorf("expected pay count 3, got %d", metrics.Pay.Count)
	}
	if metrics.Pay.PostedCount != 2 {
		t.Errorf("expected posted count 2, got %d", metrics.Pay.PostedCount)
	}
	if metrics.Pay.EstCount != 1 {
		t.Errorf("expected est count 1, got %d", metrics.Pay.EstCount)
	}
	if metrics.Pay.MaxPayMax != 200000 {
		t.Errorf("expected max pay 200000, got %f", metrics.Pay.MaxPayMax)
	}
	if metrics.Pay.MedianPayMax != 180000 {
		t.Errorf("expected median pay 180000, got %f", metrics.Pay.MedianPayMax)
	}

	// Test Pay Histogram
	if len(metrics.PayHistogram) != 5 {
		t.Errorf("expected 5 salary histogram bands, got %d", len(metrics.PayHistogram))
	}

	// Test Insights
	if len(metrics.Insights) == 0 {
		t.Errorf("expected non-empty insights, got 0")
	}
}
