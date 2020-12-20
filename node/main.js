'use strict'

const amqp = require('amqplib/callback_api')

const queue = 'queue-example'

amqp.connect({ hostname: 'rabbitmq' }, (err0, connection) => {
  if (err0) {
    throw err0
  }

  connection.createChannel((err1, channel) => {
    if (err1) {
      throw err1
    }

    channel.assertQueue(queue, {
      durable: true
    })
    channel.prefetch(1)

    console.log('Waiting for messages in %s', queue)

    channel.consume(queue, async (msg) => {
      console.log(new Date()+': '+msg.content.toString())
      channel.ack(msg)
    })
  })
})
