# middlewarego
  1. main.go => main file which work middleware
  1. server_01.go => 1st server : just print a message
  1. server_02.go => access a mongo database and give data when you request with a particular user name
  1. server_03.go => you can give any number of integers as input parameters which seperated by space. Then it gives the sum of all integers.
  1. registory.go => work as a registory to middleware. Middleware map request services with suitable functions with using this registory.
  1. servers.xml => store all information about servers
  1. client_1.py => python client. Access a port 8081 and retreive data with requesting services from middleware.
  URLConnDemo.java => java client. Working process same as python client.
