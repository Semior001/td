package proto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gotd/neo"
	"github.com/gotd/td/internal/testutil"
)

func TestMessageID(t *testing.T) {
	now := time.Date(2018, 10, 10, 23, 42, 6, 13600, time.UTC)
	id := MessageID(newMessageID(now, 0))
	if id.Type() != MessageFromClient {
		t.Fatal("invalid type")
	}
	if id != 0x5bbe8e4e00003520 {
		t.Error("mismatch")
	}
	delta := id.Time().Sub(now)
	if delta < 0 {
		delta *= -1
	}
	if delta > time.Second {
		t.Fatal("unexpected time drift")
	}
	t.Run("NewMessageID", func(t *testing.T) {
		if NewMessageID(now, MessageFromServer).Type() != MessageFromServer {
			t.Error("Mismatch")
		}
		if NewMessageID(now, 100).Type() != MessageFromClient {
			t.Error("Mismatch")
		}
	})
	t.Run("String", func(t *testing.T) {
		require.Equal(t, "5bbe8e4e00003520 (FromClient, 2018-10-10T23:42:06Z)", id.String())
	})
}

func BenchmarkNewMessageID(b *testing.B) {
	// Note that most overhead will be from time.Now() calls.
	// Just ensuring that NewMessageID itself is reasonably fast.
	now := testutil.Date()
	for i := 0; i < b.N; i++ {
		if NewMessageID(now, MessageFromServer).Type() != MessageFromServer {
			b.Fatal("Mismatch")
		}
	}
}

func TestMessageIDGen(t *testing.T) {
	date := testutil.Date()
	clock := neo.NewTime(date)

	gen := NewMessageIDGen(clock.Now, 10)
	met := make(map[int64]bool)

	for i := 0; i < 1000; i++ {
		if i%10 == 0 {
			clock.Travel(time.Millisecond * 100)
		}

		id := gen.New(MessageFromClient)
		if met[id] {
			t.Fatal("met")
		}

		met[id] = true
	}
}

func BenchmarkMsgIDGen_New(b *testing.B) {
	b.ReportAllocs()

	date := testutil.Date()
	var dateCalls int
	now := func() time.Time {
		if dateCalls%100 == 0 {
			date = date.Add(time.Millisecond)
		}
		return date
	}

	gen := NewMessageIDGen(now, 100)

	for i := 0; i < b.N; i++ {
		_ = gen.New(MessageFromServer)
	}
}
