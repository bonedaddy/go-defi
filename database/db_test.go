package database

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDB(t *testing.T) {
	t.Run("DBOpts", func(t *testing.T) {
		tests := []struct {
			name    string
			args    *Opts
			wantErr bool
		}{
			{
				"pg-ssl", &Opts{
					Type:           "postgres",
					Host:           "localhost",
					Port:           "5432",
					User:           "user",
					Password:       "pass",
					DBName:         "indexed",
					SSLModeDisable: true,
				}, false,
			},
			{
				"pg-nossl", &Opts{
					Type:           "postgres",
					Host:           "localhost",
					Port:           "5432",
					User:           "user",
					Password:       "pass",
					DBName:         "indexed",
					SSLModeDisable: false,
				}, false,
			},
			{
				"sqlite", &Opts{
					Type:   "sqlite",
					DBName: "indexed",
				}, false,
			},
			{
				"invalid-type", &Opts{
					Type:   "eos",
					DBName: "getrekt",
				}, true,
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := tt.args.Open()
				if (err != nil) != tt.wantErr {
					t.Fatalf("Open() err %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	})
	t.Run("AutoMigrate", func(t *testing.T) {
		// TODO: enable postgresql
		t.Cleanup(func() {
			os.Remove("indexed.db")
		})
		db := newTestDB(t)
		require.NoError(t, db.Close())
	})
}

func newTestDB(t *testing.T) *Database {
	db, err := New(&Opts{Type: "sqlite", DBName: "coomer"})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate())
	return db
}
