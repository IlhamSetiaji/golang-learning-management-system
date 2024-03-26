
# Golang Learning Management System

Golang Learning Management System is a robust Learning Management System (LMS) built using the Go programming language (GoLang) and RabbitMQ messaging broker. It provides a comprehensive platform for educational institutions, organizations, and online course providers to efficiently manage their learning content, courses, and student interactions.


## Key Features

 - **Course Management**: Easily create, update, and organize courses with flexible content structures.
- **User Management**: Manage students, instructors, and administrative staff with authentication and authorization controls.
- **Content Delivery**: Efficiently deliver multimedia learning content such as videos, documents, and quizzes.
- **Real-time Notifications**: Utilize RabbitMQ for real-time notifications on course updates, submissions, and announcements.
- **Scalability**: Built with GoLang's concurrency features, ensuring scalability to handle large numbers of users and courses.
- **Analytics**: Gather insights into student progress, course effectiveness, and engagement metrics.
- **Customization**: Highly customizable to adapt to various educational contexts and institutional requirements.
- **Integration**: Seamless integration capabilities with existing Learning Management Systems and third-party services.
- **Security**: Implement robust security measures to safeguard user data and ensure compliance with privacy regulations.


## Auth API Reference

#### Register

```http
  POST /auth/register
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `name` | `string` | **Required**. Your Fullname |
| `username` | `string` | **Required** |
| `email` | `string` | **Required**. Your Valid Email |
| `password` | `string` | **Required**. **Min: 8** |
| `role_id` | `int` | **Required**. **Mentee: 1, Mentor: 2** |

#### Login

```http
  GET /auth/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `username`      | `string` | **Required** |
| `password`      | `string` | **Required**. **Min: 8** |

#### Me

```http
  GET /auth/me
```

| Authorization Type | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Bearer`      | `string` | **Required**. Token that you'll get from login |


## Configuration variables

To run this project, you will need to add the following environment variables to your config.json file

```json
{
  "app": {
    "name": "your-app-name"
  },
  "web": {
    "prefork": false,
    "port": 3000
  },
  "log": {
    "level": "debug",
    "output": "stdout"
  },
  "database": {
    "driver": "your-database-driver",
    "host": "your-database-host",
    "port": 1234,
    "username": "your-database-user",
    "password": "your-database-password",
    "name": "your-database-name",
    "pool": {
      "idle": 10,
      "max": 100,
      "lifetime": 300
    }
  },
  "rabbitmq": {
    "url": "amqp://guest:guest@localhost:5672/",
    "queue": "your-queue-name"
  },
  "jwt": {
    "secret": "your-jwt-secret"
  }
}
```


## Installation

Install this project using go

```bash
  git clone https://github.com/IlhamSetiaji/golang-learning-management-system.git
  cd golang-learning-management-system
  cp .config.example.json config.json
  go mod tidy
```

to run this project

```bash
go run main.go
```

To run, watch, and build this project

```bash
CompileDaemon -command="./golang-learning-management-system"
```

To migrate the database
```bash
go run ./cmd/migration/main.go
```

To run the consumer worker
```bash
go run ./cmd/worker/consumer.go
```
    
## Contributing

Contributions are always welcome! Because I'll code this project after my works are done

- Please make a new branch based on **master** branch.
- Naming convention for the branch is **prefix/suffix** with **prefix** is based on what the module and **suffix** what will you do. For example **feature/master-class**
- Create Pull Request from **your_branch** to **master** and add me as reviewer

Please adhere to this project's `code of conduct`.


## To-Do List

- ~Setup Base Project~
- ~Setup Base Config and Packages~
- ~Setup Authentication~
- Create User Management Feature
- Create Programs and Batches
- Create Master Class
- Create Material Feature
- Create Task Feature
- Create Quiz Feature
- Create Grading Feature
- Create Webinar Feature
- Create Certification Feature

