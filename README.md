# Welcome to F Society Chat 

This is a simple command-line based chat application written in Go. It allows users to connect using `netcat` and share encrypted & anonymous messages in a terminal environment. This is our playground.

## Setup

1. Clone this repository to your local machine.

2. Open the `main.go` file and locate the `password` and `port` constants at the top of the file:

   ```go
   const (
       password = "exploit" 
       port     = "1337"   
   )
   ```
   Change the password to the desired password for authentication and adjust the port if necessary.
   
   Compile and run the server by executing the following command in your terminal:
      ```go
      go run main.go
      ```

## Connecting

1. Open a new terminal

2. Connect to the server using one of the following methods:

   - **netcat Command**:

     ```
     nc <server_public_ip> <port>
     ```

   - **PuTTY (Windows)**: If you're on Windows, select raw connection type and fill in the host's IP and port in PuTTY's configuration.
   Alternatively, run cmd and enter:
       ```
      putty -raw <server_public_ip> <port>
       ```

   - **SSH (Secure Shell)**: Not supported yet.

## Closing Connection

   1. Execute the Ctrl+C command.
   2. enter `clear` to clear the terminal

