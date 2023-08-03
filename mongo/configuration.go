package mongo

import "os"

func MongoHost() string {
	return os.Getenv("MONGOCONF_MONGO_HOST")
}

func MongoPort() string {
	return os.Getenv("MONGOCONF_MONGO_PORT")
}

func MongoUser() string {
	return os.Getenv("MONGOCONF_MONGO_USER")
}

func MongoPassword() string {
	return os.Getenv("MONGOCONF_MONGO_PASSWORD")
}

func MongoDataBase() string {
	return os.Getenv("MONGOCONF_MONGO_DATABASE")
}
