package mongoconf

import (
	"panicmode/enum"
	"panicmode/mongo"
	Scheduler "panicmode/scheduler"
	"panicmode/tools"
	"time"
)

type MongoConf struct {
	scheduler *Scheduler.JobScheduler
	table     *PanicTable
}

func CreateMongoConf(name string, interval int64, variables []*Variable) *MongoConf {
	return &MongoConf{
		scheduler: Scheduler.CreateJobScheduler(time.Duration(interval*int64(time.Second)), time.Duration(0)),
		table:     CreatePanicTable(name, variables),
	}
}

func (mongoConf *MongoConf) Sync() {
	mongoConf.sync()
	_, _ = mongoConf.scheduler.AddFunction(mongoConf.sync)
	mongoConf.scheduler.Start()
}

func (mongoConf *MongoConf) sync() {
	updatedVariables := make(map[string]bool)
	for _, updatedVariable := range mongo.GetDAO().Load(mongoConf.table.GetName()) {
		if variable, isAvailable := mongoConf.table.variables[updatedVariable.Name]; isAvailable {
			if variable.GetVariableType() == enum.GetVariableTypeFrom(updatedVariable.Type) {
				if err := variable.SetValue(updatedVariable.Value); err != nil {
					tools.Logger.Error("MONGOCONF_BAD_UPDATED_ERROR", "table", mongoConf.table.GetName(), "reason", err.Error())
				} else {
					updatedVariables[updatedVariable.Name] = true
				}
			} else {
				tools.Logger.Error("MONGOCONF_BAD_VARIABLE_TYPE_ERROR", "table", mongoConf.table.GetName(), "variable", variable.GetName(), "original-type", variable.GetVariableType(), "updated-type", updatedVariable.Type)
			}
		} else {
			tools.Logger.Error("MONGOCONF_UNKNOWN_VARIABLE_ERROR", "table", mongoConf.table.GetName(), "variable", updatedVariable.Name)
		}
	}
	for _, variable := range mongoConf.table.variables {
		if _, isAvailable := updatedVariables[variable.GetName()]; !isAvailable {
			_ = variable.SetValue(variable.GetDefaultValue())
		}
	}
}

func (mongoConf *MongoConf) GetVariable(name string) *Variable {
	return mongoConf.table.GetVariable(name)
}
