package server.message;

public class Message {
    private String text;
    private MessageType type;
    private String sender;



    public Message(String message) {
        fromString(message);
    }

    private void fromString(String message) {
        String[] parts = message.split("\\|");
//        System.out.println(parts[0] + " " + parts[1]);
        switch (parts[0].toLowerCase()) {
            case "t":
                this.type = MessageType.TEXT_TCP;
                this.text = parts[1];
                this.sender = parts[2];
                break;
            case "q":
                this.type = MessageType.QUIT;
                this.text = "Quit";
                break;
            case "i":
                this.type = MessageType.INIT;
                this.text = parts[1];
                System.out.println("message: " + parts[1]);
                break;
            case "u":
                this.type = MessageType.TEXT_UDP;
                this.text = parts[1];
                this.sender = parts[2];
                break;
            default:
                this.type = MessageType.UKWN;
                this.text = "Unknown message type";
                break;
        }

    }
    public MessageType getType() {
        return type;
    }
    public String getText() {
        return text;
    }

    public String getMessage() {
        return new StringBuilder().append(sender).append(": ").append(text).toString();
    }

    public String getSender() {
        return sender;
    }
}
