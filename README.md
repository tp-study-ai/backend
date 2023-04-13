# api

## Like

> описание:
>> поставить лайк задаче
>
> метод:
>> POST
>
> URL:
>> http://127.0.0.1:8000/api/like_task
>
>> REQUEST EXAMPLE
>>```json
>>{
>>  "task_id": 41
>>}
>>```
>
>> RESPONSE
>>```json
>>{
>>  "message": "лайк поставлен"
>>}
>>```

> описание:
>> поставить лайк задаче
>
> метод:
>> POST
>
> URL:
>> http://127.0.0.1:8000/api/delete_like
>
>> REQUEST EXAMPLE
>>```json
>>{
>>  "task_id": 41
>>}
>>```
>
>> RESPONSE
>>```json
>>{
>>  "message": "лайк удален"
>>}
>>```

> описание:
>> поставить лайк задаче
>
> метод:
>> GET
>
> URL:
>> http://127.0.0.1:8000/api/get_like_tasks
>
>> RESPONSE
>>```json
>>{
>>  
>>    "count_tasks": 1,
>>    "tasks_id_list": [
>>        41
>>    ],
>>    "tasks": [
>>        {
>>            "id": 41,
>>            "name": "505_B. Mr. Kitayuta's Colorful Graph",
>>            "description": "description",
>>            "public_tests": [
>>                "input",
>>                "output"
>>            ],
>>            "difficulty": 8,
>>            "cf_contest_id": 505,
>>            "cf_index": "B",
>>            "cf_points": 1000,
>>            "cf_rating": 1400,
>>            "cf_tags_ID": [
>>                10,
>>                12,
>>                13,
>>                20
>>            ],
>>            "cf_tags_RU": [
>>                "поиск в глубину и подобное",
>>                "дп",
>>                "системы непересекающихся множеств",
>>                "графы"
>>            ],
>>            "cf_tags_en": [
>>                "dfs and similar",
>>                "dp",
>>                "dsu",
>>                "graphs"
>>            ],
>>            "time_limit": 1,
>>            "memory_limit_bytes": 1,
>>            "link": "link",
>>            "task_ru": "task_ru",
>>            "input": "input",
>>            "output": "output",
>>            "note": "note"
>>        },
>>        {
>>            ...
>>        }
>>        ...
>>    ]
>>}
>>```

# Tasks

> описание:
>> получение посылок
>
> метод:
>> POST
>
> URL:
> > https://study-ai.ru/api/get_send_tasks
>
>> RESPONSE
>>```json
>>{
>>    "tasks": [
>>        {
>>            "id": 1,
>>            "user_id": 1,
>>            "task_id": 1,
>>            "check_time": 1.002,
>>            "build_time": 0,
>>            "check_result": 6,
>>            "check_message": "check_message",
>>            "tests_passed": 0,
>>            "tests_total": 31,
>>            "lint_success": false,
>>            "code_text": "#include <bits/stdc++.h> ...",
>>            "date": "12 апреля 2023, 12:14"
>>        },
>>        {
>>            ...
>>        },
>>        ...
>>    ]
>>}
>>```

> описание:
>> рекомендаций
>
> метод:
>> POST
>
> URL:
> > https://study-ai.ru/api/get_similar
>
>> REQUEST
>>```json
>>{
>>  string `json:"source_code"`
>>  string `json:"problem_url"`
>>  int    `json:"rating"`
>>  int    `json:"difficulty"`
>>}
>>```
>
>> REQUEST EXAMPLE
>>```json
>>{
>>  "source_code": "#include <iostream>\n ...",    
>>  "problem_url": "/contest/1900/problem/D",
>>  "rating": 1900,
>>  "difficulty": 2
>>}
>>```
>
>> RESPONSE
>>```json
>>{
>>  "tasks": [
>>        {
>>            "id": 1,
>>            "name": "name",
>>            "description": "description",
>>            "public_tests": [
>>                "input",
>>                "",
>>                "output",
>>                ""
>>            ],
>>            "difficulty": 1,
>>            "cf_contest_id": 1,
>>            "cf_index": "C",
>>            "cf_points": 1,
>>            "cf_rating": 1,
>>            "cf_tags_ID": [
>>                1
>>            ],
>>            "cf_tags_RU": [
>>                "cf_tags_RU"
>>            ],
>>            "cf_tags_en": [
>>                "cf_tags_en"
>>            ],
>>            "time_limit": 1,
>>            "memory_limit_bytes": 1,
>>            "link": "linku",
>>            "task_ru": "task_ru",
>>            "input": "input",
>>            "output": "output",
>>            "note": "note"
>>        },
>>        {
>>            ...
>>        },
>>        ...
>>    ]
>>}
>>```

