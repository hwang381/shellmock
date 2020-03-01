package lib

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_getStateFilepath(t *testing.T) {
	assert.Equal(t, getStateFilepath("/usr/local/bin/mongod"), "/tmp/usr_local_bin_mongod.json")
}
