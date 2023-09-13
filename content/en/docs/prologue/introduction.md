---
title: "API"


date: 2020-10-06T08:48:57+00:00
lastmod: 2020-10-06T08:48:57+00:00
draft: false
images: []
menu:
  docs:
    parent: "prologue"
weight: 100
toc: true
---


<hr/>

### POST : Create Challenge

 <!-- {{< alert icon="👉" text="{{base_url}}/v1/challenges" />}} -->
 {{<alert icon="URL" text="{{base_url}}/v1/challenges" />}}


{{<alert text="Body"  />}}
<p >
{<br/>
  "created_at": "2017-09-11T15:12:59.152Z", <br/>
  "created_by": 1, <br/>
  "description": "History",<br/>
  "duration_minutes": 0,<br/>
  "end_date": "2017-09-11T15:12:59.152Z",<br/>
  "is_active": true,<br/>
  "is_public": true,<br/>
  "max_participants": 0,<br/>
  "start_date": "2017-09-11T15:12:59.152Z",<br/>
  "title": "History",<br/>
  "updated_at": "2017-09-11T15:12:59.152Z",<br/>
  "updated_by": 1<br/>
}

</p>


{{<alert  text="Expected Response"  />}}

{<br/>
    "challenges_id": 1,<br/>
    "player_userid": 123,<br/>
    "player_useremail": "abc.com",<br/>
    "player_gamestartime": "2023-09-11T14:58:10.9614009+05:30",<br/>
    "player_endgametime": "2023-09-11T14:58:10.9614009+05:30",<br/>
    "total_questions": 0,<br/>
    "challenge_duration_seconds": 0,<br/>
    "player_time_elapsed": 0,<br/>
    "player_right_answers": 0,<br/>
    "player_correct_ans_percentage": 0<br/>
}
<hr/>

### GET :  Challenge

{{<alert icon="URL" text="{{base_url}}/v1/challenges?limit=10&offset=0" />}}

{{<alert icon="Response" text="Body"  />}}
[ <br/>
  { <br/>
    "id": 1,<br/>
    "title": "History",<br/>
    "description": "History",<br/>
    "game_room_conn_strg": "not generated",<br/>
    "start_date": "2017-09-11T15:12:59.152Z",<br/>
    "end_date": "2017-09-12T15:12:59.152Z",<br/>
    "no_of_questions": 0,<br/>
    "duration_minutes": 0,<br/>
    "max_participants": 0,<br/>
    "is_public": true,<br/>
    "is_active": true,<br/>
    "created_at": "2023-09-11T09:02:13.754526Z",<br/>
    "updated_at": "2023-09-11T09:02:13.754526Z",<br/>
    "created_by": 1,<br/>
    "updated_by": 1<br/>
  }<br/>
]<br/>
<hr/>

### PUT : Create Challenge

 <!-- {{< alert icon="👉" text="{{base_url}}/v1/challenges" />}} -->
 {{<alert icon="URL" text="{{base_url}}/v1/challenges" />}}

{{<alert icon="Request" text=" Body" />}}

{ <br/>
  "description": "History Upd",<br/>
  "duration_minutes": 0,<br/>
  "end_date": "2017-05-15T15:12:59.152Z",<br/>
  "id": 1,<br/>
  "is_active": true,<br/>
  "is_public": true,<br/>
  "max_participants": 0,<br/>
  "start_date": "2017-05-15T15:12:59.152Z",<br/>
  "title": "History Upd",<br/>
  "updated_at": "2017-05-15T15:12:59.152Z",<br/>
  "updated_by": 2<br/>
} 


{{<alert icon="Response" text=" Body" />}}

{ <br/>
  "id": 1, <br/>
  "title": "History Upd", <br/>
  "description": "History Upd", <br/>
  "game_room_conn_strg": "not generated", <br/>
  "start_date": "2017-09-11T15:12:59.152Z",<br/>
  "end_date": "2017-09-12T15:12:59.152Z",<br/>
  "no_of_questions": 0, <br/>
  "duration_minutes": 0, <br/>
  "max_participants": 0, <br/>
  "is_public": true, <br/>
  "is_active": true, <br/>
  "created_at": "2023-09-11T09:02:13.754526Z", <br/>
  "updated_at": "2023-09-11T09:02:13.754526Z", <br/>
  "created_by": 1, <br/>
  "updated_by": 1 <br/>
}
<hr/>

### POST :  Questions

{{<alert icon="URL" text="{{base_url}}/v1/questions" />}}

{{<alert icon="Request" text="Body"  />}}

{ <br/>
  "challenges_id": 1,<br/>
  "created_at": "2023-09-11T09:02:13.754526Z",<br/>
  "created_by": 1,<br/>
  "is_active": true,<br/>
  "question": "When did the First War Of Independence happen ?",<br/>
   "answer" : "1857"<br/>
  "sequence_no": 0,<br/>
  "updated_at": "2023-09-11T09:02:13.754526Z",<br/>
  "updated_by": 1<br/>
} <br/>