> описание:
>> получение рандомной задачи
>
> метод:
>> GET
>
> URL:
> > https://study-ai.ru/api/get_task
> 
>> RESPONSE 
>>```json
>>{
>>  int      `json:"id"`
>>  string   `json:"name"`
>>  string   `json:"description"`
>>  []string `json:"public_tests"`
>>  []string `json:"private_tests"`
>>  []string `json:"generated_tests"`
>>  int      `json:"difficulty"`
>>  int      `json:"cf_contest_id"`
>>  string   `json:"cf_index"`
>>  float64  `json:"cf_points"`
>>  int      `json:"cf_rating"`
>>  []int32  `json:"cf_tags"`
>>  string   `json:"time_limit"`
>>  int      `json:"memory_limit_bytes"`
>>  string   `json:"link"`
>>  string   `json:"task_ru"`
>>  string   `json:"input"`
>>  string   `json:"output"`
>>  string   `json:"note"`
>>}
>>```

> описание:
>>полчение задачи по id
>
> метод:
>> GET
> 
> URL
>> https://study-ai.ru/api/get_task_by_id
>
>> RESPONSE
>>```json
>>{
>>  int      `json:"id"`
>>  string   `json:"name"`
>>  string   `json:"description"`
>>  []string `json:"public_tests"`
>>  []string `json:"private_tests"`
>>  []string `json:"generated_tests"`
>>  int      `json:"difficulty"`
>>  int      `json:"cf_contest_id"`
>>  string   `json:"cf_index"`
>>  float64  `json:"cf_points"`
>>  int      `json:"cf_rating"`
>>  []int32  `json:"cf_tags"`
>>  string   `json:"time_limit"`
>>  int      `json:"memory_limit_bytes"`
>>  string   `json:"link"`
>>  string   `json:"task_ru"`
>>  string   `json:"input"`
>>  string   `json:"output"`
>>  string   `json:"note"`
>>}
>>```

> описание:
>>отправка решения задачи
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/check_solution
>
>> REQUEST
>>```json
>>{
>>  int    `json:"task_id"`
>>  string `json:"solution"`
>>}
>>```
> 
>> RESPONSE
>>```json
>>{
>>  float64 `json:"checkTime"`
>>  float64 `json:"buildTime"`
>>  int     `json:"checkResult"`
>>  string  `json:"checkMessage"`
>>  int     `json:"testsPassed"`
>>  int     `json:"testsTotal"`
>>  bool    `json:"lintSuccess"`
>>}
>>```

# Авторизация регистрация юзер

> описание:
>>регистрация
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/register
>
>> REQUEST
>> ```json
>> {
>>   string `json:"username"`
>>   string `json:"password"`
>> }
>> ```
> 
>> RESPONSE
>>```json
>> {
>>    int    `json:"id"`
>>    string `json:"username"`
>> }
>> ```

> описание:
>>авторизация
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/login
>
>> REQUEST
>> ```json
>> {
>>   string `json:"username"`
>>   string `json:"password"`
>> }
>> ```
>
> > RESPONSE
>>```json
>> {
>>    int    `json:"id"`
>>    string `json:"username"`
>> }
>> ```

> описание:
>>logout
>
> метод:
>> GET
>
> URL
>> https://study-ai.ru/api/logout
> 
>>  RESPONSE
>> ```json
>> {
>>   string `json:"message"`
>> }
>>```

> описание:
>>полчение юзера по id
>
> метод:
>> get
>
> URL
>> https://study-ai.ru/api/get_user
>> ```
>
>> RESPONSE
>>```json
>> {
>>    int    `json:"id"`
>>    string `json:"username"`
>> }
>> ```