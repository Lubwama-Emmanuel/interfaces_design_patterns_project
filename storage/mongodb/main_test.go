package mongodb_test

import (
	"testing"

	log "github.com/sirupsen/logrus"

	"github.com/Lubwama-Emmanuel/Interfaces/config"
)

var cfig config.Config

func TestMain(m *testing.M) {
	testConfig, err := config.NewConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to load configs")
	}

	cfig = testConfig
}
