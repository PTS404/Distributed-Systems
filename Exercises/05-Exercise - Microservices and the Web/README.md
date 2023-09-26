# Week 05 - Microservices and the Web (Agata)

https://learnit.itu.dk/course/view.php?id=3022221#section-4


In this lecture, we discuss how we can make distributed systems do useful work for us, by creating different kind of Services: The world of RPC - Webservices, Microservices, Functions as a Service

## Mandatory Reading
  * Designing Data Intensive Applications, Chapter 1 - Scalable and Maintainable Applications, p. 3 - p. 22
  * Designing Data Intensive Applications, Chapter 4 - Encoding and evolution, p. 111- p. 143

## Supplemental Reading
   
   <!-- [![Watch the video](https://img.youtube.com/vi/S2osKiqQG9s/0.jpg)](https://www.youtube.com/watch?v=S2osKiqQG9s) -->

  * Martin Kleppmann - [YouTube Video](https://www.youtube.com/watch?v=S2osKiqQG9s)

  * A short, concise understandable [introduction to REST by IBM](https://developer.ibm.com/articles/ws-restful/)
  * Martin Fowler - [An Introduction to Micro Services](https://martinfowler.com/articles/microservices.html)
  * Research paper: [Architectural Debt in Microservices: A Case Study in a Large Company](https://www.researchgate.net/publication/331904375_Architectural_Technical_Debt_in_Microservices_A_Case_Study_in_a_Large_Company), Toledo, Martini, Przybyszewska et.al.
  * A bit of IT history, about the WOES of the old days before REST, [when we were dealing with the dreadful SOAP services](http://www.tbray.org/ongoing/When/200x/2004/09/18/WS-Oppo) (and they are still hanging around)
  * [Basic gRPC with Go tutorial](https://grpc.io/docs/languages/go/basics/)
  * You will need Swagger for the exercises: [https://swagger.io/tools/open-source/getting-started/](https://swagger.io/tools/open-source/getting-started/)

## Lecture Demos

  * Calling a RESTful API: [checking the weather at ITU](https://api.open-meteo.com/v1/forecast?latitude=55.6763&longitude=12.5681&hourly=temperature_2m)
  * Creating your own [webservice in Golang](https://github.itu.dk/agpr/DISYS/tree/master/restful)
  * Design a RESTful API using OpenAPI ([Swagger](https://swagger.io/)), a hello world example [here](https://app.swaggerhub.com/apis/themathmagician/HelloDisys/0.0.1#/default/hello)
  * Creating a gRPC service in Golang






## Exercises

<img src="https://media.giphy.com/media/13GIgrGdslD9oQ/giphy.gif" width=50%/>

**User Story for Exercises**

Imagine ITU wants to be even better at delivering awesome courses, and having fabulous teachers.
Let us help to reach this objective, by designing a REST api for a webservice, that will expose data about

  * ITU students and enrollments
  * course workloads for each student
  * course teachers, and their student popularity scores
  * student satisfaction ratings for courses

## Exercise 1

Let us create an RPC / gRPC service

  1. Define service endpoints **student**, **course**, **teacher** using gRPC
  2. Discuss whether the operations should use one way, or bi-directional streaming
  3. Implement the RPC / gRPC service in Golang, that exposes your course endpoint
  4. Consume the RPC / gRPC course service endpoint by creating a client in Golang

## Exercise 2
Let us examine the if our service is a micro service

  1. Discuss which endpoints of your web service \*could\* have a different lifecycle / supporting team. 
  2. Discuss which operations could be asynchronous
  3. Now, try to re-design your API to a set of microservices, and discuss your architectural choices.
  4. Explain the difference between what is an API and what is a micro service, based on your design

## Exercise 3 - optional 
 Let us create an API for a  RPC / REST service:

  1. Define service endpoints - **student, course, teacher** using the Swagger editor
  2. Discuss what operations should be using GET, PUT, POST, DELETE
  3. Implement an RPC / REST service in Golang, that exposes your course endpoint
  4. Consume the RPC / REST course service endpoint by creating a client in Golang

## Exercise 4
Let us debate, when to use REST over HTTP  versus gRPC over HTTP to implement the API for a service:

  1. Which of your implementations has a contract for service provider / consumer
  2. Which was easier to implement ? Where was it easier to implement a client? a server?
  3. Which do you think will be easier to change in the future? 
  4. Discuss, when gRPC should be favored, and when REST should be favored