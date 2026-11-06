@echo off
chcp 65001 >nul
title FINAL_REI_PRO - 前端服务
cd /d "%~dp0front"

if not exist package.json (
  echo [错误] 没有找到 front\package.json
  pause
  exit /b 1
)

echo 正在启动前端服务...
echo 浏览器地址：http://127.0.0.1:5188
echo.

where pnpm >nul 2>nul
if %errorlevel%==0 (
  call pnpm dev
) else (
  echo 未找到 pnpm，改用 npm 启动。
  call npm run dev
)

echo.
echo 前端服务已经停止。
pause
