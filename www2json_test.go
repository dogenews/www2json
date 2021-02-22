package www2json

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Feedvc(t *testing.T) {
	feed, err := FeedFromLink("http://vc.ru")
	assert.NoError(t, err)
	fmt.Println(feed)
}
