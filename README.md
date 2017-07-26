# middlewarego
main working file
  main.go => main file which work middleware
  server_01.go => 1st server : just print a message
  server_02.go => access a mongo database and give data when you request with a particular user name
  server_03.go => you can give any number of integers as input parameters which seperated by space. Then it gives the sum of all integers.
  registory.go => work as a registory to middleware. Middleware map request services with suitable functions with using this registory.
  
  servers.xml => store all information about servers
  
  client_1.py => python client. Access a port 8081 and retreive data with requesting services from middleware.
  URLConnDemo.java => java client. Working process same as python client.
