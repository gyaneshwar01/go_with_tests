package main

import (
	"io"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Aman", "Wins": 10},
			{"Name": "Sujal", "Wins": 32}
		]`)

		defer cleanDatabase()
		store := FileSystemPlayerStore{database}
		got := store.GetLeague()

		want := []Player{
			{"Aman", 10},
			{"Sujal", 32},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Aman", "Wins": 10},
			{"Name": "Sujal", "Wins": 32}
		]`)

		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		got := store.GetPlayerScore("Aman")
		want := 10

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Aman", "Wins": 10},
			{"Name": "Sujal", "Wins": 32}
		]`)

		defer cleanDatabase()

		store := FileSystemPlayerStore{database}

		store.RecordWin("Sujal")

		got := store.GetPlayerScore("Sujal")
		want := 33
		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Aman", "Wins": 10},
			{"Name": "Sujal", "Wins": 32}
		]`)

		defer cleanDatabase()

		store := FileSystemPlayerStore{database}
		store.RecordWin("Nitesh")

		got := store.GetPlayerScore("Nitesh")
		want := 1

		assertScoreEquals(t, got, want)
	})
}

func createTempFile(t testing.TB, initialdata string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		t.Fatalf("could not create a temp file, %v", err)
	}

	tmpFile.Write([]byte(initialdata))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertScoreEquals(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
