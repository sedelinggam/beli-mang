---
runme:
  id: 01HX76ZK1SK79VM7CJ415Y58ZT
  version: v3
---

# 🍔 Beli Mang

BeliMang! is a food delivery app that user can use to buy food and drinks!

## 🌟Features

BeliMang offers the following features:

- **Admin Authentication**:
- Admin registration
- Admin login
- **User Authentication**:
- User registration
- User login
- **Image Upload**:
- Upload Image
- **Manage Merchant**:
- Create Merchant
- Get Merchants
- Create Merchant's Items
- Get Merchant's Items
- **Purchase**:
- Get Merchant Location
- Create Price and Delivery Estimation
- Create an Order
- Get Order History

## ⛔️ Requirements

Before running this app make sure you have this installed on your puter :

- [Golang 1.22.0](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [docker](https://docs.docker.com/engine/install/ubuntu/)
- [AWS S3](https://aws.amazon.com/s3/)
- [AWS ECR](https://aws.amazon.com/ecr/)
- [AWS ECS](https://aws.amazon.com/ecs/)

## 🎖Prerequisite

To run the application, follow these steps before run the program:

1. Make sure you have Golang, PostgreSQL, Golang Migrate, and Docker installed and configured on your system.
2. Clone this repository:

```bash {"id":"01HXBJ7XEECXDYSM92BBJFY4V5"}

git clone https://github.com/sedelinggam/beli-mang.git

```

3. Navigate to the project directory:

```bash {"id":"01HXBJ7XEECXDYSM92BC18F9P1"}

cd beli-mang

```

4. Run the following command to install dependencies:

```bash {"id":"01HXBJ7XEECXDYSM92BDDP7D4A"}

go mod download

```

5. Run the following command to create environment for the application:

```bash {"id":"01HXBJ7XEECXDYSM92BG3X43GZ"}

mv .env.sample .env

```

## 🚀 Run The Program

1. **Setting Up Environment Variables**

Before starting the application, you need to set up the following environment variables:

- `DB_NAME`: Name of your PostgreSQL database
- `DB_PORT`: Port of your PostgreSQL database (default: 5432)
- `DB_HOST`: Hostname or IP address of your PostgreSQL server
- `DB_USERNAME`: Username for your PostgreSQL database
- `DB_PASSWORD`: Password for your PostgreSQL database
- `DB_PARAMS`: Additional connection parameters for PostgreSQL (e.g., sslmode=disabled)
- `JWT_SECRET`: Secret key used for generating JSON Web Tokens (JWT)
- `BCRYPT_SALT`: Salt for password hashing (use a higher value than 8 in production!)
- `AWS_ACCESS_KEY_ID`: AWS Access Key ID
- `AWS_SECRET_ACCESS_KEY`: AWS Secret Access Key
- `AWS_S3_BUCKET_NAME`: AWS S3 Bucket Name
- `AWS_REGION`: AWS Region (e.g. ap-southeast-1)

2. **Database Migrations**

- Apply migrations to the database:

```bash {"id":"01HXBJ7XEECXDYSM92BKBXS47Z"}

make migrate-dev

```

3. **Running the Application**

```bash {"id":"01HXBJ7XEECXDYSM92BNS4FSD8"}

make run

```

You can access the application in your web browser at http://localhost:8080

## 🐋 Build Image

Make sure you already installed Docker on your computer.
Adjust the `.env` file to make sure it's connected to the database and then you can build the Docker image by running:
```bash
make build-image
```

After the image is build, adjust the `.env` and you can run it using:
```bash
make run-image
```