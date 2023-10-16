package com.kafka.kafka;

import com.kafka.payload.User;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service
public class JsonKafkaConsumer {

    private static final Logger LOGGER  = LoggerFactory.getLogger(JsonKafkaConsumer.class) ;

//    @KafkaListener(topics = "topic_json",groupId = "myGroup")
    @KafkaListener(topics = "${spring.kafka.topic-json.name}",groupId = "myGroup")
    public  void consume(User user){
        LOGGER.info(String.format("JSON Message Received --> %s" , user));
    }

}
