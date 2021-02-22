package www2json_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Feedvc(t *testing.T) {
	feed, err := www2json.FeedFromLink("http://vc.ru")
	assert.NoError(t, err)
	fmt.Println(feed)
}
