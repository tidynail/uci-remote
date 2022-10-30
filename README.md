# uci-remote

toolkits running UCI chess engine on a remote machine

 - 'uciserver' running on a machine having UCI chess engine
 - 'uciproxy' works like a UCI chess engine by connecting uciserver

# usage example (Windows)

## Servers (running enigens)

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

Run the script
```
servers.bat
```

## Clients

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

192.168.2.30 is the IP address of the machine running servers.bat.

Run or add remote stockfish to chess application
```
remote_stockfish.exe
```

Run or add remote komodo to chess application
```
remote_komodo.exe
```

# uciserver usage

```
uciserver [ip]:port /path/to/engine
```

## examples

```
uciserver :7900 stockfish
```

open port 7900, will run ./stockfish (or ./stockfish.exe on Windows)

```
uciserver 192.168.2.3:7900 "D:\Chess Engines\stockfish_15_win_x64_avx2.exe"
```

bind interface 192.168.2.3, open port 7900, will run D:\Chess Engines\stockfish_15_win_x64_avx2.exe upon connection.

# uciproxy usage

```
uciproxy ip:port
```
or
```
copy uciproxy to <engine name>
create <engine name>.txt with ip:port
```

## examples

```
uciproxy 192.168.2.30:7900
```

emulate UCI chess engine by connecting a uciserver at 192.168.2.30:7900

```
copy uciproxy.exe remote_komodo13.exe
echo 192.168.2.30:7900 > remote_komodo13.txt
remote_komodo13
```
create remote_komodo13 proxy engine connecting to 192.168.2.30:7900 
