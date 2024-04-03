package server.handler;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import server.Client;
import server.message.Message;
import server.message.MessageType;
import server.Server;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;

public class ClientTCPHandler implements Runnable {
    private Socket clientSocket;
    private Client client;
    private PrintWriter out;
    private BufferedReader in;
    private static final Logger logger = LogManager.getLogger(ClientTCPHandler.class);
    private final Server server;

    public ClientTCPHandler(Socket clientSocket,Server server) {
            this.clientSocket = clientSocket;
            this.server = server;


        }

    @Override
    public void run() {
        try
        {
            initConnection();
            String threadName = "TCP Client - " + clientSocket.getPort();
            Thread.currentThread().setName(threadName);
            logger.info("Client connected on port : " + clientSocket.getPort());
            while (clientSocket.isConnected()&& !Thread.currentThread().isInterrupted()) {
                String inputLine = in.readLine();
                if (inputLine != null) {
                    Message message = new Message(inputLine);
                    switch (message.getType()) {
                        case TEXT_TCP:
                            server.sendToAllTCP(client,message.getMessage());
                            logger.info("Received message from "+ client +" on port : " + inputLine +" "+ clientSocket.getPort());
                            break;
                        case QUIT:
                            logger.info("Client " + client + " disconnected on port : " + clientSocket.getPort());
                            server.removeClient(client);
                            break;
                        default:
                            break;
                    }

                }

            }
        }catch (IOException e) {
            logger.error("Exception caught when trying to listen on port "
                    + clientSocket.getPort() + " run" + e.getMessage());
        }

    }
    private void initConnection() {
        try {
            this.out = new PrintWriter(clientSocket.getOutputStream(), true);
            this.in = new BufferedReader(new InputStreamReader(clientSocket.getInputStream()));
            String inputLine = in.readLine();
            if (inputLine != null) {
                System.out.println("inputline: "+inputLine);
                Message message = new Message(inputLine);
                if (message.getType() == MessageType.INIT) {
                    String name = message.getText();
                    if(server.nameExists(name)){
                        sendQuitMessage();
                        logger.error("Name already exists");
                        return;
                    }
                    client = new Client(name,clientSocket);
                    server.addClient(client,out);
                    logger.info("Client " + name + " connected on port : " + clientSocket.getPort());
                    out.println("Welcome " + name);
                }

            }

        } catch (IOException e) {
            logger.error("Exception caught when trying to listen on port "
                    + clientSocket.getPort() + " initStreams ." + e.getMessage());
        }
    }

    public void sendQuitMessage() {
        out.println("server:quit|Name already exists");
        try {
            clientSocket.close();
        } catch (IOException e) {
            logger.error("Error closing socket in tcphandler send quit msg", e);
        }
    }
}
