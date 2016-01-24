package reader

import (
    . "../../Config"
    . "../../../helper"
    "gopkg.in/yaml.v2"
)

var singleReader *configReader
type configReader struct{
    errChannel errorChannel
}

func GetReader() *configReader {
    if singleReader == nil {
        singleReader = new(configReader)
        err := new(panicChannel)
        singleReader.setErrorChannel(err)
    }
    return singleReader
}

func (this* configReader) setErrorChannel(errCh errorChannel){
    this.errChannel = errCh
}

func (this configReader) Read(pathToConfig string) *Config {
    content := ReadInFile(pathToConfig)
    config := this.unmarshalToConfig(content)
    if (this.configIsMorallyOk(config)) {
        return &config
    } else  {
        return nil
    }
}


func (this configReader) configIsMorallyOk(config Config) bool {
    if len(config.Teammembers) == 1 {
        this.errChannel.Error("Sorry, I do not allow you to track the times of single persons, use more then 1 person in the team!")
        return false
    }
    return true
}

func (cr configReader) unmarshalToConfig(content []byte) Config {
    var config Config
    err := yaml.Unmarshal(content, &config)
    PanicOnError(err)
    return config
}


/*
Panic channel
 */

type errorChannel interface{
    Error(msg string)
}
type panicChannel struct {}
func (this *panicChannel) Error(msg string) {
    panic(msg)
}