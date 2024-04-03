import sys
import threading
from tcp_handler import init_tcp_connection, receive_tcp_messages
from udp_handler import init_udp_connection, receive_udp_messages
from message_handler import send_nickname, send_messages, send_quit_message
from multicast_handler import init_multicast_connection, receive_multicast_messages



if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python3 main.py <nickname>")
        sys.exit(1)
    nickname = sys.argv[1]
    if nickname == "server":
        print("Invalid nickname. Please choose another")
        sys.exit(1)

    client_tcp = init_tcp_connection()
    client_udp = init_udp_connection(client_tcp)
    client_multicast = init_multicast_connection()
    send_nickname(client_tcp, nickname)
    try:

        sending_thread = threading.Thread(target=send_messages,
                                          args=(client_tcp, client_udp, client_multicast, nickname))
        sending_thread.start()
        receiving_tcp_thread = threading.Thread(target=receive_tcp_messages, args=(client_udp, client_tcp,
                                                                                   client_multicast))
        receiving_tcp_thread.start()
        receiving_udp_thread = threading.Thread(target=receive_udp_messages, args=(client_udp, client_tcp,
                                                                                   client_multicast))
        receiving_udp_thread.start()
        receiving_multicast_thread = threading.Thread(target=receive_multicast_messages, args=(client_multicast,))
        receiving_multicast_thread.start()

        sending_thread.join()
        receiving_tcp_thread.join()
        receiving_udp_thread.join()
        receiving_multicast_thread.join()

    finally:
        send_quit_message(client_tcp, client_udp, client_multicast, nickname)
