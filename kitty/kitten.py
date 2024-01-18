from kittens.tui.handler import result_handler
from kitty.fast_data_types import *

def main(args): pass

@result_handler(no_ui=True)
def handle_result(args, answer, target_window_id, boss):
    # {{- print "\n" -}}{{ .Script | indent }}
    # {{- if .SendResult }}
    import socket
    import sys
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    server_address = "{{ .Socket }}"
    try:
        sock.connect(server_address)
    except socket.error as msg:
        print(msg)
        sys.exit(1)
    try:
        message = answer.encode()
        sock.sendall(message)
    finally:
        sock.close()
    # {{- end -}}
