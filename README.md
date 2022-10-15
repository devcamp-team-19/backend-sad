# SAD - Scan Area Danger
 Sad Area Danger was created to reduce crime rates and improve security in Indonesia by forming a healthy society.

## Tokopedia Devcamp - Team 19
- Fajar Muhammad Hamka
- Marcello Faria
- Ghozi Mahdi

## Techstacks Backend
- Golang
- Clean Architecture
- Gin
- JWT
- UUID
- godotoenv
- GORM
- PostgreSQL

## How to use it
- clone this repo
- set the env
- run `docker compose up` for database
- create database sad-db (should be same with env)
- run `go install`
- run `go mod tidy`
- run `go run .`

## Feature Backend (All Done Handle)
 ```
 Feature API Usecase:
 - Get All User
 - Register User
 - Login User
 - Get File
 - Upload File
 - Create Report
 - Get Reports
 - Create Comments
 - Get All Coments in 1 Posts
 - Voting Report 1 Post (can unvote / change like to unliked or otherwise)
 - Get Specific Votes
 ```
- Postman: [Collection Json](https://drive.google.com/file/d/1WIMDwdWNBd1YfBvnGpnwKwX8jzfSc3bY/view?usp=sharing)