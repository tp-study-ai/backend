## api

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