package com.kafka.kafka;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;

@Service

public class KafkaConsumer {


    public static final Logger LOGGER = LoggerFactory.getLogger(KafkaConsumer.class) ;
    @KafkaListener(topics = "${spring.kafka.topic.name}" , groupId = "myGroup")
    public String consume(String message){
        LOGGER.info(String.format("Message recevived -> %s", message));
   return message ;
    }

}
