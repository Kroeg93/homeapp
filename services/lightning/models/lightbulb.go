package models

type LightsDTO = map[string]LightDTO

type LightDTO struct {
	State            State        `json:"state"`
	Swupdate         Swupdate     `json:"swupdate"`
	Type             string       `json:"type"`
	Name             string       `json:"name"`
	Modelid          string       `json:"modelid"`
	Manufacturername string       `json:"manufacturername"`
	Productname      string       `json:"productname"`
	Capabilities     Capabilities `json:"capabilities"`
	Config           Config       `json:"config"`
	Uniqueid         string       `json:"uniqueid"`
	Swversion        string       `json:"swversion"`
	Swconfigid       string       `json:"swconfigid"`
	Productid        string       `json:"productid"`
}

type State struct {
	On        bool      `json:"on"`
	Bri       int       `json:"bri"`
	Hue       int       `json:"hue"`
	Sat       int       `json:"sat"`
	Effect    string    `json:"effect"`
	Xy        []float64 `json:"xy"`
	Ct        int       `json:"ct"`
	Alert     string    `json:"alert"`
	Colormode string    `json:"colormode"`
	Mode      string    `json:"mode"`
	Reachable bool      `json:"reachable"`
}

type Swupdate struct {
	State       string `json:"state"`
	Lastinstall string `json:"lastinstall"`
}

type Ct struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

type Control struct {
	Mindimlevel    int         `json:"mindimlevel"`
	Maxlumen       int         `json:"maxlumen"`
	Colorgamuttype string      `json:"colorgamuttype"`
	Colorgamut     [][]float64 `json:"colorgamut"`
	Ct             Ct          `json:"ct"`
}

type Streaming struct {
	Renderer bool `json:"renderer"`
	Proxy    bool `json:"proxy"`
}

type Capabilities struct {
	Certified bool      `json:"certified"`
	Control   Control   `json:"control"`
	Streaming Streaming `json:"streaming"`
}

type Startup struct {
	Mode       string `json:"mode"`
	Configured bool   `json:"configured"`
}

type Config struct {
	Archetype string  `json:"archetype"`
	Function  string  `json:"function"`
	Direction string  `json:"direction"`
	Startup   Startup `json:"startup"`
}
