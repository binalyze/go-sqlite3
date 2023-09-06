cd /D "%~dp0"

C:\msys64\usr\bin\bash -lc " "
IF %ERRORLEVEL% NEQ 0 GOTO :error

C:\msys64\usr\bin\bash -lc "pacman -S --needed --noconfirm wget mingw-w64-i686-toolchain mingw-w64-i686-autotools"
IF %ERRORLEVEL% NEQ 0 GOTO :error

C:\msys64\usr\bin\bash -lc "pacman -S --needed --noconfirm mingw-w64-x86_64-toolchain mingw-w64-x86_64-autotools"
IF %ERRORLEVEL% NEQ 0 GOTO :error


set CHERE_INVOKING=yes 
set MSYSTEM=MINGW64

C:\msys64\usr\bin\bash -lc "./build.sh"
IF %ERRORLEVEL% NEQ 0 GOTO :error

set MSYSTEM=MINGW32

C:\msys64\usr\bin\bash -lc "./build.sh"

:error
EXIT /B %ERRORLEVEL%