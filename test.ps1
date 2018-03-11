# enables bitlocker on the system volume
Enable-BitLocker -MountPoint "C:" -EncryptionMethod Aes256 -UsedSpaceOnly -RecoveryPasswordProtector

# gets the bitlocker protectors on the system drive
(Get-BitlockerVolume -MountPoint $env:SystemDrive).KeyProtector | ConvertTo-Json

# enable with tpm
Add-BitLockerKeyProtector -MountPoint $env:SystemDrive –TpmProtector

# if you need to partition the drive
BdeHdCfg.exe -target %SystemDrive% shrink -quiet –restart

# make sure the machine has a tpm enabled
(Get-WmiObject win32_tpm -Namespace root\cimv2\Security\MicrosoftTPM).isenabled() | Select-Object -ExpandProperty IsEnabled