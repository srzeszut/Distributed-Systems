package server;

import java.io.IOException;
import java.io.PrintWriter;
import java.net.DatagramPacket;
import java.net.DatagramSocket;
import java.net.ServerSocket;
import java.net.Socket;
import java.util.*;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import server.handler.ClientTCPHandler;
import server.handler.ClientUDPHandler;

import static java.lang.Thread.sleep;

public class Server {

    private ServerSocket serverTCPSocket;
    private DatagramSocket serverUDPSocket;
    private static final Logger logger = LogManager.getLogger(Server.class);


    private int portNumber;
    private Map<Thread, Socket> clientThreadsMap = new HashMap<>();
    private List<Client> clients = new ArrayList<>();
    private Map<Client, PrintWriter> clientWriters = new HashMap<>();

    public Server(int portNumber) {
        this.portNumber = portNumber;

    }


    public void start() {
        System.out.println("Server started on port " + portNumber);
        try {
            serverTCPSocket = new ServerSocket(portNumber);
            serverUDPSocket = new DatagramSocket(portNumber);
            Thread clientUDPThread = new Thread(new ClientUDPHandler(serverUDPSocket,this));
            clientUDPThread.start();

            while (true) {
                        Socket clientSocket = serverTCPSocket.accept();
                        Thread clientTCPThread = new Thread(new ClientTCPHandler(clientSocket,this));


                        clientThreadsMap.put(clientTCPThread, clientSocket);

                        clientTCPThread.start();



            }

        } catch (IOException e) {
            logger.error("Could not listen on port " + portNumber, e);

        }
    }

    public synchronized void sendToAllTCP(Client from,String message) {

        for (Client client : clientWriters.keySet()) {
            if (!client.equals(from)) {
                clientWriters.get(client).println(message);
                logger.info("Sending message to " + client);
            }
            }
        }
    public synchronized void sendToAllUDP(Client from,byte[] sendData) {

        for(Client client : clients) {
            if (!client.equals(from)) {
                DatagramPacket sendPacket = new DatagramPacket(sendData, sendData.length, client.socket().getInetAddress(), client.socket().getPort());
                try {
                    serverUDPSocket.send(sendPacket);
                    logger.info("Sending UDP packet to " + client);
                } catch (IOException e) {
                    logger.error("Error sending UDP packet", e);
                }
            }

        }
    }

    public synchronized void sendQuitMessage(Client client) {
        clientWriters.get(client).println("server:quit");
        try
        {
            serverUDPSocket.send(new DatagramPacket("server:quit".getBytes(), "server:quit".length(), client.socket().getInetAddress(), client.socket().getPort()));
            client.socket().close();

        } catch (IOException e) {
            logger.error("Error closing socket", e);
        }

    }

    public synchronized void addClient(Client client, PrintWriter out) {
        clients.add(client);
        clientWriters.put(client, out);
    }
    public synchronized void addClient(Client client) {
        clients.add(client);
    }
    public synchronized void removeClient(Client client) {
        sendQuitMessage(client);
        clients.remove(client);
        clientWriters.remove(client);
        Iterator<Map.Entry<Thread, Socket>> iterator = clientThreadsMap.entrySet().iterator();
            while (iterator.hasNext()) {
                Map.Entry<Thread, Socket> entry = iterator.next();
                if (entry.getValue().equals(client.socket())) {
                    Thread thread = entry.getKey();
                    thread.interrupt();
                    logger.info("Thread " + thread + " removed from server");
                    iterator.remove();
                    break;
                }
            }

            logger.info("Client " + client + " removed from server");

    }
    public synchronized boolean nameExists(String name) {
        for (Client client : clients) {
            if (client.name().equals(name)) {
                return true;
            }
        }
        return false;
    }

    public Client getClientByName(String name) {
        for (Client client : clients) {
            if (client.name().equals(name)) {
                return client;
            }
        }
        return null;
    }

}
