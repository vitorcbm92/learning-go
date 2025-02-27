# learning-go
This is a simple project to learn go


# What will this first project do?
🔹 1️⃣ **Complete REST API (CRUD + JWT Authentication)**

📌 **What you'll learn:**
✅ Create REST APIs with Go (net/http or gin-gonic)  
✅ Connect to a database (PostgreSQL, MySQL, SQLite)  
✅ Implement JWT authentication  
✅ Deploy to AWS or Railway  

📜 **What it does:**

- A REST API to manage users, products, or orders  
- Supports CRUD operations (GET, POST, PUT, DELETE)  
- Protects endpoints with JWT authentication  

🔥 **Technologies:**

- Go (gin-gonic or fiber)  
- Database (gorm, PostgreSQL)  
- Docker for local development  
- Deployment on AWS Lambda, Railway, or Render  

🎯 **Extras:**

- Create a simple front-end (React or Vue)  
- Add WebSockets for real-time notifications

# How to start
    -- This project runs on a remote EC2 instance
    -- Considering that will know how to create a simple free EC2 instance on aws:
    -- open CMD in Windows and type:
    📌 ssh -i "<path_to_your_key>" ec2-user@<your_aws_instance_ip>
    📌 Example -> ssh -i "C:\Users\John\keys\aws\my_key.pem" ec2-user@99.999.999.99

# Installation history and notes

    # Installing postgres on instance:
    - dnf list postgresql* -> List the available packages on our EC2 instance available to be installed;
    - sudo dnf install -y postgresql16 postgresql16-server -> Installed PostgresSql 16;
    - sudo /usr/bin/postgresql-setup --initdb -> Initialized the database;
    - sudo systemctl enable postgresql
    sudo systemctl start postgresql
    sudo systemctl status postgresql -> Start and enable PostgresSql;
    - psql --version -> verify installation;
 

    # Configuring installed db:
    - sudo su - postgres -> Switch to postgres system user;
    - psql -> opens the database shell;
    - CREATE DATABASE go_crud -> Create a new database for our project;
    - CREATE USER go_user WITH ENCRYPTED PASSWORD 'strongpassword' -> Create user and password;
    - GRANT ALL PRIVILEGES ON DATABASE go_crud TO go_user -> Grants admin privileges;
    - \q -> Exit the shell;
    - exit -> Exit postgres system user session;




 

