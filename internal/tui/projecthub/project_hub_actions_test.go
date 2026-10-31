package projecthub

import (
	"slices"
	"testing"

	"github.com/sglavoie/docket/internal/tui/operation"
)

// Direct operations stay out of Space; a single active row offers archiving and
// deletion. Bulk tag actions need a selection and drop out of a one-row sheet.
func TestProjectHubActiveProjectOffersArchiveAndDelete(t *testing.T) {
	hub := newProjectHubFixture(t,
		scopeTestProject("alpha", "Alpha", "1", "2"),
		scopeTestProject("beta", "Beta", "1"),
	).open()
	focusProjectHubItem(t, hub, "project-hub:project:beta")
	hub.Update(syntheticKey(' '))

	got := projectHubActionIDs(hub)
	want := []operation.OperationID{operation.OperationScopeArchiveProject, operation.OperationProjectDelete}
	if !slices.Equal(got, want) {
		t.Fatalf("active project actions = %v, want %v", got, want)
	}
}

// Archived projects remain deletable without restoring them first. Active-only
// operations, including notes, stay absent from their sheet.
func TestProjectHubArchivedProjectOffersUnarchiveAndDelete(t *testing.T) {
	fixture := newProjectHubFixture(t, scopeTestProject("alpha", "Alpha", "1", "2"))
	archived := scopeTestProject("archived", "Archived", "1")
	archived.Archived = true
	fixture.add(archived)

	hub := fixture.open()
	focusProjectHubItem(t, hub, "project-hub:archived:archived")
	hub.Update(syntheticKey(' '))
	got := projectHubActionIDs(hub)
	want := []operation.OperationID{operation.OperationScopeRestoreProject, operation.OperationProjectDelete}
	if !slices.Equal(got, want) {
		t.Fatalf("archived project actions = %v, want %v", got, want)
	}
}
