package alert

import (
	"encoding/json"
	"fmt"
	"github.com/yuanyp8/alertbot/conf"
	"io/ioutil"
	"net/http"
	"time"
)

type Alert struct {
	Annotations  *Annotations `json:"annotations"`
	EndsAt       time.Time    `json:"endsAt"`
	Fingerprint  string       `json:"fingerprint"`
	Receivers    []*Receiver  `json:"receivers"`
	StartsAt     time.Time    `json:"startsAt"`
	Status       *Status      `json:"status"`
	UpdatedAt    time.Time    `json:"updatedAt"`
	GeneratorURL string       `json:"generatorURL"`
	Labels       *Labels      `json:"labels"`
}
type Status struct {
	InhibitedBy []interface{} `json:"inhibitedBy"`
	SilencedBy  []interface{} `json:"silencedBy"`
	State       string        `json:"state"`
}
type Receiver struct {
	Name string `json:"name"`
}
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
}

type Labels struct {
	Agent        string `json:"agent,omitempty"`
	Alertname    string `json:"alertname"`
	Dc           string `json:"dc,omitempty"`
	Device       string `json:"device,omitempty"`
	Env          string `json:"env,omitempty"`
	Fstype       string `json:"fstype,omitempty"`
	Hostname     string `json:"hostname,omitempty"`
	Instance     string `json:"instance,omitempty"`
	Job          string `json:"job,omitempty"`
	Mountpoint   string `json:"mountpoint,omitempty"`
	Severity     string `json:"severity"`
	Type         string `json:"type,omitempty"`
	ExportedType string `json:"exported_type,omitempty"`
	Name         string `json:"name,omitempty"`
	State        string `json:"state,omitempty"`
}

func NewAnnotations() *Annotations {
	return &Annotations{}
}
func NewLabels() *Labels {
	return &Labels{}
}

func NewReceiver() *Receiver {
	return &Receiver{}
}

func NewStatus() *Status {
	return &Status{}
}

func NewAlert() *Alert {
	return &Alert{
		Annotations: NewAnnotations(),
		Labels:      NewLabels(),
		Receivers:   make([]*Receiver, 0, 100),
		Status:      NewStatus(),
	}
}

func Get() ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", conf.C().AlertManager.HttpAddr(), nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	defer response.Body.Close()

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}

func Unmarshal(data []byte) ([]*Alert, error) {
	ret := make([]*Alert, 0, 300)
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return ret, nil

}

func GetAlert() ([]*Alert, error) {
	data, err := Get()
	if err != nil {
		return nil, err
	}
	return Unmarshal(data)
}

func (a *Alert) ToString() string {
	return fmt.Sprintf("STATE: %s\nSUMMARY: %s\nSTART_AT: %v\nGENERATOR_URL: %s\n\nLABEL: %s", a.Status.State, a.Annotations.Summary, a.StartsAt, a.GeneratorURL, a.Labels.ToString())
}

func (l *Labels) ToString() string {
	data, err := json.Marshal(l)
	if err != nil {
		return "Label Parse Failed"
	}

	return string(data)
}
