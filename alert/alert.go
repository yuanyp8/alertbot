package alert

import (
	"fmt"
	"strings"
	"time"
)

type Notification struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	Status            string            `json:"status"`
	Receiver          string            `json:receiver`
	GroupLabels       map[string]string `json:groupLabels`
	CommonLabels      map[string]string `json:commonLabels`
	CommonAnnotations map[string]string `json:commonAnnotations`
	ExternalURL       string            `json:externalURL`
	Alerts            []*Alert          `json:alerts`
}

func NewNotification() *Notification {
	return &Notification{
		GroupLabels:       make(map[string]string),
		CommonLabels:      make(map[string]string),
		CommonAnnotations: make(map[string]string),
		Alerts:            make([]*Alert, 0, 20),
	}
}

type Alert struct {
	Annotations  *Annotations      `json:"annotations"`
	EndsAt       time.Time         `json:"endsAt"`
	Fingerprint  string            `json:"fingerprint"`
	StartsAt     time.Time         `json:"startsAt"`
	Status       string            `json:"status"`
	UpdatedAt    time.Time         `json:"updatedAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
}

type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

func NewAnnotations() *Annotations {
	return &Annotations{}
}

func NewAlert() *Alert {
	return &Alert{
		Annotations: NewAnnotations(),
		Labels:      make(map[string]string),
	}
}

//func Get() ([]byte, error) {
//	client := &http.Client{}
//	req, err := http.NewRequest("GET", conf.C().AlertManager.HttpAddr(), nil)
//
//	if err != nil {
//		return nil, err
//	}
//
//	response, err := client.Do(req)
//	defer response.Body.Close()
//
//	if err != nil {
//		return nil, err
//	}
//
//	return ioutil.ReadAll(response.Body)
//}

//func Unmarshal(data []byte) (*Notification, error) {
//	ret := NewNotification()
//	fmt.Println("1111")
//	if err := json.Unmarshal(data, &ret); err != nil {
//		return nil, err
//	}
//	fmt.Println("1111222")
//	return ret, nil
//
//}

//func GetAlert() (*Notification, error) {
//	data, err := Get()
//	if err != nil {
//		return nil, err
//	}
//	return Unmarshal(data)
//}

func (a *Notification) ToString() []string {
	re := make([]string, 0, 100)

	for _, v := range a.Alerts {
		dd := fmt.Sprintf("STATE: %s\nALERT_NAME: %s\nSUMMARY: %s\nSTART_AT: %v\nLABEL: %s\n", a.Status, a.CommonLabels["alertname"], v.Annotations.Summary, v.StartsAt, v.LabelToString())

		re = append(re, dd)
	}
	return re
}

func (aa *Alert) LabelToString() string {
	build := strings.Builder{}
	for k, v := range aa.Labels {
		build.WriteString(fmt.Sprintf("\n%s=%s", k, v))
	}
	return build.String()
}
