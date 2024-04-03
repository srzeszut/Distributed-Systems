from tcp_handler import send_tcp_message
from udp_handler import send_udp_message
from multicast_handler import send_multicast_message
from config import ASCII_ART


def send_nickname(client, nickname):
    nick = "I|"+ nickname + "\n"
    print("Sending nickname: ", nick)
    client.send(nick.encode())


def send_messages(client_tcp,client_udp,client_multicast,nickname):
    while True:
        message = input("You: ")
        try:
            msg_type = message.split("|")[0]
            match msg_type.upper():
                case "U":
                    if message.split("|")[1] == "ascii":
                        print("Sending ASCII ART")
                        send_ascii_art(client_udp, nickname)
                        continue
                    send_udp_message(client_udp, message, nickname)

                case "T":
                    send_tcp_message(client_tcp, message, nickname)

                case "Q":
                    send_quit_message(client_tcp,client_udp,client_multicast, nickname)
                    break
                case "M":
                    send_multicast_message(client_multicast, message, nickname)

                case _:
                    print("Invalid message format")
        except Exception as e:
            print("Error: ", e)
            break


def send_ascii_art(client_udp, nickname):
    send_udp_message(client_udp, ASCII_ART, nickname)


def send_quit_message(client_tcp,client_udp,client_mul, nickname):
    send_tcp_message(client_tcp, "Q|"+"quit"+"|", nickname)
    print("Quitting...")
    try:
        client_mul.close()
        client_udp.close()
        client_tcp.close()
    except Exception as e:
        print("Error: ", e)
    finally:
        exit(0)
