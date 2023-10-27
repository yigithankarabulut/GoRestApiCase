# GolangRestApiCase

## Endpoints
```http
POST    /register
POST    /login
POST    /logout
PUT     /update

POST    /plan/create
GET     /plan/get/?planNumber=?
GET     /plan/getByState/?state=?

PUT     /plan/cancel/?planNumber=?
PUT     /plan/complete/?planNumber=?
PUT     /plan/start/?planNumber=?

PUT     /plan/update/?planNumber=?
DELETE  /plan/delete/?planNumber=?

GET     /plan/listAll
POST    /plan/listWeekly/
POST    /plan/listMonthly?lastweek=?
```

## Json Templates
### `/register`
```json
{
    "name": "your_name_here",
    "lastname": "your_lastname_here",
    "password": "your_password_here",
    "student_number": "your_student_number_here"
}
```

### `/login`
```json
{
    "student_number": 42,
    "password": "your_password_here"
}
```
### `/update`
```json
{
    "password": "your_current_password_here",
    "newPassword": "your_new_password_here"
}
```
### `/plan/create`
```json
{
    "plan_number": 321,
    "plan_description": "your_plan_description_here",
    "date": "26.10.2024",
    "start_hour": "20:45",
    "end_hour": "21:55",
    "status": "created"
}
```
### `/plan/update/?planNumber`
```json
{
    "plan_number": 320,
    "plan_description": "your_updated_plan_description_here",
    "date": "26.11.2023",
    "start_hour": "20:55",
    "end_hour": "20:55",
    "status": "cancelled"
}
```

### `/plan/listMonthly`
```json
{
    "month": 10,
    "year": 2023
}
```
