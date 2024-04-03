from config import SERVERIP, SERVERPORT
import socket
from common import check_send_message,handle_quit_message


def init_tcp_connection():
    client_tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    client_tcp.connect((SERVERIP, SERVERPORT))
    print("Connected to server at {}:{}".format(SERVERIP, SERVERPORT))

    return client_tcp


def send_tcp_message(client, message, nickname):
    if check_send_message(message):
        final_message = message + "|" + nickname + "|\n"
        client.send(final_message.encode())
    else:
        print("Invalid message format")


def receive_tcp_messages(client_udp,client_tcp,client_multicast):
    while True:
        try:
            message = client_tcp.recv(1024).decode()
            if message.strip():
                if handle_quit_message(message,client_tcp,client_udp,client_multicast):
                    break
                print("\n-> ", message, "\nYou: ", end="")
        except Exception as e:
            print("Error: ", e)
            break
