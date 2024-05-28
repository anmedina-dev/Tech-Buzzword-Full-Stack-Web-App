# Tech Buzzword Newsletter Full Stack Web App

Ever feel like you're out on the tech lingo? Look no further, Tech Buzzword Newletter is here!

Here you may find the entire codebase for the full stack web application.

Each service will be in it's own independent branch.

## Learning Objectives

1. Learn Go
2. Practice proper system design techniques through documentation and implementation
3. Have fun!

## Functional Objectives

1. Allow Users to learn a new tech buzzword every day through either our UI application, Email newsletter, and/or Text Newsletter.

## Tech Buzzword Newsletter System Diagram

<img src="images\Tech-Buzzword-Newsletter-System-Design-Diagram.png" />

## Technologies

1. ReactJS
2. Go
3. MongoDB Atlas
4. Twilio

## DB Layer

Because of the lack of need for a relational database in this web application, the web app will be using MongoDB Atlas for our DB needs. In our MongoDB Atlas cluster, we will have 3 seperate collections to store our data.

### Email Users Collection

In this collection will be stored user data for those who signed up to be apart of the email newsletter.

<b>Schema</b>
id: string
email: string

### Text Users Collection

In this collection will be stored user data for those who signed up to be apart of the text message newsletter.

<b>Schema</b>
id: string
number: string

### Tech Buzzword Collection

In this collection will be stored the tech buzzwords, their definitions, whether or not they have been sent, and date if they have been sent.

<b>Schema</b>
id: string
buzzword: string
definition: string
hasBeenSaid: boolean
date: date || null

## Service Layer

### Load Balancer

<b>Branch: </b> service/load-balancer
<b>Description: </b> Redirect HTTP requests from the Client Layer to the appropriate service.

### Email Service

<b>Branch: </b> service/email-service
<b>Description: </b> Service written in Go.
<b>Functions: </b>

1. Handle signing users up for the email newsletter.
2. Removing users if they choose to be removed from email newsletter.
3. Communicate with tech buzzword service to retrieve the tech buzzword of the day.

### SMS Service

<b>Branch: </b> service/sms-service
<b>Description: </b> Service written in Go.
<b>Functions: </b>

1. Handle signing users up for the sms newsletter.
2. Removing users if they choose to be removed from sms newsletter.
3. Communicate with tech buzzword service to retrieve the tech buzzword of the day.

### Tech Buzzword Service

<b>Branch: </b> service/tech-buzzword-service
<b>Description: </b> Service written in Go.
<b>Functions: </b>

1. Handle choosing tech buzzword of the day.
2. Send tech buzzword of the day via HTTP Request.
3. Update and handle tech buzzword collection.

### Cron Service

<b>Branch: </b> service/cron-service
<b>Description: </b> Service written in Go.
<b>Functions: </b>

1. Handle changing tech buzzword everyday.
2. Handle sending text message and email on new tech buzzword.

## Client Layer

Implement ReactJS UI application where users can

1. Sign up for / remove themselves from the email newletter.
2. Sign up for / remove themselves from the sms newsletter.
3. View the tech buzzword of the day.
4. View all previous tech buzzwords that have been already said.
