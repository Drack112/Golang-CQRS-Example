package rabbitmqcontainer

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_RabbitMQ_Container(t *testing.T) {
	rabbitmqConn, err := Start(context.Background(), t)
	require.NoError(t, err)

	assert.NotNil(t, rabbitmqConn)
}
