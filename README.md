# api

## difficulty

> описание:
>> оценить сложность задачи
>
> метод:
>> POST
>
> URL:
>> http://127.0.0.1:8000/api/set_difficulty
>
>> -1 легко\
>> 1 сложно
> 
>> REQUEST EXAMPLE
>>```json
>>{
>>  "task_id": 41,
>>  "difficulty": -1
>>}
>>```
>
>> RESPONSE
>>```json
>>{
>>  "message": "оценка поставленна"
>>}
>>```

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
>>  {
>>  "username": "yf",
>>  "password": "yf"
>>}
>> ```
> 
>> RESPONSE
>>```json
>> {
>>  "id": 2,
>>  "username": "yf"
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

## Recommendation

>>> описание:
>> страница рекомендаций
>
> метод:
>> POST
>
> URL
>> http://127.0.0.1:8000/api/recommendations
>> ```
>
>> RESPONSE
>>```json
>>{
>>  "rec": [
>>      {
>>          "recommended_tag": "тернарный поиск",
>>          "priority": 1,
>>          "problems": [
>>              {
>>                  "id": 3886,
>>                  "name": "304_D. Rectangle Puzzle II",
>>                  "description": "You are given a rectangle grid. That grid's size is n × m. Let's denote the coordinate system on the grid. So, each point on the grid will have coordinates — a pair of integers (x, y) (0 ≤ x ≤ n, 0 ≤ y ≤ m).\n\nYour task is to find a maximum sub-rectangle on the grid (x1, y1, x2, y2) so that it contains the given point (x, y), and its length-width ratio is exactly (a, b). In other words the following conditions must hold: 0 ≤ x1 ≤ x ≤ x2 ≤ n, 0 ≤ y1 ≤ y ≤ y2 ≤ m, <image>.\n\nThe sides of this sub-rectangle should be parallel to the axes. And values x1, y1, x2, y2 should be integers.\n\n<image>\n\nIf there are multiple solutions, find the rectangle which is closest to (x, y). Here \"closest\" means the Euclid distance between (x, y) and the center of the rectangle is as small as possible. If there are still multiple solutions, find the lexicographically minimum one. Here \"lexicographically minimum\" means that we should consider the sub-rectangle as sequence of integers (x1, y1, x2, y2), so we can choose the lexicographically minimum one.\n\nInput\n\nThe first line contains six integers n, m, x, y, a, b (1 ≤ n, m ≤ 109, 0 ≤ x ≤ n, 0 ≤ y ≤ m, 1 ≤ a ≤ n, 1 ≤ b ≤ m).\n\nOutput\n\nPrint four integers x1, y1, x2, y2, which represent the founded sub-rectangle whose left-bottom point is (x1, y1) and right-up point is (x2, y2).\n\nExamples\n\nInput\n\n9 9 5 5 2 1\n\n\nOutput\n\n1 3 9 7\n\n\nInput\n\n100 100 52 50 46 56\n\n\nOutput\n\n17 8 86 92",
>>                  "public_tests": [
>>                      "input",
>>                      "100 100 52 50 46 56\n",
>>                      "output",
>>                      "17 8 86 92\n",
>>                      "input",
>>                      "9 9 5 5 2 1\n",
>>                      "output",
>>                      "1 3 9 7\n"
>>                  ],
>>                  "difficulty": 10,
>>                  "cf_contest_id": 304,
>>                  "cf_index": "D",
>>                  "cf_points": 500,
>>                  "cf_rating": 1700,
>>                  "cf_tags_ID": [
>>                      23,
>>                      25
>>                  ],
>>                  "cf_tags_RU": [
>>                      "реализация",
>>                      "математика"
>>                  ],
>>                  "cf_tags_en": [
>>                      "implementation",
>>                      "math"
>>                  ],
>>                  "time_limit": 2,
>>                  "memory_limit_bytes": 256000000,
>>                  "link": "https://codeforces.com/contest/304/problem/D?locale=ru",
>>                  "short_link": "",
>>                  "name_ru": "D. Прямоугольная загадка II",
>>                  "task_ru": "<div>\n<p>Дана прямоугольная сетка размера <span class=\"tex-span\"><i>n</i> × <i>m</i></span>. Введем систему координат на сетке. Итак, каждая точка на сетке имеет координаты — пару целых чисел <span class=\"tex-span\">(<i>x</i>, <i>y</i>)</span> <span class=\"tex-span\">(0 ≤ <i>x</i> ≤ <i>n</i>, 0 ≤ <i>y</i> ≤ <i>m</i>)</span>.</p>\n<p>Ваша задача — найти наибольший подпрямоугольник на сетке <span class=\"tex-span\">(<i>x</i><sub class=\"lower-index\">1</sub>, <i>y</i><sub class=\"lower-index\">1</sub>, <i>x</i><sub class=\"lower-index\">2</sub>, <i>y</i><sub class=\"lower-index\">2</sub>)</span>, содержащий данную точку <span class=\"tex-span\">(<i>x</i>, <i>y</i>)</span>, такой, что соотношение длин его сторон равняется <span class=\"tex-span\">(<i>a</i>, <i>b</i>)</span>. Иными словами, должны выполняться следующие условия: <span class=\"tex-span\">0 ≤ <i>x</i><sub class=\"lower-index\">1</sub> ≤ <i>x</i> ≤ <i>x</i><sub class=\"lower-index\">2</sub> ≤ <i>n</i></span>, <span class=\"tex-span\">0 ≤ <i>y</i><sub class=\"lower-index\">1</sub> ≤ <i>y</i> ≤ <i>y</i><sub class=\"lower-index\">2</sub> ≤ <i>m</i></span>, <img align=\"middle\" class=\"tex-formula\" src=\"https://espresso.codeforces.com/0c36d1eac2611e4dc89474c61ba622f0db130f00.png\" style=\"max-width: 100.0%;max-height: 100.0%;\">.</p>\n<p>Стороны этого подпрямоугольника должны быть параллельны осям координат. Величины <span class=\"tex-span\"><i>x</i><sub class=\"lower-index\">1</sub>, <i>y</i><sub class=\"lower-index\">1</sub>, <i>x</i><sub class=\"lower-index\">2</sub>, <i>y</i><sub class=\"lower-index\">2</sub></span> должны быть целыми.</p>\n<center> <img class=\"tex-graphics\" src=\"https://espresso.codeforces.com/ceee17be587b1dfbe1ae39292474b973fd250da4.png\" style=\"max-width: 100.0%;max-height: 100.0%;\"> </center>\n<p>Если существует несколько ответов, найдите ближайший к <span class=\"tex-span\">(<i>x</i>, <i>y</i>)</span> подпрямоугольник. Здесь «ближайший» означает, что Евклидово расстояние между <span class=\"tex-span\">(<i>x</i>, <i>y</i>)</span> и центром прямоугольника как можно меньше. Если все равно существует несколько ответов, выведите лексикографически минимальный. Здесь «лексикографически минимальный» означает, что мы должны рассматривать подпрямоугольник как последовательность целых чисел <span class=\"tex-span\">(<i>x</i><sub class=\"lower-index\">1</sub>, <i>y</i><sub class=\"lower-index\">1</sub>, <i>x</i><sub class=\"lower-index\">2</sub>, <i>y</i><sub class=\"lower-index\">2</sub>)</span>, так, что можно выбрать из них лексикографически минимальную.</p>\n</div>\n",
>>                  "input": "<div class=\"input-specification\">\n<div class=\"section-title\">Входные данные</div>\n<p>В первой строке записано шесть целых чисел <span class=\"tex-span\"><i>n</i>, <i>m</i>, <i>x</i>, <i>y</i>, <i>a</i>, <i>b</i></span> <span class=\"tex-span\">(1 ≤ <i>n</i>, <i>m</i> ≤ 10<sup class=\"upper-index\">9</sup>, 0 ≤ <i>x</i> ≤ <i>n</i>, 0 ≤ <i>y</i> ≤ <i>m</i>, 1 ≤ <i>a</i> ≤ <i>n</i>, 1 ≤ <i>b</i> ≤ <i>m</i>)</span>.</p>\n</div>\n",
>>                  "output": "<div class=\"output-specification\">\n<div class=\"section-title\">Выходные данные</div>\n<p>Выведите четыре целых числа <span class=\"tex-span\"><i>x</i><sub class=\"lower-index\">1</sub>, <i>y</i><sub class=\"lower-index\">1</sub>, <i>x</i><sub class=\"lower-index\">2</sub>, <i>y</i><sub class=\"lower-index\">2</sub></span>, обозначающие обнаруженный вложенный прямоугольник с левым нижним углом в <span class=\"tex-span\">(<i>x</i><sub class=\"lower-index\">1</sub>, <i>y</i><sub class=\"lower-index\">1</sub>)</span>, а правым верхним — в <span class=\"tex-span\">(<i>x</i><sub class=\"lower-index\">2</sub>, <i>y</i><sub class=\"lower-index\">2</sub>)</span>.</p>\n</div>\n",
>>                  "note": ""
>>              },
>>              {
>>                  ...
>>              },
>>              ...
>>          ]
>>      },
>>      {
>>          "recommended_tag": "кратчайшие пути",
>>          "priority": 2,
>>          "problems": [
>>              {
>>                  ...
>>              },
>>              ...
>>          ]
>>      },
>>      {
>>            ...
>>      }
>>  ]
>>}
>> ```
