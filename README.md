# eventslogger

Events logger is service, which tracks events in the system. Since micro services architecture works differently as monolithic application, we need some more sophisticated way, to gather and view events, which are happening inside our xdevices ecosystem. Simple log file is not enough, because micro services involved in process of creating whole application mainly do not know each other, and are working asynchronous (are communicating through rabbit messaging broker). Events logger keeps track of the processes, errors and warnings, which are happening inside our micro services.

---

Part of xdevices application - simple micro services-based system for monitoring. <br/>
**Developed for udemy course "Beginners guide to microservices with Go, Spring and RaspPi".**
