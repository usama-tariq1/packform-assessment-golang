# packform Golang Api

This template should help get you started developing with Golang

Current Project is Based on Initial Api Platform that I made 
support external cli and based on Gorm and Gin-Gonic


## Project Setup
Clone project and run
```sh
go get
```
assuming go 1.18 or greater is already installed on system

## Add Create Database and Import data
Download sample database [Google Drive](https://drive.google.com/file/d/1Cq8chsUKfCugmbZBA81g1CEd7IIjE5_y/view?usp=sharing).

### Compile with 
This will also create migrations in database if tables are not already created
```sh
go run main.go
```
note that project does not support hot reload 