import socket
from config import SERVERIP, SERVERPORT
from common import check_send_message, handle_quit_message


def init_udp_connection(client_tcp):
    client_udp = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    client_udp.bind((SERVERIP, client_tcp.getsockname()[1]))
    print("UDP connection initialized")
    return client_udp


def send_udp_message(client, message, nickname):
    if check_send_message(message):
        final_message = message + "|" + nickname + "|\n"
        client.sendto(final_message.encode(), (SERVERIP, SERVERPORT))

    else:
        print("Invalid message format")


def receive_udp_messages(client_udp,client_tcp,client_multicast):
    while True:
        try:
            message,_ = client_udp.recvfrom(4096)
            if message.strip():
                if handle_quit_message(message.decode(),client_tcp,client_udp,client_multicast):
                    break
            print("\n-> ", message.decode(), "\nYou: ", end="")
        except Exception as e:
            print("Error: ", e)
            break
