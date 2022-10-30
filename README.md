# uci-remote

toolkits running UCI chess engine on a remote machine

 - 'uciserver' running on a machine having UCI chess engine
 - 'uciproxy' works like a UCI chess engine by connecting uciserver

# uciserver

```
uciserver [ip]:port /path/to/engine
```

## examples

Listening port 7900 with engine ./stockfish (or ./stockfish.exe on Windows)

```
uciserver :7900 stockfish
```

Binding interface 192.168.2.3, Listening port 7900, with engine D:\Chess Engines\stockfish_15_win_x64_avx2.exe

```
uciserver 192.168.2.3:7900 "D:\Chess Engines\stockfish_15_win_x64_avx2.exe"
```

# uciproxy

```
uciproxy ip:port
```
or
```
copy uciproxy to <engine name>
create <engine name>.txt with ip:port
```

## examples

Emulate UCI chess engine by connecting a uciserver at 192.168.2.30:7900

```
uciproxy 192.168.2.30:7900
```

Many chess software doesn't expect an argument for chess engine.

Create remote_komodo13 proxy engine connecting to 192.168.2.30:7900 

```
copy uciproxy remote_komodo13
echo 192.168.2.30:7900 > remote_komodo13.txt
./remote_komodo13
```
(or on Windows)
```
copy uciproxy.exe remote_komodo13.exe
echo 192.168.2.30:7900 > remote_komodo13.txt
remote_komodo13
```

# use case (Windows)

## uciserver example (running multiple engines)

```
2 uci engines on a machine 192.168.2.30
D:\Chess Engines\stockfish_15_win_x64_avx2.exe
D:\Chess Engines\komodo-13.02-64bit-bmi2.exe 
```

Create a script as servers.bat
```
echo start uciserver :7900 "D:\Chess Engines\stockfish_15_win_x64_avx2.exe" >> servers.bat
echo start uciserver :7901 "D:\Chess Engines\komodo-13.02-64bit-bmi2.exe" >> servers.bat
```

Run the script to run servers
```
servers.bat
```

## uciproxy example

192.168.2.30 is the IP address of the machine running servers.bat.

Create a proxy stockfish engine
```
copy uciproxy.exe remote_stockfish.exe
echo 192.168.2.30:7900 > remote_stockfish.txt
```

Create a proxy komodo engine
```
copy uciproxy.exe remote_komodo.exe
echo 192.168.2.30:7901 > remote_komodo.txt
```

Run or add remote stockfish to chess application
```
remote_stockfish.exe
```

Run or add remote komodo to chess application
```
remote_komodo.exe
```

