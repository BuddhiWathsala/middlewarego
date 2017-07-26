#!/usr/bin/python           

import socket               

s = socket.socket(socket.AF_INET,socket.SOCK_STREAM)        
host = socket.gethostbyname("localhost")

port = 8081       

service = input("Enter service name: ")

if(service=="getservice2" or service == "getservice3"):
    parameters = input("Enter Parameters : ")
    parameters = parameters.replace(" ","%20")
    url = "/"+service + "/" + parameters
else:
     url = "/"+service
request = "GET "+url+" HTTP/1.1\nHost:"+host+"\n\n "

s.connect((host, port))
s.sendto(request.encode(),(host,port))

result = s.recv(4096)
finalresult = result.decode().split("\r")
print(finalresult[-1])
s.close()      
