# The Noted Project - A Knowledge Management and Note-Taking Platform.
>This project is designed for note-taking and knowledge management, allowing users to create notes for various topics. It primarily focuses on backend development while prioritizing security measures for ensuring data safety. As a beginner, I'm using this project to practice Golang.

> **Project structure**
+ /project-root
    + backend/
        + main.go
        + handlers/
            + Middleware
              + CORS.go 
              + Auth 
                + auth.go
            + Routes
              + page.go
              + ...  
        + Models/
            + user.go
            + Note.go
            + ...
        + Databases/
            + database.go
            + table.go
            + ...
    + .env
    + frontend/
        + svelte
            + svelte defualt structure...
+ gitignore
+ go.mod
+ go.sum
+ README.md

>**Backend/:** This directory contains the Golang backend code.<br/>
>**main.go:** The entry point of your Golang application.<br/>
>**Handlers/:** Handles HTTP requests and contains Middleware and API endpoints.<br/>
>**Models/:** Defines data models and structures.<br/>
>**Databases/:** Manages database connections and migrations.<br/>
>**frontend/:** https://kit.svelte.dev/docs/project-structure.<br/>
>**go.mod:** This file contain version of dependencies.<br/>
>**go.sum** This file is check sum of each dependencies version in go.mod.<br/>
### *Everything in this project is not complete yet*
>*PS. I make this project for practing golang. This is my first project. If you have any advices for me I'll be grateful to receive and improve myself thx.*

==========================================================================================
# How to install
You need to install [Golang](https://go.dev/dl/) first, after you installed you can Go to next step.

### Step 1:
> *git clone https://github.com/AokDevkid/TheNotedProject.git* <br/>
use this command to clone this repo to you directory.

Once you've cloned the repository to your directory. create a **.env** file within the Backend directory. Open the **.env** file and add the following:
> DATABASE_URL = "your_database_here"<br/>
> DATABASE_PASSWORD = "your_database_password_here"<br/>
> SECRET_KEY = "your_secretkey_here"



### Step 2:
Open terminal and execute the following commands.
> cd Backend <br/>
> go run main.go

**And then BOOM!. (HOPE IT WORKS :D)**

