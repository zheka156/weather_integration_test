package weather

type ValueItem struct {
	Value string `json:"value"`
}

type Response struct {
	CurrentCondition []CurrentCondition `json:"current_condition"`
	NearestArea      []NearestArea      `json:"nearest_area"`
	Request          []RequestItem      `json:"request"`
	Weather          []WeatherDay       `json:"weather"`
}

type CurrentCondition struct {
	FeelsLikeC       string      `json:"FeelsLikeC"`
	FeelsLikeF       string      `json:"FeelsLikeF"`
	Cloudcover       string      `json:"cloudcover"`
	Humidity         string      `json:"humidity"`
	LangRU           []ValueItem `json:"lang_ru"`
	LocalObsDateTime string      `json:"localObsDateTime"`
	ObservationTime  string      `json:"observation_time"`
	PrecipInches     string      `json:"precipInches"`
	PrecipMM         string      `json:"precipMM"`
	Pressure         string      `json:"pressure"`
	PressureInches   string      `json:"pressureInches"`
	TempC            string      `json:"temp_C"`
	TempF            string      `json:"temp_F"`
	UVIndex          string      `json:"uvIndex"`
	Visibility       string      `json:"visibility"`
	VisibilityMiles  string      `json:"visibilityMiles"`
	WeatherCode      string      `json:"weatherCode"`
	WeatherDesc      []ValueItem `json:"weatherDesc"`
	WeatherIconURL   []ValueItem `json:"weatherIconUrl"`
	Winddir16Point   string      `json:"winddir16Point"`
	WinddirDegree    string      `json:"winddirDegree"`
	WindspeedKmph    string      `json:"windspeedKmph"`
	WindspeedMiles   string      `json:"windspeedMiles"`
}

type NearestArea struct {
	AreaName   []ValueItem `json:"areaName"`
	Country    []ValueItem `json:"country"`
	Latitude   string      `json:"latitude"`
	Longitude  string      `json:"longitude"`
	Population string      `json:"population"`
	Region     []ValueItem `json:"region"`
	WeatherURL []ValueItem `json:"weatherUrl"`
}

type RequestItem struct {
	Query string `json:"query"`
	Type  string `json:"type"`
}

type WeatherDay struct {
	Astronomy   []Astronomy  `json:"astronomy"`
	AvgtempC    string       `json:"avgtempC"`
	AvgtempF    string       `json:"avgtempF"`
	Date        string       `json:"date"`
	Hourly      []HourlyItem `json:"hourly"`
	MaxtempC    string       `json:"maxtempC"`
	MaxtempF    string       `json:"maxtempF"`
	MintempC    string       `json:"mintempC"`
	MintempF    string       `json:"mintempF"`
	SunHour     string       `json:"sunHour"`
	TotalSnowCM string       `json:"totalSnow_cm"`
	UVIndex     string       `json:"uvIndex"`
}

type Astronomy struct {
	MoonIllumination string `json:"moon_illumination"`
	MoonPhase        string `json:"moon_phase"`
	Moonrise         string `json:"moonrise"`
	Moonset          string `json:"moonset"`
	Sunrise          string `json:"sunrise"`
	Sunset           string `json:"sunset"`
}

type HourlyItem struct {
	DewPointC        string      `json:"DewPointC"`
	DewPointF        string      `json:"DewPointF"`
	FeelsLikeC       string      `json:"FeelsLikeC"`
	FeelsLikeF       string      `json:"FeelsLikeF"`
	HeatIndexC       string      `json:"HeatIndexC"`
	HeatIndexF       string      `json:"HeatIndexF"`
	WindChillC       string      `json:"WindChillC"`
	WindChillF       string      `json:"WindChillF"`
	WindGustKmph     string      `json:"WindGustKmph"`
	WindGustMiles    string      `json:"WindGustMiles"`
	ChanceOfFog      string      `json:"chanceoffog"`
	ChanceOfFrost    string      `json:"chanceoffrost"`
	ChanceOfHighTemp string      `json:"chanceofhightemp"`
	ChanceOfOvercast string      `json:"chanceofovercast"`
	ChanceOfRain     string      `json:"chanceofrain"`
	ChanceOfRemDry   string      `json:"chanceofremdry"`
	ChanceOfSnow     string      `json:"chanceofsnow"`
	ChanceOfSunshine string      `json:"chanceofsunshine"`
	ChanceOfThunder  string      `json:"chanceofthunder"`
	ChanceOfWindy    string      `json:"chanceofwindy"`
	Cloudcover       string      `json:"cloudcover"`
	DiffRad          string      `json:"diffRad"`
	Humidity         string      `json:"humidity"`
	LangRU           []ValueItem `json:"lang_ru"`
	PrecipInches     string      `json:"precipInches"`
	PrecipMM         string      `json:"precipMM"`
	Pressure         string      `json:"pressure"`
	PressureInches   string      `json:"pressureInches"`
	ShortRad         string      `json:"shortRad"`
	TempC            string      `json:"tempC"`
	TempF            string      `json:"tempF"`
	Time             string      `json:"time"` // "0", "300", ...
	UVIndex          string      `json:"uvIndex"`
	Visibility       string      `json:"visibility"`
	VisibilityMiles  string      `json:"visibilityMiles"`
	WeatherCode      string      `json:"weatherCode"`
	WeatherDesc      []ValueItem `json:"weatherDesc"`
	WeatherIconURL   []ValueItem `json:"weatherIconUrl"`
	Winddir16Point   string      `json:"winddir16Point"`
	WinddirDegree    string      `json:"winddirDegree"`
	WindspeedKmph    string      `json:"windspeedKmph"`
	WindspeedMiles   string      `json:"windspeedMiles"`
}
