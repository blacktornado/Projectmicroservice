FETCH TOP 2 TRACK FROM MUSIXMATCH API WITH A GIVEN LOCATION

Before Moving ON KINDLY GENERATE YOUR API KEY FROM HERE AND PUT IT IN all ***three .env(in root folder , in Lyricsmicroservice & Trackmicroservice folder) files with key MUMATCH_API_KEY***
***https://developer.musixmatch.com/signup***

***HTTP METHOD GET***
****localhost:8082/top-tracks-with-lyrics?location=in***

In root folder main.go hits endpoint /top-tracks-with-lyrics and calls gateway.GetTopTrackL handler(located in ApiGateway folder) through a logging middleware (this middleware resides in middleware folder)
This main.go server is running in "8082"

This handler(inside ApiGateway) 
1. Calls http://localhost:8080/top-tracks?location=" that is the server running inside TrackMicroservice folder .
   This folder(TrackMicroservice) has its independent main.go file that runs in 8080 port. The source GetTrackMicro inside TrackMicroservice/Api handles the request and returns a slice
2. After getting the response from the above Track Handler it again calls LyricsMicroservice to get the respective lyrics of the tracks that is received from step 1
   This Folder(LyricsMicroservice) has its independent main.go file that runs in 8081 port. The source GetLyricsMicro inside LyricsMicroservice/Api handles the request and returns a string
3. After getting both the response from different sources it appends and process the final result and write it back as response to the client.


Overall Logic

***Three Server running indpendently communicating through a gateway in a RESTFUL WAY.***

RUN THESE THREE SERVERS SEPARATELY

***from Root Folder run***
***go run main.go to initiate 8082 port***

***from separate terminal goto TrackMicroservice folder and run***
***go run main.go to initiate 8080 port***

***from another terminal goto Lyricsmicroservice and run***  
***go run main.go to initiate 8081 port***

***Open Postman and hit GET localhost:8082/top-tracks-with-lyrics?location=in***

Rooms for Improvement

1. Graceful shutdown
2. Time Constraint
3. Rate Limitter
4. Handling Concurrent Request
5. More on Error Handling
6. Persistent Logging
7. For Vertical scaling Microservice pattern
8. Dockerizing the Project



![NEW](https://github.com/blacktornado/Projectmicroservice/assets/8749326/de2c2129-73c7-463f-9309-322134900d8f)

 
