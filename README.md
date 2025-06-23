# RECORDMAKER

## Flattens the contents of multiple lines between a record seperator  into one line

```
./recordmaker_mac
Usage: ./recordmaker_mac -file <inputFile> -rs <recordSeparator> [-output <outputFile>] [-ls <lineSeparator>] [-rstop <recordStop>]

Options:
  -file string
    	Path to input file
  -ls string
    	Line separator. Defines the character between each newline of the record that has been flattened into a single line (optional, must be in quotes)
  -output string
    	Path to output file (optional)
  -rs string
    	Record separator. Define where a record starts (required, must be in quotes)
  -rstop string
    	Defines where the record stops. Needs to be in quotes (optional)
```

## Executables available for download

```
recordmaker_linux_32		recordmaker_rpi_arm64		recordmaker_windows_64.exe
recordmaker_linux_64		recordmaker_rpi_armv6		
recordmaker_mac			recordmaker_rpi_armv7
recordmaker_mac_arm64		recordmaker_windows_32.exe
```

## Go source code
```
recordmakerv3.go
```
## Example File:
```
interface GigabitEthernet1/0/1
 switchport access vlan 999
 switchport mode access
 switchport voice vlan 1999
 spanning-tree portfast
!
interface GigabitEthernet1/0/2
 switchport access vlan 999
 switchport mode access
 switchport voice vlan 1999
 spanning-tree portfast
!
interface GigabitEthernet1/0/3
 switchport access vlan 999
 switchport mode access
 switchport voice vlan 1999
 spanning-tree portfast
```
#### Command ./recordmaker_mac -file file.txt -rs interface -rstop !
```
interface GigabitEthernet1/0/1  switchport access vlan 999  switchport mode access  switchport voice vlan 1999  spanning-tree portfast 
interface GigabitEthernet1/0/2  switchport access vlan 999  switchport mode access  switchport voice vlan 1999  spanning-tree portfast 
interface GigabitEthernet1/0/3  switchport access vlan 999  switchport mode access  switchport voice vlan 1999  spanning-tree portfast 
```


