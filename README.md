# api

## chock mode

> описание:
>> получение ударного режима
>
> метод:
>> GET
>
> URL:
>> http://127.0.0.1:8000/api/shock_mode
>
>> RESPONSE
>>```json
>>{
>>  "today": false,
>>  "chock_mode": 2
>>}
>>```

> описание:
>> получение данных для календаря
>
> метод:
>> GET
>
> URL:
>> http://127.0.0.1:8000/api/calendar
>
>> RESPONSE
>>```json
>>{
>>  "days": [
>>    1,
>>    1,
>>    0,
>>    0,
>>    0,
>>    0,
>>    0,
>>    0,
>>    0,
>>    0,
>>    1,
>>    0,
>>    0,
>>    0,
>>    0,
>>    0,
>>    ...
>>  ]
>>}
>>```

> описание:
>> получение всех решенных задач
>
> метод:
>> GET
>
> URL:
>> http://127.0.0.1:8000/api/get_done_task
>
>> RESPONSE
>>```json
>>{
>>  "count_done_task": 2,
>>  "done_task": [
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

> описание:
>> получение всех начатых задач задач
>
> метод:
>> GET
>
> URL:
>> http://127.0.0.1:8000/api/get_done_task
>
>> RESPONSE
>>```json
>>{
>>  "count_done_task": 2,
>>  "done_task": [
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
>> получение посылок по id задачи
>
> метод:
>> POST
>
> URL:
> > https://study-ai.ru/api/get_send_tasks_by_task_id
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
>>  "id": 1,
>>  "name": "name",
>>  "description": "description",
>>  "public_tests": [
>>      "input",
>>      "2\nM\nXS\nXS\nM\n",
>>      "output",
>>      "0\n"
>>  ],
>>  "difficulty": 7,
>>  "cf_contest_id": 1000,
>>  "cf_index": "A",
>>  "cf_points": 0,
>>  "cf_rating": 1200,
>>  "cf_tags_ID": [
>>      21,
>>      23
>>  ],
>>  "cf_tags_RU": [
>>      "жадные алгоритмы",
>>      "реализация"
>>  ],
>>  "cf_tags_en": [
>>      "greedy",
>>      "implementation"
>>  ],
>>  "time_limit": 2,
>>  "memory_limit_bytes": 256000000,
>>  "link": "link",
>>  "task_ru": "task_ru",
>>  "input": "input",
>>  "output": "output",
>>  "note": "note"
>>}
>>```

> описание:
>>полчение задачи по id
>
> метод:
>> GET
> 
> URL
>> https://study-ai.ru/api/get_task_by_id?id=1
>
>> RESPONSE
>>```json
>>{
>>  "id": 1,
>>  "name": "name",
>>  "description": "description",
>>  "public_tests": [
>>      "input",
>>      "2\nM\nXS\nXS\nM\n",
>>      "output",
>>      "0\n"
>>  ],
>>  "difficulty": 7,
>>  "cf_contest_id": 1000,
>>  "cf_index": "A",
>>  "cf_points": 0,
>>  "cf_rating": 1200,
>>  "cf_tags_ID": [
>>      21,
>>      23
>>  ],
>>  "cf_tags_RU": [
>>      "жадные алгоритмы",
>>      "реализация"
>>  ],
>>  "cf_tags_en": [
>>      "greedy",
>>      "implementation"
>>  ],
>>  "time_limit": 2,
>>  "memory_limit_bytes": 256000000,
>>  "link": "link",
>>  "task_ru": "task_ru",
>>  "input": "input",
>>  "output": "output",
>>  "note": "note"
>>}
>>```

> описание:
>> сортировка задач
>
> метод:
>> GET
>
> URL
>> http://127.0.0.1:8000/api/tasks_list?page=0&min_rating=1000&max_rating=4000&sort=rating_up&tags=1,4
>
>> PARAMS  
>> page\
>> sort
>> * rating_up рэйтинг по возрастанию
>> * rating_down рэйтинг по убыванию
>> * min_rating
>> * max_rating
>> * во всех остальных случаях сортирует по id\
>>
>> tags
>> * 1
>> * 1,2
> 
>> RESPONSE
>>```json
>>{
>>  "task_count": 284,
>>  "tasks": [
>>      {
>>          "id": 4655,
>>          "name": "1170_F. Wheels",
>>          "description": "Polycarp has n wheels and a car with m slots for wheels. The initial pressure in the i-th wheel is a_i.\n\nPolycarp's goal is to take exactly m wheels among the given n wheels and equalize the pressure in them (then he can put these wheels in a car and use it for driving). In one minute he can decrease or increase the pressure in any (single) wheel by 1. He can increase the pressure no more than k times in total because it is hard to pump up wheels.\n\nHelp Polycarp and say what is the minimum number of minutes he needs to spend to equalize the pressure of at least m wheels among the given n wheels.\n\nInput\n\nThe first line of the input contains three integers n, m and k (1 ≤ m ≤ n ≤ 2 ⋅ 10^5, 0 ≤ k ≤ 10^9) — the number of wheels, the number of slots for wheels in a car and the number of times Polycarp can increase by 1 the pressure in a wheel.\n\nThe second line of the input contains n integers a_1, a_2, ..., a_n (1 ≤ a_i ≤ 10^9), where a_i is the pressure in the i-th wheel.\n\nOutput\n\nPrint one integer — the minimum number of minutes Polycarp needs to spend to equalize the pressure in at least m wheels among the given n wheels.\n\nExamples\n\nInput\n\n\n6 6 7\n6 15 16 20 1 5\n\n\nOutput\n\n\n39\n\n\nInput\n\n\n6 3 1\n4 8 15 16 23 42\n\n\nOutput\n\n\n8\n\n\nInput\n\n\n5 4 0\n5 5 5 4 5\n\n\nOutput\n\n\n0",
>>          "public_tests": [
>>              "input",
>>              "6 6 7\n6 15 16 20 1 5\n",
>>              "output",
>>              "\n39\n"
>>          ],
>>          "difficulty": 12,
>>          "cf_contest_id": 1170,
>>          "cf_index": "F",
>>          "cf_points": 0,
>>          "cf_rating": 0,
>>          "cf_tags_ID": [
>>              1,
>>              3,
>>              21
>>          ],
>>          "cf_tags_RU": [
>>              "*особая задача",
>>              "бинарный поиск",
>>              "жадные алгоритмы"
>>          ],
>>          "cf_tags_en": [
>>              "*special",
>>              "binary search",
>>              "greedy"
>>          ],
>>          "time_limit": 3,
>>          "memory_limit_bytes": 256000000,
>>          "link": "https://codeforces.com/contest/1170/problem/F?locale=ru",
>>          "task_ru": "",
>>          "input": "",
>>          "output": "",
>>          "note": ""
>>      },
>>      {
>>          ...
>>      },
>>      ... 
>>  ]  
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
>>  "task_id": 48,
>>  "solution": "#include <iostream> ..."
>>}
>>```
> 
>> RESPONSE
>>```json
>>{
>>  "checkTime": 0.622,
>>  "buildTime": 0,
>>  "checkResult": 3,
>>  "checkMessage": "main.cpp:8:29: warning: ...",
>>  "testsPassed": 0,
>>  "testsTotal": 31,
>>  "lintSuccess": false
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

> описание:
>> обновление usera
>
> метод:
>> POST
>
> URL
>> http://127.0.0.1:8000/api/update
>> ```
>
>> REQUEST
>> ```json
>> {
>>    "id": 1,
>>    "username": "yutfut2",
>>    "new_username": "yutfut2",
>>    "new_password": "yutfut3"
>> }
>> ```
>
>> RESPONSE
>>```json
>> {
>>    "id": 1,
>>    "username": "yutfut2"
>> }
>> ```