{{<alert icon="Expected" text="Response Body" />}}

{ <br/>
  "question_id": 1,<br/>
  "challenges_id": 1,<br/>
  "question": "When did the First War Of Independence happen ?",<br/>
  "answer": "1857",<br/>
  "sequence_no": 0,<br/>
  "is_active": true,<br/>
  "created_at": "2023-09-11T09:02:13.754526Z",<br/>
  "updated_at": "2023-09-11T09:02:13.754526Z",<br/>
  "created_by": 1,<br/>
  "updated_by": 1<br/>
}<br/>
<hr/>

### GET :  Questions

{{<alert icon="URL" text="{{base_url}}/v1/questions/1" />}}

{{<alert icon="Response" text="Body"  />}}

{ <br/>
  "question_id": 1, <br/>
  "challenges_id": 1, <br/>
  "question": "When did the First War Of Independence happen ?", <br/>
  "answer": "1857", <br/>
  "sequence_no": 0, <br/>
  "is_active": true, <br/>
  "created_at": "2023-09-11T09:02:13.754526Z", <br/>
  "updated_at": "2023-09-11T09:02:13.754526Z", <br/>
  "created_by": 1, <br/>
  "updated_by": 1 <br/>
} 
<hr/>

### POST :  Questions

{{<alert icon="URL" text="{{base_url}}/v1/choices" />}}

{{<alert icon="Request" text="Body"  />}}

<p>

{ <br/>
  "choice_text": "1857", <br/>
  "created_at": "2023-09-11T09:02:13.754526Z", <br/>
  "created_by": 1, <br/>
  "is_active": true, <br/>
  "is_correct": true, <br/>
  "question_id": 1, <br>
  "updated_at": "2023-09-11T09:02:13.754526Z", <br/>
  "updated_by": 1 <br/>
}



</p>



{{<alert icon="Expected" text="Response Body" />}}

<p>
{
  "choice_id": 1, <br/>
  "question_id": 1, <br/>
  "choice_text": "1857", <br/>
  "sequence_no": 0, <br/>
  "is_correct": true, <br/>
  "is_active": true, <br/>
  "created_at": "2023-09-11T09:02:13.754526Z", <br/>
  "updated_at": "2023-09-11T09:02:13.754526Z", <br/>
  "created_by": 1, <br/>
  "updated_by": 1 <br/>
}



</p>

### PUT :  Generate Connection String

{{<alert icon="URL" text="{{base_url}}/v1/challenges/generateConnString" />}}

{{<alert icon="Request" text="Body"  />}}

{ <br/>
    
     "challenges_id" : 1 <br/>
} <br/>


{{<alert icon="Response" text="Body"  />}}

{ <br/>
    "id": 1, <br/>
    "title": "Current Affairs", <br/>
    "description": "Current Affairs", <br/>
    "game_room_conn_strg": "RXWULU",<br/>
    "start_date": "2023-09-07T23:12:59.152Z",<br/>
    "end_date": "2023-09-08T23:12:59.152Z", <br/>
    "no_of_questions": 2,<br/>
    "duration_minutes": 0,<br/>
    "max_participants": 0,<br/>
    "is_public": true,<br/>
    "is_active": true,<br/>
    "created_at": "2023-09-12T12:31:03.876676Z",<br/>
    "updated_at": "2023-09-12T12:31:03.876676Z",<br/>
    "created_by": 1,<br/>
    "updated_by": 1<br/>
}

<hr/>



### POST :  Player Create Session

{{<alert icon="URL" text="{{base_url}}/v1/players" />}}

{{<alert icon="Request" text="Body"  />}}

{ <br/>
    "challenges_id" : 1, <br/>
    "player_userid" : 123, <br/>
    "player_useremail" : "abc.com", <br/>
    "player_gamestartime" : "2017-09-05T20:12:59.152Z" <br/>
    
}<br/>

{{<alert icon="Response" text="Body"  />}}

{ <br/>
    "challenges_id": 1, <br/>
    "player_userid": 123, <br/>
    "player_useremail": "abc.com", <br/>
    "player_gamestartime": "2023-09-11T23:35:49.0340732+05:30", <br/>
    "player_endgametime": "2023-09-11T23:36:34.0340732+05:30", <br/>
    "total_questions": 1, <br/>
    "challenge_duration_seconds": 45, <br/>
    "player_time_elapsed": 0, <br/>
    "player_right_answers": 0, <br/>
    "player_correct_ans_percentage": 0 <br/>
} <br/>


### POST :  Player Create Records QnA

{{<alert icon="URL" text="{{base_url}}/v1/players" />}}

{{<alert icon="Request" text="Body"  />}}

{ <br/>
    "player_userid" : 2, <br/>
    "player_useremail" : "b.com", <br/>
    "challenges_id" : 1,<br/>
    "question_id_attempted" : 2 ,<br/>
    "player_answer" : "Lotus"<br/>

}<br/>

