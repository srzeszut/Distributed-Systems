package server;

import java.net.Socket;
import java.util.UUID;

public record Client(String name, Socket socket) {
    public Client {
        if (name == null || name.isBlank()) {
            throw new IllegalArgumentException("Name cannot be null or empty");
        }
    }
}
