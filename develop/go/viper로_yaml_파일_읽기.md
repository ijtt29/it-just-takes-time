# Go viper libray로 yaml을 읽어 Config로 사용하는 방법

API서버 구동시 혹은, AWS와 연관된 기능을 수행하기 위해 필요한 정보들을 Config파일로 관리하려고한다.

GoLang은 어떤식으로 Config파일을 읽는지 찾아보다가 viper라는걸 찾게 되었다.

[링크](https://github.com/spf13/viper)

간단한 예시코드

```
go get github.com/spf13/viper
```

예시 config.yaml파일
```
# configs/configs.yaml
database:
  databaseType: "mysql"
  host: "localhost:3306"
application:
  port: 9000
aws:
  provider:
    region: XXXXXXXXXXX
    accessKey: XXXXXXXXXXXXXXX
    secretKey: XXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
  s3:
    bucket: sample
```

## config.yaml 읽어오기
```
# configs/configs.go
package configs

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init(env string) {
	var err error

	viper.AddConfigPath("./configs")
	viper.SetConfigName("config.yaml")
	viper.SetConfigType("yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file")
	}

    # 읽어온 yaml파일의 k,v값을 struct로 매핑한다.
	err = viper.Unmarshal(&Conf)
	if err != nil {
		log.Fatal(err)
	}
}

```

## 읽어온 config에 매핑할 struct
```
# configs/configure.go
package configs

var Conf = Configure{}

type Configure struct {
	Application Application `yaml: "server"`
	Database    Database    `yaml: "database"`
	AWS         AWS         `yaml: "aws"`
}

type Application struct {
	Port string `yaml: "port"`
}

type Database struct {
	DatabaseType string `yaml: "databaseType"`
	Host         string `yaml: "host"`
	Port         string `yaml: "port"`
	Database     string `yaml: "database"`
	Username     string `yaml: "username"`
	Password     string `yaml: "password"`
}

type AWS struct {
	Provider Provider `yaml: "provider"`
	S3       S3       `yaml: "s3"`
}

type Provider struct {
	AccessKey string `yaml: "accessKey"`
	SecretKey string `yaml: "secretKey"`
	Region    string `yaml: "region"`
}

type S3 struct {
	Bucket string `yaml: "bucket"`
}

```

## 사용하기
어플리케이션 실행후 confing파일을 읽어와 Configure struct로 매핑이 되었다면 다음과 같이 사용할 수 있다.
```
v = configs.Conf.AWS.Provider.Region
```