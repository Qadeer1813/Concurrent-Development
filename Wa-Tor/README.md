# Concurrent Development
## Wa-tor Simulation
****
**Student Name: Qadeer Hussain**

**Student ID: C00270632** 

**Lecture: Joesph Kehoe**
** **
Before running or compiling the Wa-tor simulation make sure that 
[Fyne](https://docs.fyne.io/started/) is set up correctly.

    To compile go code
        1. go mod init
        2. go run Wa-tor.go

****
## Description

**System Information**

    Operating System: POP OS/Windows 11
    Ram: 16Gb
    Storage: 1TB
    GPU: Nvidia 3050
    Threads: 12

**Documentation**

    To View Documentation below are steps to be followed if using Gold

    1. go install go101.org/golds@latest

    2. run the following: golds .

    3. Head over to this: github.com/Qadeer1813/Concurrent-Development/tree/main/Wa-Tor section of the document

    4. This will now display the documentation

    To View Documentation below are steps to be followed if using Go Doc
    
    1. go install golang.org/x/tools/cmd/godoc@latest

    2. change the package name from apckage main to package Wa-Tor
    (Warning: Once this package has beeen changed to Wa-Tor the code will no longer compile.)

    3. then run 

    4. godoc -http=:8080

    5. Go to Third Pirty from the Index
    
    6. In there searchh for Wa-Tor this will produce the necessary documentation

**What is Wa-tor**

    Wa-Tor is a population dynamics simulation devised by A. K. Dewdney and 
    presented in the December 1984 issue of Scientific American in a five-page 
    article entitled "Computer Recreations: Sharks and fish wage an ecological 
    war on the toroidal planet Wa-Tor".

**Implementation**

    The Development of the WA-TOR simulastion started with a sequential implmentation
    and then moved to a more complex version concurrent version.

    1. Sequential Implmentation
    Straightforward sequential version of the Wa-Tor simulation, 
    focusing on creating the core functionality

    2. Concurrent Implmentation
    After the completion of the sequential implementation, concurrency was introduced

****

**BenchMark Results**

    Simulations were ran for 250 iterations

    Wator Thread Speed Up	
	
    Number of Sharks	30
    Number of Fish	35
    Fish Breeding Time	7
    Shark Breeding Time	9
    Shark Starvation	7

![image](https://github.com/user-attachments/assets/d7a341ce-0740-4647-85ce-5476125814d6)

![image](https://github.com/user-attachments/assets/8c40a84d-6a03-46cd-a0ed-ec19730b5f5a)

## License 

Concurrent Development © 2024 by Qadeer Hussain is licensed 
under CC BY-NC-ND 4.0. To view a copy of this license, 
visit https://creativecommons.org/licenses/by-nc-nd/4.0/
    
