package server.handler;

import org.apache.logging.log4j.LogManager;
import org.apache.logging.log4j.Logger;
import server.Client;
import server.message.Message;
import server.Server;


import java.io.IOException;

import java.net.DatagramPacket;
import java.net.DatagramSocket;

public class ClientUDPHandler implements Runnable{



    private static final Logger logger = LogManager.getLogger(ClientUDPHandler.class);
    private final Server server;

    private final DatagramSocket serverUDPSocket;

    public ClientUDPHandler( DatagramSocket serverUDPSocket, Server server) {
        this.server = server;
        this.serverUDPSocket = serverUDPSocket;
    }

    @Override
    public void run() {
        Thread.currentThread().setName("UDP thread");
        try
        {
            while (serverUDPSocket.isBound() && !Thread.currentThread().isInterrupted()) {
                byte[] receiveData = new byte[4096];
                DatagramPacket receivePacket = new DatagramPacket(receiveData, receiveData.length);
                serverUDPSocket.receive(receivePacket);
                String sentence = new String(receivePacket.getData());
                Message message = new Message(sentence);
                logger.info("Received message from "+ message.getSender() + " : " + message.getText());
                //tu cos nie dziala
                Client client = server.getClientByName(message.getSender());
                logger.info("Received message from "+ client );
                byte[] sendData  = message.getMessage().getBytes();
//                DatagramPacket sendPacket = new DatagramPacket(sendData, sendData.length, client.socket().getInetAddress(), client.socket().getPort());


                server.sendToAllUDP(client,sendData);

            }
        }catch (IOException e) {
            logger.error("Exception caught when trying to listen on port "
                    + " run" + e.getMessage());
        }


    }

}
