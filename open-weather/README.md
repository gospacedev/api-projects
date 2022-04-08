# Open Weather
A Weather API implementation

## Usage

Get an API Weather Key from:

https://openweathermap.org/

and paste it in .apiConfig:

    {
    "OpenWeatherMapApiKey":"<paste-API-key-here>"
    }

You can test it by running main.go and pasting this to your browser:

    http://localhost:8000/hello

Then enter this and a name of a city in your browser:

    http://localhost:8000/weather/<name-of-a-city>
    
This format should show in your browser:

    http://localhost:8000/weather/manila
    {"Name":"Manila","main":{"temp":27.09}}
    
You can add more stats from their documentation:

https://openweathermap.org/current
