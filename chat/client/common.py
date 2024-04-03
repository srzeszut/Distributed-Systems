

def check_send_message(text):
    if "|" in text[:3]:
        return True
    return False


def handle_quit_message(message,client_tcp,client_udp,client_multicast):

    if message.split("|")[0] == "server:quit":
        if len(message.split("|")) > 1:
            print(message.split("|")[1])
        print("Quit message received. Exiting...")
        try:

            client_multicast.close()
            client_udp.close()
            client_tcp.close()
        except Exception as e:
            print("Error: ", e)
        finally:
            return True
    return False






