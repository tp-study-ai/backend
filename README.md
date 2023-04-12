## api

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
>>            "code_text": "#include <bits/stdc++.h>\n\nusing namespace std;\n\nvector<int> proc(const vector<int> &v1, const vector<int> &v2, int sz) {\n  vector<int> ret;\n  int p1 = 0, p2 = 0;\n  for (int i = 0; i < sz ; i++) {\n    if (i % 2 == 0) {\n      if (p1 >= v1.size()) return vector<int>();\n      ret.push_back(v1[p1++]);\n    } else {\n      if (p2 >= v2.size()) return vector<int>();\n      ret.push_back(v2[p2++]);\n    }\n  }\n\n  for (int i = 0; i + 1 < sz; i++)\n    if (abs(ret[i] - ret[i+1]) != 1)\n      return vector<int>();\n\n  return ret;\n}\n\nvoid solve() {\n  int sz = 0;\n  vector<vector<int>> v(2);\n  for (int i = 0; i < 4; i++) {\n    int n; cin >> n;\n    sz += n;\n    while (n > 0) {\n      v[i%2].push_back(i);\n      --n;\n    }\n  }\n\n  auto r1 = proc(v[0], v[1], sz);\n  auto r2 = proc(v[1], v[0], sz);\n\n  if (r1.size() == sz) {\n    cout << \"YES\" << endl;\n    for (int r: r1)\n      cout << r << \" \";\n    cout << endl;\n  } else if (r2.size() == sz) {\n    cout << \"YES\" << endl;\n    for (int r: r2)\n      cout << r << \" \";\n    cout << endl;\n  } else {\n    cout << \"NO\" << endl;\n  }\n}\n\nint main() {\n  ios_base::sync_with_stdio(0);\n  cin.tie(0);\n  solve();\n  return 0;\n}",
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