from config import MCAST_GRP, MCAST_PORT
import socket,struct
from config import MULTICAST_TTL
from common import check_send_message


def init_multicast_connection():
    multi_sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, socket.IPPROTO_UDP)
    multi_sock.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    multi_sock.bind(('', MCAST_PORT))#moze byc bez grp
    mreq = struct.pack("4sl", socket.inet_aton(MCAST_GRP), socket.INADDR_ANY)
    multi_sock.setsockopt(socket.IPPROTO_IP, socket.IP_ADD_MEMBERSHIP, mreq)
    print("Multicast connection initialized")
    return multi_sock


def send_multicast_message(client, message, nickname):
    if check_send_message(message):

        final_message = nickname + ": " + message + "\n"
        client.sendto(final_message.encode(), (MCAST_GRP, MCAST_PORT))
    else:
        print("Invalid message format")


def receive_multicast_messages(client):
    client.setsockopt(socket.IPPROTO_IP, socket.IP_MULTICAST_TTL, MULTICAST_TTL)
    while True:
        try:
            message= client.recv(1024)
            if message.strip():
                print("\n-> ", message.decode(), "\nYou: ", end="")
        except Exception as e:
            print("Error: ", e)
            break
