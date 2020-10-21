# Config
_A simple means of updating configuration provided_

## Example
This package leverages previous work of utilising struct tags in order to caputre desired inputs.

Environment parsing is handled by [envconfig](https://github.com/kelseyhightower/envconfig).  
Flag parsing is handled by [go-flags](https://github.com/jessevdk/go-flags).  

In order to utilise it, you can do something like the following:
```
type Values struct {
  LogLevel  string `long:"log-level" description:"Set the logging level for the application"`
  EnvSecret string `envconfig:"env_secret" split_words:"true"`
}

func example() {
  v := &Values{LogLevel: "info"}
  conf := config.Default(v)
  if err := conf.ParseFlags().ParseEnv().Err(); err != nil {
    panic(err)
  }
}
```