package app

//
//type Application struct {
//	Repository   *repository.Repository
//	KafkaHandler *handlers.KafkaHandler
//}
//
//func InitializeApp() *Application {
//	database := di.NewORM()
//	kafkaProducer := kafka.InitKafkaProducer()
//	kafkaConsumer := kafka.InitKafkaConsumer()
//	redisClient := redis.InitRedis()
//
//	return &Application{
//		Repository:   handlers.NewRepository(database),
//		KafkaHandler: handlers.NewKafkaHandler(kafkaConsumer),
//	}
//}
