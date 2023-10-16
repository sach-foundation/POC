package com.kafka.controller;

import com.kafka.kafka.KafkaConsumer;
import com.kafka.kafka.KafkaProducer;
import jakarta.websocket.server.PathParam;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/v1/kafka")
public class MessageController {

    private KafkaProducer kafkaProducer ;

    @Autowired
    private KafkaConsumer kafkaConsumer;
    public MessageController(KafkaProducer kafkaProducer) {
        this.kafkaProducer = kafkaProducer;
    }

    //http://localhost:8080/api/v1/kafka/publish?message=hello
    @GetMapping("/publish")
    public ResponseEntity<String> publish(@RequestParam(    "message") String message){
        kafkaProducer.sendMessage(message);
        return ResponseEntity.ok( "Entered Message is " + kafkaConsumer.consume(message));
    }
}
