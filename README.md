# Пример решения задачи с использованием Go
# Data Science Contest  

## Общая информация  

В рамках Sberbank Data Science Journey проходит Data Science Contest. Во время контеста участникам предлагается поработать с банковскими данными и решить несколько исследовательских задач: определить пол клиента по его финансовым тратам в первой задаче, предсказать общий оборот в той или иной категории услуг во второй задаче и предсказать траты каждого клиента в каждой из категорий в третьей. Итоговый рейтинг участника рассчитывается на основе сумме баллов по каждой из задач.  

Основные данные представляют из себя историю банковских транзакций, а также демографическую информацию по некоторой выборке клиентов (данные обезличены и специальным образом искажены).  

### Таблица transactions.csv
#### > 300 Mb, поэтому сохранена только выборка

Описание  

Таблица содержит историю транзакций клиентов банка за один год и три месяца.  

Формат данных  

customer_id,tr_datetime,mcc_code,tr_type,amount,term_id  
111111,15 01:40:52,1111,1000,-5224,111111  
111112,15 15:18:32,3333,2000,-100,11122233  
...  
Описание полей  

customer_id — идентификатор клиента;  
tr_datetime — день и время совершения транзакции (дни нумеруются с начала данных);  
mcc_code — mcc-код транзакции;  
tr_type — тип транзакции;  
amount — сумма транзакции в условных единицах со знаком; + — начисление средств клиенту (приходная транзакция), - — списание средств (расходная транзакция);  
term_id — идентификатор терминала;  

### Таблица customers_gender_train.csv

Описание  

Данная таблица содержит информацию по полу для части клиентов, для которых он известен. Для остальных клиентов пол необходимо предсказать в задаче A.  

Формат данных  

customer_id,gender  
111111,0  
111112,1  
...  
Описание полей  

customer_id — идентификатор клиента;  
gender — пол клиента; 0 — женский, 1 — мужской;  

### Таблица tr_mcc_codes.csv

Описание  

Данная таблица содержит описание mcc-кодов транзакций.  

Формат данных  

mcc_code;mcc_description  
1000;словесное описание mcc-кода 1000  
2000;словесное описание mcc-кода 2000  
...  
Описание полей  

mcc_code – mcc-код транзакции;  
mcc_description — описание mcc-кода транзакции.  

### Таблица tr_types.csv  

Описание  

Данная таблица содержит описание типов транзакций.  

Формат данных  

tr_type;tr_description  
1000;словесное описание типа транзакции 1000  
2000;словесное описание типа транзакции 2000  
...  
Описание полей  

tr_type – тип транзакции;  
tr_description — описание типа транзакции;  

## Задача A

Для клиентов, у которых неизвестен пол (которых нет в обучающей выборке customers_gender_train.csv, но которые есть в transactions.csv), необходимо предсказать вероятность быть мужского пола (значение 1).  

### Ожидаемый формат посылки решения

customer_id,gender  
1111111,0  
1111112,1  
1111113,0.2  
...  
