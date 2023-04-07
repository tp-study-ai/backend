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
> RESPONSE 
>>{\
> id\
> name\
> description\
> public_tests\
> private_tests\
> generated_tests\
> difficulty\
> cf_contest_id\
> cf_index\
> cf_points\
> cf_rating\
> cf_tags\
> time_limit\
> memory_limit_bytes\
> link\
> task_ru\
> input\
> output\
> note\
> }

> описание:
>>полчение задачи по id
>
> метод:
>> GET
> 
> URL
>> https://study-ai.ru/api/get_task_by_id
> 
> RESPONSE
>>{\
> id\
> name\
> description\
> public_tests\
> private_tests\
> generated_tests\
> difficulty\
> cf_contest_id\
> cf_index\
> cf_points\
> cf_rating\
> cf_tags\
> time_limit\
> memory_limit_bytes\
> link\
> task_ru\
> input\
> output\
> note\
> }

> описание:
>>отправка решения задачи
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/check_solution
>
> REQUEST
>>{\
> task_id\
> solution\
> }
> 
> RESPONSE
>>{\
> checkTime\
> buildTime\
> checkResult\
> checkMessage\
> testsPassed\
> testsTotal\
> lintSuccess\
> }

> описание:
>>регистрация
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/register
>
> REQUEST
>>{\
> username\
> password\
> }

> описание:
>>авторизация
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/login
>
> REQUEST
>>{\
> username\
> password\
> }

> описание:
>>logout
>
> метод:
>> POST
>
> URL
>> https://study-ai.ru/api/logout