{{<alert icon="Respone" text="Body"  />}}

{
    "player_userid": 2,
    "player_useremail": "b.com",
    "challenges_id": 1,
    "question_id_attempted": 2,
    "player_answer": "Lotus",
    "question_id_attempted_is_correct": true,
    "answered_at": "2023-09-12T19:33:18.007685+05:30",
    "player_session_started": "2023-09-12T12:38:02.290115Z",
    "time_elapsed_from_session_to_current": 5115
}

<hr/>

### GET :  Player GET Records QnA

{{<alert icon="URL" text="{{base_url}}/v1/players/search?player_userid=2&challenges_id=1" />}}

{{<alert icon="Response" text="Body"  />}}

{ <br/>
    "id": 2,<br/>
    "challenges_id": 1,<br/>
    "total_questions": 2,<br/>
    "challenge_duration_seconds": 90,<br/>
    "player_userid": 2,<br/>
    "player_useremail": "b.com",<br/>
    "player_gamestartime": "2023-09-12T12:38:02.290115Z",<br/>
    "player_endgametime": "2023-09-12T12:39:32.290115Z",<br/>
    "player_game_submitted_at": "2023-09-12T13:04:12.8203Z",<br/>
    "player_time_elapsed": -1427845636,<br/>
    "player_right_answers": 1,<br/>
    "player_correct_ans_percentage": 0<br/>
}


### GET :  Player GET All Records QnA

{{<alert icon="URL" text="{{base_url}}/v1/records?player_userid=1&challenges_id=1" />}}

{{<alert icon="Response" text="Body"  />}}

[ <br/>
    { <br/>
        "id": 1,<br/>
        "challenges_id": 1,<br/>
        "player_userid": 1,<br/>
        "player_useremail": "a.com",<br/>
        "question_id_attempted": 1,<br/>
        "question_id_attempted_is_correct": {<br/>
            "Bool": true,<br/>
            "Valid": true<br/>
        },<br/>
        "player_answer": "Lotus",<br/>
        "answered_at": "2023-09-12T12:56:55.815805Z",<br/>
        "player_session_started": "2023-09-12T12:37:35.243329Z",<br/>
        "time_elapsed_from_session_to_current": 1160<br/>
    }<br/>
]<br/>

<hr/>


### POST :  Submit All Answers Challenge

{{<alert icon="URL" text="{{base_url}}/v1/records/submitAnswers" />}}

{{<alert icon="Request" text="Body"  />}}

{ <br/>
    "challenges_id" : 1,<br/>
    "player_userid" : 2<br/>
}<br/>

{{<alert icon="Response" text="Body"  />}}

{ <br/>
    "id": 2,<br/>
    "challenges_id": 1,<br/>
    "total_questions": 2,<br/>
    "challenge_duration_seconds": 90,<br/>
    "player_userid": 2,<br/>
    "player_useremail": "b.com",<br/>
    "player_gamestartime": "2023-09-12T12:38:02.290115Z",<br/>
    "player_endgametime": "2023-09-12T12:39:32.290115Z",<br/>
    "player_game_submitted_at": "2023-09-12T13:26:02.12752Z",<br/>
    "player_time_elapsed": -2085651016,<br/>
    "player_right_answers": 1,<br/>
    "player_correct_ans_percentage": 50<br/>
}<br/>

<hr/>


### GET :  Leaderboard List

{{<alert icon="URL" text="{{base_url}}/v1/players/leaderboard?challenges_id=1" />}}

{{<alert icon="Response" text="Body"  />}}

[ <br/>
    { <br/>
        "id": 2, <br/>
        "challenges_id": 1, <br/>
        "total_questions": 2, <br/>
        "challenge_duration_seconds": 90, <br/>
        "player_userid": 2, <br/>
        "player_useremail": "b.com", <br/>
        "player_gamestartime": "2023-09-12T12:38:02.290115Z"<br/>
        "player_endgametime": "2023-09-12T12:39:32.290115Z",<br/>
        "player_game_submitted_at": "2023-09-12T13:26:02.12752Z",<br/>
        "player_time_elapsed": -2085651016,<br/>
        "player_right_answers": 2,<br/>
        "player_correct_ans_percentage": 100<br/>
    }<br/>
    {<br/>
        "id": 1,<br/>
        "challenges_id": 1,<br/>
        "total_questions": 2,<br/>
        "challenge_duration_seconds": 90,<br/>
        "player_userid": 1,<br/>
        "player_useremail": "a.com",<br/>
        "player_gamestartime": "2023-09-12T12:37:35.243329Z",<br/>
        "player_endgametime": "2023-09-12T12:39:05.243329Z",<br/>
        "player_game_submitted_at": "2023-09-12T13:25:50.831984Z",<br/>
        "player_time_elapsed": 780696996,<br/>
        "player_right_answers": 1,<br/>
        "player_correct_ans_percentage": 50<br/>
    }<br/>
]<br/>
