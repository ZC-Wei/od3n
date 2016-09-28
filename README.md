# On-Demand Dynamic Delivery Network (OD3N)
## 0.Background
E-commerce is a very poplar field. It grows rapidly and own an unpredictable marketing potential. However, there is an evident limitation—delivery. An Post could provide basic delivery service, but its user experience is that good and can’t satisfy quick daily business delivery. DHL, UPS, DPD own professional high quality delivery service, but their service cost is very high and inflexible. Almost small and medium business can’t afford to build their own delivery service.

## 1. Overview
On-Demand Dynamic Delivery Network (OD3N) is a platform that aim to providing a flexible, effective, intelligent, low-cost same-city delivery service. With crowdsourcing, OD3N try to find out an affordable delivery solution for small and medium business and provide a good delivery experience for public.
There are three kind of characters in the system—user, courier and customer.
-	User----publish new deliveries.
-	Courier----complete deliveries.
- Customer----benefit from entire OD3N service.
The project will allow an end user to publish an assignment. The assignment contains a delivery request. The crowd worker will be assigned a proper courier according to his status (e.g. availability, location…) and local road situation. Then system will provide route navigation to the courier to avoid traffic jam and getting lost.

## 2. Implementation
The system will be divided into client side project and server side project. Client side project focus on interaction of all characters in the system. Server side project establish stable continuous services and handle all the business logic. These two parts communicate directly.

### 2.1. Client Side Applications
In general, there three types of client: courier mobile client, user web and mobile client, customer message interface. They individually handle all the interaction from different type of characters. Their actions including send request, update status, upload location data and so on.

#### 2.1.1.	Courier Mobile Client
Responsibilities:
-	Identification of courier
-	Upload courier location and status
-	Receive delivery assignment
-	Update delivery status
-	Route navigation

#### 2.1.2.	User Web/Mobile Client
Responsibilities:
-	Create delivery assignment
-	Manage delivery assignment
-	Request courier

#### 2.1.3.	Customer Message Interface
Responsibilities:
-	Notify customer delivery status

### 2.2 Server Side Architecture
Server side project contains multiple node. There is a unique node called Control Node (CN). Other node called functional node. Each functional node provides individual service, such as real-time dispatch, payment, management, navigation and so on.

#### 2.2.1.	Control Node
This unique node is the gateway of whole server side system. It will handle all inbound and outbound traffic. According to the type of request, it would distribute requests to the appropriate functional node, and forward outgoing responses to outside. Meanwhile it’s also responsible for geolocation information update.

#### 2.2.2.	Real-time Node
Real-time Node (RTN) is core functional node. It provides the core functionality of system—dynamic dispatch. RTN contain delivery request queue and courier pool. For each delivery request, the dispatch algorithm will dynamically calculate the courier scores and rank them. And then chose the courier who has best score, assign the delivery to him. The score will be affect by multiple factors (e.g. delivery priority, courier availability, courier real-time location, vehicle type, road condition). It is also responsible for assignment status monitoring.
#### 2.2.3.	Business Node
Business Node will contain all other business logic. These business logics include user management, payment, system monitoring. Those logics will be package as several individual modules. Each module exposes public APIs to make sure all module could cooperate together. Module could be deployed and attached to the main system flexibly.
## 3.	Workflow
Main business logic: New delivery -> Delivery finish

## 4.	System Design Topology

## 5.	Tech Stack
- Web Application:      React.js + Node.js
- Mobile Application: 	React Native
- Server Side: 		      Node.js/Go + Mongo DB + Redis + Docker + Nginx

## 6.	Development Route Map

|Timeline|Phase|Guideline|
|---|---|---|
|Current|1.Prototype|a. Basic usable client b. Implement dispatch module b. Fake map modeling d. Achieve main workflow demo|
|N/A|2.Demonstrable|a. Assignment management b. Tracking system c. Real-world simulation|
|N/A|3.Advanced|Dynamic navigation and more…|


