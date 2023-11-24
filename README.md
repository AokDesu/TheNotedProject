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

>**Backend/:** This directory contains the Golang backend code.
>**main.go:** The entry point of your Golang application.
>**Handlers/:** Handles HTTP requests and contains Middleware and API endpoints.
>**Models/:** Defines data models and structures.
>**Databases/:** Manages database connections and migrations.
>**frontend/:** https://kit.svelte.dev/docs/project-structure
>**go.mod:** This file contain version of dependencies.
>**go.sum** This file is check sum of each dependencies version in go.mod.
### *Everything in this project is not complete yet*
>*PS. I make this project for practing golang. This is my first project. If you have any advices for me I'll be grateful to receive and improve myself thx.*
