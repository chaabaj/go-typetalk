package v2

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	. "github.com/nulab/go-typetalk/typetalk/internal"
)

func Test_TopicsService_GetMyTopics_should_get_some_topics(t *testing.T) {
	setup()
	defer teardown()
	b, _ := ioutil.ReadFile(fixturesPath + "get-my-topics.json")
	mux.HandleFunc("/topics",
		func(w http.ResponseWriter, r *http.Request) {
			TestMethod(t, r, "GET")
			TestQueryValues(t, r, Values{
				"spaceKey": "qwerty",
			})
			fmt.Fprint(w, string(b))
		})

	result, _, err := client.Topics.GetMyTopics(context.Background(), "qwerty")
	if err != nil {
		t.Errorf("Returned error: %v", err)
	}
	var want *struct {
		Topics []*FavoriteTopicWithUnread `json:"topics"`
	}
	json.Unmarshal(b, &want)
	if !reflect.DeepEqual(result, want.Topics) {
		t.Errorf("Returned result:\n result  %v,\n want %v", result, want.Topics)
	}
}
