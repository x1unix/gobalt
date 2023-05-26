package dateutil_test

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/x1unix/gobalt/dateutil"
)

func Example_DurationString_UnmarshalJSON() {
	source := `{"requestTimeout": "2h30m10s"}`
	result := struct {
		RequestTimeout dateutil.DurationString `json:"requestTimeout"`
	}{}

	_ = json.Unmarshal([]byte(source), &result)
	fmt.Println(result.RequestTimeout.Duration.String())
	// Output: 2h30m10s
}

func Example_DurationString_MarshalJSON() {
	source := struct {
		RequestTimeout dateutil.DurationString `json:"requestTimeout"`
	}{
		RequestTimeout: dateutil.NewDurationString(time.Second + 500*time.Millisecond),
	}

	result, _ := json.Marshal(source)
	fmt.Println(string(result))
	// Output: {"requestTimeout":"1.5s"}
}

func Example_DurationString_TextUnmarshaler() {
	source := `2h30m10s`
	var result dateutil.DurationString

	_ = result.UnmarshalText([]byte(source))
	fmt.Println(result.Duration.String())
	// Output: 2h30m10s
}
