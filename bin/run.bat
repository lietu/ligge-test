@echo off
go build -i ligge-test.go

if %errorlevel% neq 0 exit /b %errorlevel%

ligge-test.exe
