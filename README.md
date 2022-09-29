# Go Svelte Chat App

It is a full stack project of the WebSocket chat application. <br>For backend I used <b>Go, Gorilla WebSocket, MongoDB.</b><br> <b>Svelte</b> was used for the frontend (it was my first project built with it).<br>
The user enters his nickname and joins the chat. The user gets the chat history. Other users are notified about the new user.<br><br>

## How to run it:
To run the application, you must have the <b>Docker</b> installed and turned on. Then download this project. You can delete the Illustrations folder, it was only used on gitchub for presentation purposes and it is not needed for the program to run. Then open the folder in the command prompt and enter the command:<br>
<b>docker compose build<b> <br><br>
and after it finishes downloading:<br>
<b>docker compose up</b><br><br>
To open the application, open the browser and enter the address:<br><br>
<b>http://localhost:5173/</b> <br><br>
Additionally, you can manage the database using Mongo Express, for this you need to open the browser with the address:<br><br>
<b>http://localhost:8081/</b> <br><br>

## Table of contents:
* [Technologies](#technologies)
* [Illustrations](#illustrations)

## Technologies
Project is created with:<br>
Backend:
* Go
* Gorilla Mux WebSocket
* MongoDB
* Mongo Express

Frontend:
* Svelte

## Illustrations
<p float="left">
 <img src="Illustrations/1.png" height = "200">
 <img src="Illustrations/2.png" height = "200">
 <img src="Illustrations/3.png" height = "200">
 <img src="Illustrations/4.png" height = "200">
 <img src="Illustrations/5.png" height = "200">
 <img src="Illustrations/6.png" height = "200">
 <img src="Illustrations/7.png" height = "200">
 <img src="Illustrations/8.png" height = "200">
 <img src="Illustrations/9.png" height = "200">
</p>