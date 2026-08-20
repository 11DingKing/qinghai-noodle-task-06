package brand

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestQinghaiBrandTask06(t *testing.T) {
	now := time.Date(2026, 8, 20, 10, 0, 0, 0, time.UTC)
	r := NewRegistry()
	l := activeLicense(now)
	require.NoError(t, r.SaveLicense(context.Background(), l, 0))
	s := NewService(r, func() time.Time { return now })
	_, err := s.RenewLicense(context.Background(), l.ID, compliantStore(now), 24*time.Hour)
	require.NoError(t, err)
}
