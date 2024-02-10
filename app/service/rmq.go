package service

import (
	"fmt"
	"hr/app/utils"
	configs "hr/configs/config"
	"hr/configs/models"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func Initrmq(c *gin.Context) *models.RabbitMQMiddleware {
	// 加载配置
	rabbitMQurl := configs.Config.GetString("rabbitMQ.url")
	rabbitMQuser := configs.Config.GetString("rabbitMQ.user")
	rabbitMQpassword := configs.Config.GetString("rabbitMQ.password")

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s", rabbitMQuser, rabbitMQpassword, rabbitMQurl))
	if err != nil {
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
	}
	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
	}
	// 声明交换机
	// 用户信息交换机
	DeclareExchange(c, utils.UserExchange, "direct")
	// 全局信息交换机
	DeclareExchange(c, utils.GlobalExchange, "fanout")

	return &models.RabbitMQMiddleware{
		Connection: conn,
		Channel:    ch,
	}

}

func Closermq(r *models.RabbitMQMiddleware) {
	if r != nil {
		r.Channel.Close()
		r.Connection.Close()
	}
}

// rmq
func DeclareQueue(c *gin.Context, queueName string) amqp.Queue {
	r := GetRabbitMQMiddle(c)
	q, err := r.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		true,      // no-wait
		nil,       // arguments
	)
	if err != nil {
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
	}
	return q
}

func DeclareExchange(c *gin.Context, exchangeName, kind string) {
	// 声明交换机
	r := GetRabbitMQMiddle(c)
	err := r.Channel.ExchangeDeclare(
		exchangeName,
		kind,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
		return
	}
}

func BindQueue(c *gin.Context, queueName, routeKey, exchangeName string) {
	r := GetRabbitMQMiddle(c)
	err := r.Channel.QueueBind(
		queueName,
		routeKey,
		exchangeName,
		false,
		nil,
	)
	if err != nil {
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
		return
	}
}

func ConsumeMessage(c *gin.Context, queueName string) <-chan amqp.Delivery {
	// 声明一个消费者，并返回一个接收（receive）操作符用于从通道中接收数据的表达式
	r := GetRabbitMQMiddle(c)
	msgs, err := r.Channel.Consume(
		queueName, // queueName
		"",        // consumer
		false,     //autoAck 自动确认已读，false为手动
		false,     // exclusive 独占
		true,      // 接收自己的信息
		true,      // 不等待服务器响应，false表示等待
		nil,       // 其他参数
	)
	if err != nil {
		c.Error(utils.GetError(utils.RMQ_INIT_ERROR, err.Error()))
		c.Abort()
		return nil
	}
	return msgs
}

// for msg := range messages {
// 	// 将消息发送到前端
// 	sendMessageToClient(msg.Body)
// }

func PublishMessage(c *gin.Context, exchangeName, queueName, message string) {
	r := GetRabbitMQMiddle(c)
	err := r.Channel.Publish(
		exchangeName, // exchange
		queueName,    // routing key (队列名即为路由键) ,如果为空就是发布到全部队列
		false,        // mandatory
		false,        //
		//  immediate 参数为 false 时（默认值），如果消息无法被立即路由到队列，消息将会被存储在队列中等待消费者接收
		amqp.Publishing{
			ContentType:  "text/plain",
			Body:         []byte(message),
			DeliveryMode: amqp.Persistent, // 持久化
		},
	)
	if err != nil {
		c.Error(utils.GetError(utils.QUEUE_OPERATION_ERROR, err.Error()))
		c.Abort()
		return
	}
}
