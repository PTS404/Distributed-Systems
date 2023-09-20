# Distributed-Systems

### a) What are packets in your implementation? What data structure do you use to transmit data and meta-data?
The packets aren't specifically defined but are implemented within a buffer (byte slice), that transmits the data between our client and server. The net package offers the net.conn interface, which handles information in the background.  

The data structure used for transmitting data and meta-data is byte slices ([]byte, <buffer siZe>)

### b) Does your implementation use threads or processes? Why is it not realistic to use threads?
We are using threads to run client and server simultaneously instead of using the nc command to talk to the tcp server. This is because we don't have linux along wanting to provide a simplified way of using/testing the program.

It is not realistic to use threads since the protocol should run across a network instead of locally. 

### c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
Using sequence numbers. Gives certain numbers to part of a message (requiring the message to be in more than one part), and then put them in the right sequence.

### d) In case messages can be delayed or lost, how does your implementation handle message loss?
This could be solves using either acknowledgements, timeouts or re-transmissions.
Acknowledgements works by waiting until the message has arrived, this can give some errors, since it will keep waiting. This can be solves working with the timeout, where after a certain amount of time, we assume that the message was never sent. 
When assuming the message was never sent, we can use re-transmissions to send the request/message again, and then going through the process again. 

### e) Why is the 3-way handshake important?
To make sure both the client and the server are agreeing. Both parts need to acknowledge the message. When both the client and server has a connection, then the TCP can move on, they need to have established connection.
