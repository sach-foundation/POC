package com.kafka.controller;

import com.kafka.kafka.JsonKafkaProducer;
import com.kafka.payload.User;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/vi/kafka")
public class JsonMessageController {

    private JsonKafkaProducer kafkaProducer ;

    public JsonMessageController(JsonKafkaProducer kafkaProducer) {
        this.kafkaProducer = kafkaProducer;
    }

    @PostMapping("/publish")
    public ResponseEntity<String> publish(@RequestBody User user) {

        kafkaProducer.sendMessage(user);

        return new ResponseEntity<>("Message Created", HttpStatus.CREATED) ;
    }
}
