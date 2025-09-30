# Схема БД


```mermaid
---
config:
  theme: neutral
---
erDiagram
	direction TB
	statuses {
		int id PK ""  
		string name  ""  
	}
	orders {
		int id PK ""  
		int user_id FK ""  
		int status_id FK ""  
		int number  ""  
		numeric accrual  "numeric(15,2)"  
		datetime created_at  ""  
		datetime uploaded_at  ""  
	}
	users {
		int id PK ""  
		string login  ""  
		string password  "hash BCRYPT"  
		datetime created_at  ""  
		datetime uploaded_at  ""  
	}
	balances {
		int id PK ""  
		int user_id FK ""  
		numeric current  "numeric(15,2)"  
		numeric withdrawn  "numeric(15,2)"  
	}
	withdrawals {
		int id PK "Пока по связям в данной таблице вопрос "  
		int order_id  ""  
		datetime processed_at  ""  
	}
	users||--|{orders:"  "
	statuses||--|{orders:"  "
	balances||--||users:"  "
```