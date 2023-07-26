package mongodb_test

import (
	"context"
	"fmt"
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/Lubwama-Emmanuel/Interfaces/config"
	"github.com/Lubwama-Emmanuel/Interfaces/models"
	"github.com/Lubwama-Emmanuel/Interfaces/storage/mongodb"
)

type args struct {
	data models.DataObject
}

func TestMongo(t *testing.T) {
	t.Parallel()

	tests := []struct {
		testName string
		config   mongodb.MongoConfig
		args     args
		wantErr  assert.ErrorAssertionFunc
	}{
		{
			testName: "success",
			args: args{
				data: models.DataObject{
					"0704660968": "Emmanuel",
				},
			},
			wantErr: assert.NoError,
		},
		// {
		// 	testName: "Error/wrong connection string",
		// 	dbURL:    "mongdb://localhost:27017",
		// 	args: args{
		// 		data: models.DataObject{
		// 			"0704660968": "Emmanuel",
		// 		},
		// 	},
		// 	wantErr: assert.Error,
		// },
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			ctx := context.Background()

			config, err := config.NewConfig()
			if err != nil {
				log.WithError(err).Fatal("failed to load config")
			}

			mongoDB, _ := mongodb.NewMongoDB(ctx, config.Mongo)
			storage := mongodb.NewPhoneNumberStorage(mongoDB)

			performMongoTest(t, tc, storage)
		})
	}
}

func performMongoTest(t *testing.T, tc struct {
	testName string
	config   mongodb.MongoConfig
	args     args
	wantErr  assert.ErrorAssertionFunc
}, mongoDB *mongodb.PhoneNumberStorage,
) {
	createErr := mongoDB.Create(tc.args.data)
	if createErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, createErr))
		return
	}

	_, readErr := mongoDB.Read("0704660968")
	if readErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readErr))
		return
	}

	updateObj := models.DataObject{
		"0704660968": "Rex",
	}

	updateErr := mongoDB.Update(updateObj)
	if updateErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, updateErr))
		return
	}

	deleteErr := mongoDB.Delete("0704660968")
	if deleteErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, deleteErr))
		return
	}

	_, readAllErr := mongoDB.ReadAll()
	if readAllErr != nil && tc.wantErr == nil {
		assert.Fail(t, fmt.Sprintf("Test %v Error not expected but got one:\n"+"error: %q", tc.testName, readAllErr))
		return
	}
}
