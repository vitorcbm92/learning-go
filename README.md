# TODO improve readme file.
# learning-go
This is a simple project to learn Go.

## What will this first project do?

🔹 1️⃣ **Complete REST API (CRUD + JWT Authentication)**

📌 **What you'll learn:**
- ✅ Create REST APIs with Go (net/http or gin-gonic)  
- ✅ Connect to a database (PostgreSQL, MySQL, SQLite)  
- ✅ Implement JWT authentication  
- ✅ Deploy to AWS or Railway  

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

## How to start

This project runs on a remote EC2 instance. Here's how to get started:

1. Ensure you know how to create a simple free EC2 instance on AWS.
2. Open CMD on Windows and use the following command to SSH into your EC2 instance:

```bash
# Replace "<path_to_your_key>" with the path to your private key, and "<your_aws_instance_ip>" with your EC2 instance's public IP.
ssh -i "<path_to_your_key>" ec2-user@<your_aws_instance_ip>

# Example:
# ssh -i "C:\Users\John\keys\aws\my_key.pem" ec2-user@99.999.999.99
