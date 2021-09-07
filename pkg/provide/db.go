package provide

import (
	"context"

	"github.com/hans-m-song/archive-ingest/pkg/config"
	"github.com/hans-m-song/archive-ingest/pkg/ingest"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Provider struct {
	ready      bool
	connection *pgxpool.Pool
	batch      *pgx.Batch
}

func (p *Provider) Flush() error {
	if p.batch == nil {
		p.batch = &pgx.Batch{}
		return nil
	}

	result := p.connection.SendBatch(context.Background(), p.batch)
	logrus.WithField("actions", p.batch.Len()).Debug("batch flushed")

	return result.Close()
}

func (p *Provider) ListNames(table string) (*[]*NameRow, error) {
	result, err := listNameTable(p.connection, table)
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	return result, nil
}

func (p *Provider) Disconnect() error {
	if !p.ready {
		logrus.Warn("attempting to disconnect when already disconnected")
		return nil
	}

	logrus.Debug("disconnecting provider")
	err := p.Flush()
	if err != nil {
		return err
	}

	p.connection.Close()

	p.ready = false
	return nil

}

func NewProvider() (*Provider, error) {
	params := ingest.ConnectionParams{
		User: viper.GetString(config.PostgresUser),
		Pass: viper.GetString(config.PostgresPass),
		Host: viper.GetString(config.PostgresHost),
		Port: viper.GetString(config.PostgresPort),
		Name: viper.GetString(config.PostgresDatabase),
	}

	connection, err := ingest.ConnectToPostgres(params)
	if err != nil {
		return nil, err
	}

	provider := Provider{
		ready:      true,
		connection: connection,
		batch:      &pgx.Batch{},
	}

	if err != nil {
		return nil, err
	}

	return &provider, nil
}
