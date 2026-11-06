@echo off
chcp 65001 >nul
cd /d "%~dp0back"
echo Starting backend at http://127.0.0.1:8091
go run .
echo.
echo Backend stopped. Exit code: %errorlevel%
pause
