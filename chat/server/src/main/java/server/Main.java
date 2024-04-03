package server;

import java.lang.management.ManagementFactory;
import java.lang.management.ThreadInfo;
import java.lang.management.ThreadMXBean;

import static java.lang.Thread.sleep;

public class Main {
    public static void main(String[] args) {
        int portNumber = 9008;
        Server server = new Server(portNumber);
        server.start();

        while(true){

            try {
                Thread.sleep(10000);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

        }
    }
}